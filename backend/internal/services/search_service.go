package services

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SearchResult struct {
	Kind       string  `json:"kind"`
	Title      string  `json:"title,omitempty"`
	Slug       string  `json:"slug,omitempty"`
	TrackSlug  string  `json:"trackSlug,omitempty"`
	ModuleSlug string  `json:"moduleSlug,omitempty"`
	Excerpt    string  `json:"excerpt,omitempty"`
	Question   string  `json:"question,omitempty"`
	Score      float64 `json:"score"`
}

type SearchResponse struct {
	Query   string         `json:"query"`
	Type    string         `json:"type"`
	Total   int            `json:"total"`
	Results []SearchResult `json:"results"`
}

type SearchService struct {
	db *mongo.Database
}

func NewSearchService(db *mongo.Database) *SearchService {
	return &SearchService{db: db}
}

func (s *SearchService) Search(ctx context.Context, query, searchType string) (*SearchResponse, error) {
	results := make([]SearchResult, 0)

	trackResults, err := s.searchTracks(ctx, query)
	if err == nil {
		results = append(results, trackResults...)
	}

	flashcardResults, err := s.searchFlashcards(ctx, query)
	if err == nil {
		results = append(results, flashcardResults...)
	}

	return &SearchResponse{
		Query:   query,
		Type:    searchType,
		Total:   len(results),
		Results: results,
	}, nil
}

func (s *SearchService) searchTracks(ctx context.Context, query string) ([]SearchResult, error) {
	filter := bson.M{
		"$text": bson.M{"$search": query},
	}
	opts := options.Find().
		SetProjection(bson.M{"score": bson.M{"$meta": "textScore"}}).
		SetSort(bson.M{"score": bson.M{"$meta": "textScore"}}).
		SetLimit(20)

	coll := s.db.Collection("tracks")

	textIdx := mongo.IndexModel{
		Keys: bson.D{
			{Key: "title", Value: "text"},
			{Key: "description", Value: "text"},
			{Key: "modules.title", Value: "text"},
			{Key: "modules.lessons.title", Value: "text"},
		},
	}
	_, _ = coll.Indexes().CreateOne(ctx, textIdx)

	cursor, err := coll.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []SearchResult
	for cursor.Next(ctx) {
		var doc struct {
			Slug        string  `bson:"slug"`
			Title       string  `bson:"title"`
			Description string  `bson:"description"`
			Score       float64 `bson:"score"`
		}
		if err := cursor.Decode(&doc); err != nil {
			continue
		}
		results = append(results, SearchResult{
			Kind:    "lesson",
			Title:   doc.Title,
			Slug:    doc.Slug,
			Excerpt: doc.Description,
			Score:   doc.Score,
		})
	}
	return results, nil
}

func (s *SearchService) searchFlashcards(ctx context.Context, query string) ([]SearchResult, error) {
	filter := bson.M{
		"$text": bson.M{"$search": query},
	}
	opts := options.Find().
		SetProjection(bson.M{"score": bson.M{"$meta": "textScore"}}).
		SetSort(bson.M{"score": bson.M{"$meta": "textScore"}}).
		SetLimit(20)

	cursor, err := s.db.Collection("flashcards").Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []SearchResult
	for cursor.Next(ctx) {
		var doc struct {
			ModuleSlug string  `bson:"module_slug"`
			Question   string  `bson:"question"`
			Score      float64 `bson:"score"`
		}
		if err := cursor.Decode(&doc); err != nil {
			continue
		}
		results = append(results, SearchResult{
			Kind:       "flashcard",
			Question:   doc.Question,
			ModuleSlug: doc.ModuleSlug,
			Score:      doc.Score,
		})
	}
	return results, nil
}
