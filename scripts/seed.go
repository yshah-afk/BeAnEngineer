package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

type Lesson struct {
	Slug             string `json:"slug" bson:"slug"`
	Title            string `json:"title" bson:"title"`
	Order            int    `json:"order" bson:"order"`
	ContentPath      string `json:"content_path,omitempty" bson:"content_path"`
	HasQuiz          bool   `json:"has_quiz" bson:"has_quiz"`
	EstimatedMinutes int    `json:"estimated_minutes" bson:"estimated_minutes"`
}

type Module struct {
	Slug        string   `json:"slug" bson:"slug"`
	Title       string   `json:"title" bson:"title"`
	Description string   `json:"description" bson:"description"`
	Order       int      `json:"order" bson:"order"`
	Lessons     []Lesson `json:"lessons" bson:"lessons"`
}

type Track struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Slug        string             `json:"slug" bson:"slug"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
	Icon        string             `json:"icon" bson:"icon"`
	Order       int                `json:"order" bson:"order"`
	Modules     []Module           `json:"modules" bson:"modules"`
}

type Flashcard struct {
	ID         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	ModuleSlug string             `json:"module_slug" bson:"module_slug"`
	Question   string             `json:"question" bson:"question"`
	Answer     string             `json:"answer" bson:"answer"`
	Tags       []string           `json:"tags" bson:"tags"`
	Difficulty string             `json:"difficulty" bson:"difficulty"`
}

type QuizFile struct {
	Questions []json.RawMessage `json:"questions"`
}

type FlashcardFile struct {
	Cards []Flashcard `json:"cards"`
}

func main() {
	mongoURI := getEnv("MONGO_URI", "mongodb://localhost:27017")
	dbName := getEnv("MONGO_DB_NAME", "mastery_hub")
	contentDir := getEnv("CONTENT_DIR", "./content")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer client.Disconnect(ctx)

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}
	fmt.Println("Connected to MongoDB")

	db := client.Database(dbName)

	trackCount := seedTracks(ctx, db, contentDir)
	flashcardCount := seedFlashcards(ctx, db, contentDir)
	seedAdminUser(ctx, db)

	fmt.Println("\n--- Seed Summary ---")
	fmt.Printf("Tracks seeded:     %d\n", trackCount)
	fmt.Printf("Flashcards seeded: %d\n", flashcardCount)
	fmt.Println("Admin user:        admin@mastery-hub.dev / Admin123!")
	fmt.Println("Done.")
}

func seedTracks(ctx context.Context, db *mongo.Database, contentDir string) int {
	col := db.Collection("tracks")
	_, _ = col.DeleteMany(ctx, bson.M{})

	entries, err := os.ReadDir(contentDir)
	if err != nil {
		fmt.Printf("Warning: could not read content directory %q: %v\n", contentDir, err)
		return 0
	}

	var tracks []Track
	trackOrder := 0

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		trackOrder++
		trackSlug := entry.Name()
		trackPath := filepath.Join(contentDir, trackSlug)

		track := Track{
			Slug:        trackSlug,
			Title:       slugToTitle(trackSlug),
			Description: fmt.Sprintf("Learn %s from fundamentals to advanced topics.", slugToTitle(trackSlug)),
			Icon:        pickIcon(trackSlug),
			Order:       trackOrder,
		}

		moduleEntries, err := os.ReadDir(trackPath)
		if err != nil {
			continue
		}

		moduleOrder := 0
		for _, modEntry := range moduleEntries {
			if !modEntry.IsDir() {
				continue
			}
			moduleOrder++
			modSlug := modEntry.Name()
			modPath := filepath.Join(trackPath, modSlug)

			mod := Module{
				Slug:        modSlug,
				Title:       slugToTitle(modSlug),
				Description: fmt.Sprintf("Module: %s", slugToTitle(modSlug)),
				Order:       moduleOrder,
			}

			lessonFiles, err := os.ReadDir(modPath)
			if err != nil {
				continue
			}

			var lessonNames []string
			for _, lf := range lessonFiles {
				name := lf.Name()
				if !lf.IsDir() && (strings.HasSuffix(name, ".md") || strings.HasSuffix(name, ".mdx")) {
					lessonNames = append(lessonNames, name)
				}
			}
			sort.Strings(lessonNames)

			for i, name := range lessonNames {
				ext := filepath.Ext(name)
				lessonSlug := strings.TrimSuffix(name, ext)
				hasQuiz := fileExists(filepath.Join(modPath, lessonSlug+"_quiz.json"))

				mod.Lessons = append(mod.Lessons, Lesson{
					Slug:             lessonSlug,
					Title:            slugToTitle(lessonSlug),
					Order:            i + 1,
					ContentPath:      filepath.Join(trackSlug, modSlug, name),
					HasQuiz:          hasQuiz,
					EstimatedMinutes: 30,
				})
			}

			track.Modules = append(track.Modules, mod)
		}

		tracks = append(tracks, track)
	}

	if len(tracks) == 0 {
		fmt.Println("No tracks found in content directory.")
		return 0
	}

	docs := make([]interface{}, len(tracks))
	for i := range tracks {
		docs[i] = tracks[i]
	}
	res, err := col.InsertMany(ctx, docs)
	if err != nil {
		log.Fatalf("Failed to insert tracks: %v", err)
	}

	fmt.Printf("Inserted %d tracks\n", len(res.InsertedIDs))
	return len(res.InsertedIDs)
}

func seedFlashcards(ctx context.Context, db *mongo.Database, contentDir string) int {
	col := db.Collection("flashcards")
	_, _ = col.DeleteMany(ctx, bson.M{})

	var allCards []interface{}

	err := filepath.Walk(contentDir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return err
		}
		if strings.HasSuffix(info.Name(), "quiz.json") || info.Name() == "flashcards.json" {
			data, readErr := os.ReadFile(path)
			if readErr != nil {
				return nil
			}

			var file FlashcardFile
			if json.Unmarshal(data, &file) == nil && len(file.Cards) > 0 {
				for _, card := range file.Cards {
					allCards = append(allCards, card)
				}
			}
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Warning: error walking content dir for flashcards: %v\n", err)
	}

	if len(allCards) == 0 {
		fmt.Println("No flashcards found.")
		return 0
	}

	res, err := col.InsertMany(ctx, allCards)
	if err != nil {
		log.Fatalf("Failed to insert flashcards: %v", err)
	}

	fmt.Printf("Inserted %d flashcards\n", len(res.InsertedIDs))
	return len(res.InsertedIDs)
}

func seedAdminUser(ctx context.Context, db *mongo.Database) {
	col := db.Collection("users")

	count, _ := col.CountDocuments(ctx, bson.M{"email": "admin@mastery-hub.dev"})
	if count > 0 {
		fmt.Println("Admin user already exists, skipping.")
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte("Admin123!"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Failed to hash password: %v", err)
	}

	now := time.Now()
	admin := bson.M{
		"email":         "admin@mastery-hub.dev",
		"name":          "Admin",
		"avatar_url":    "",
		"auth_provider": "local",
		"password_hash": string(hash),
		"role":          "admin",
		"streak": bson.M{
			"current":       0,
			"longest":       0,
			"last_activity": now,
		},
		"created_at": now,
		"updated_at": now,
	}

	_, err = col.InsertOne(ctx, admin)
	if err != nil {
		log.Fatalf("Failed to create admin user: %v", err)
	}

	fmt.Println("Created admin user: admin@mastery-hub.dev")
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}

func slugToTitle(slug string) string {
	parts := strings.Split(slug, "-")
	for i, p := range parts {
		if len(p) > 0 {
			parts[i] = strings.ToUpper(p[:1]) + p[1:]
		}
	}
	return strings.Join(parts, " ")
}

func pickIcon(slug string) string {
	icons := map[string]string{
		"llm":       "Brain",
		"ai":        "Brain",
		"go":        "Code2",
		"golang":    "Code2",
		"react":     "Layout",
		"frontend":  "Layout",
		"fullstack": "Layers",
		"devops":    "Server",
		"system":    "Cpu",
	}
	lower := strings.ToLower(slug)
	for key, icon := range icons {
		if strings.Contains(lower, key) {
			return icon
		}
	}
	return "BookOpen"
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
