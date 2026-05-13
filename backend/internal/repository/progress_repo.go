package repository

import (
	"context"
	"time"

	"github.com/mastery-hub/backend/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ProgressRepo struct {
	coll *mongo.Collection
}

func NewProgressRepo(db *mongo.Database) *ProgressRepo {
	return &ProgressRepo{coll: db.Collection("progress")}
}

func (r *ProgressRepo) FindByUser(ctx context.Context, userID primitive.ObjectID, trackSlug string) ([]models.Progress, error) {
	filter := bson.M{"user_id": userID}
	if trackSlug != "" {
		filter["track_slug"] = trackSlug
	}

	cursor, err := r.coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []models.Progress
	if err := cursor.All(ctx, &results); err != nil {
		return nil, err
	}
	return results, nil
}

func (r *ProgressRepo) FindByUserAndLesson(ctx context.Context, userID primitive.ObjectID, lessonSlug string) (*models.Progress, error) {
	var p models.Progress
	err := r.coll.FindOne(ctx, bson.M{
		"user_id":     userID,
		"lesson_slug": lessonSlug,
	}).Decode(&p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *ProgressRepo) Upsert(ctx context.Context, p *models.Progress) error {
	now := time.Now()
	p.UpdatedAt = now

	filter := bson.M{
		"user_id":     p.UserID,
		"lesson_slug": p.LessonSlug,
	}

	update := bson.M{
		"$set": bson.M{
			"status":       p.Status,
			"track_slug":   p.TrackSlug,
			"module_slug":  p.ModuleSlug,
			"completed_at": p.CompletedAt,
			"time_spent":   p.TimeSpent,
			"updated_at":   now,
		},
		"$setOnInsert": bson.M{
			"created_at": now,
		},
	}

	opts := options.Update().SetUpsert(true)
	result, err := r.coll.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return err
	}
	if result.UpsertedID != nil {
		p.ID = result.UpsertedID.(primitive.ObjectID)
		p.CreatedAt = now
	}
	return nil
}

func (r *ProgressRepo) AddQuizScore(ctx context.Context, userID primitive.ObjectID, lessonSlug string, score models.QuizScore) error {
	filter := bson.M{
		"user_id":     userID,
		"lesson_slug": lessonSlug,
	}
	update := bson.M{
		"$push": bson.M{"quiz_scores": score},
		"$set":  bson.M{"updated_at": time.Now()},
	}
	_, err := r.coll.UpdateOne(ctx, filter, update)
	return err
}

func (r *ProgressRepo) CountCompletedOnDate(ctx context.Context, userID primitive.ObjectID, date time.Time) (int64, error) {
	startOfDay := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC)
	endOfDay := startOfDay.Add(24 * time.Hour)

	return r.coll.CountDocuments(ctx, bson.M{
		"user_id":      userID,
		"status":       "completed",
		"completed_at": bson.M{"$gte": startOfDay, "$lt": endOfDay},
	})
}
