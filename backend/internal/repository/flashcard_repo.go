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

type FlashcardRepo struct {
	cardsColl    *mongo.Collection
	progressColl *mongo.Collection
}

func NewFlashcardRepo(db *mongo.Database) *FlashcardRepo {
	return &FlashcardRepo{
		cardsColl:    db.Collection("flashcards"),
		progressColl: db.Collection("user_flashcard_progress"),
	}
}

func (r *FlashcardRepo) FindByModule(ctx context.Context, moduleSlug string) ([]models.Flashcard, error) {
	filter := bson.M{}
	if moduleSlug != "" {
		filter["module_slug"] = moduleSlug
	}

	cursor, err := r.cardsColl.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var cards []models.Flashcard
	if err := cursor.All(ctx, &cards); err != nil {
		return nil, err
	}
	return cards, nil
}

func (r *FlashcardRepo) FindByID(ctx context.Context, id primitive.ObjectID) (*models.Flashcard, error) {
	var card models.Flashcard
	err := r.cardsColl.FindOne(ctx, bson.M{"_id": id}).Decode(&card)
	if err != nil {
		return nil, err
	}
	return &card, nil
}

func (r *FlashcardRepo) FindDueForUser(ctx context.Context, userID primitive.ObjectID, limit int64) ([]models.UserFlashcardProgress, error) {
	now := time.Now()
	filter := bson.M{
		"user_id":     userID,
		"next_review": bson.M{"$lte": now},
	}
	opts := options.Find().SetLimit(limit).SetSort(bson.D{{Key: "next_review", Value: 1}})

	cursor, err := r.progressColl.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []models.UserFlashcardProgress
	if err := cursor.All(ctx, &results); err != nil {
		return nil, err
	}
	return results, nil
}

func (r *FlashcardRepo) CountDueForUser(ctx context.Context, userID primitive.ObjectID) (int64, error) {
	return r.progressColl.CountDocuments(ctx, bson.M{
		"user_id":     userID,
		"next_review": bson.M{"$lte": time.Now()},
	})
}

func (r *FlashcardRepo) FindProgress(ctx context.Context, userID, flashcardID primitive.ObjectID) (*models.UserFlashcardProgress, error) {
	var p models.UserFlashcardProgress
	err := r.progressColl.FindOne(ctx, bson.M{
		"user_id":      userID,
		"flashcard_id": flashcardID,
	}).Decode(&p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *FlashcardRepo) UpsertProgress(ctx context.Context, p *models.UserFlashcardProgress) error {
	now := time.Now()
	p.UpdatedAt = now

	filter := bson.M{
		"user_id":      p.UserID,
		"flashcard_id": p.FlashcardID,
	}
	update := bson.M{
		"$set": bson.M{
			"ease_factor":   p.EaseFactor,
			"interval_days": p.IntervalDays,
			"next_review":   p.NextReview,
			"repetitions":   p.Repetitions,
			"updated_at":    now,
		},
		"$setOnInsert": bson.M{
			"user_id":      p.UserID,
			"flashcard_id": p.FlashcardID,
			"created_at":   now,
		},
	}
	opts := options.Update().SetUpsert(true)
	_, err := r.progressColl.UpdateOne(ctx, filter, update, opts)
	return err
}

func (r *FlashcardRepo) BulkInsertCards(ctx context.Context, cards []models.Flashcard) error {
	if len(cards) == 0 {
		return nil
	}
	docs := make([]interface{}, len(cards))
	for i := range cards {
		docs[i] = cards[i]
	}
	_, err := r.cardsColl.InsertMany(ctx, docs)
	return err
}
