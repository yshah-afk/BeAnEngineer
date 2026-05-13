package repository

import (
	"context"
	"log/slog"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(ctx context.Context, uri, dbName string) (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	clientOpts := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	db := client.Database(dbName)
	slog.Info("connected to MongoDB", "database", dbName)
	return db, nil
}

func EnsureIndexes(ctx context.Context, db *mongo.Database) error {
	indexes := map[string][]mongo.IndexModel{
		"users": {
			{
				Keys:    bson.D{{Key: "email", Value: 1}},
				Options: options.Index().SetUnique(true),
			},
			{
				Keys: bson.D{{Key: "auth_provider", Value: 1}, {Key: "provider_id", Value: 1}},
				Options: options.Index().SetUnique(true).SetPartialFilterExpression(bson.M{
					"provider_id": bson.M{"$gt": ""},
				}),
			},
		},
		"tracks": {
			{
				Keys:    bson.D{{Key: "slug", Value: 1}},
				Options: options.Index().SetUnique(true),
			},
		},
		"progress": {
			{
				Keys:    bson.D{{Key: "user_id", Value: 1}, {Key: "lesson_slug", Value: 1}},
				Options: options.Index().SetUnique(true),
			},
			{
				Keys: bson.D{{Key: "user_id", Value: 1}, {Key: "track_slug", Value: 1}},
			},
		},
		"bookmarks": {
			{
				Keys:    bson.D{{Key: "user_id", Value: 1}, {Key: "lesson_slug", Value: 1}},
				Options: options.Index().SetUnique(true),
			},
		},
		"notes": {
			{
				Keys: bson.D{{Key: "user_id", Value: 1}, {Key: "lesson_slug", Value: 1}},
			},
		},
		"flashcards": {
			{
				Keys: bson.D{{Key: "module_slug", Value: 1}, {Key: "difficulty", Value: 1}},
			},
			{
				Keys: bson.D{{Key: "question", Value: "text"}, {Key: "answer", Value: "text"}},
			},
		},
		"user_flashcard_progress": {
			{
				Keys:    bson.D{{Key: "user_id", Value: 1}, {Key: "flashcard_id", Value: 1}},
				Options: options.Index().SetUnique(true),
			},
			{
				Keys: bson.D{{Key: "user_id", Value: 1}, {Key: "next_review", Value: 1}},
			},
		},
	}

	for coll, models := range indexes {
		_, err := db.Collection(coll).Indexes().CreateMany(ctx, models)
		if err != nil {
			slog.Error("failed to create indexes", "collection", coll, "error", err)
			return err
		}
		slog.Info("indexes ensured", "collection", coll, "count", len(models))
	}
	return nil
}

func Collection(db *mongo.Database, name string) *mongo.Collection {
	return db.Collection(name)
}
