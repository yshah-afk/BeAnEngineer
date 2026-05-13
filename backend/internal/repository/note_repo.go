package repository

import (
	"context"
	"time"

	"github.com/mastery-hub/backend/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type NoteRepo struct {
	coll *mongo.Collection
}

func NewNoteRepo(db *mongo.Database) *NoteRepo {
	return &NoteRepo{coll: db.Collection("notes")}
}

func (r *NoteRepo) FindByUser(ctx context.Context, userID primitive.ObjectID, lessonSlug string) ([]models.Note, error) {
	filter := bson.M{"user_id": userID}
	if lessonSlug != "" {
		filter["lesson_slug"] = lessonSlug
	}

	cursor, err := r.coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []models.Note
	if err := cursor.All(ctx, &results); err != nil {
		return nil, err
	}
	return results, nil
}

func (r *NoteRepo) FindByUserAndLesson(ctx context.Context, userID primitive.ObjectID, lessonSlug string) ([]models.Note, error) {
	return r.FindByUser(ctx, userID, lessonSlug)
}

func (r *NoteRepo) FindByID(ctx context.Context, id primitive.ObjectID) (*models.Note, error) {
	var note models.Note
	err := r.coll.FindOne(ctx, bson.M{"_id": id}).Decode(&note)
	if err != nil {
		return nil, err
	}
	return &note, nil
}

func (r *NoteRepo) Create(ctx context.Context, note *models.Note) error {
	now := time.Now()
	note.CreatedAt = now
	note.UpdatedAt = now
	result, err := r.coll.InsertOne(ctx, note)
	if err != nil {
		return err
	}
	note.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (r *NoteRepo) Update(ctx context.Context, note *models.Note) error {
	note.UpdatedAt = time.Now()
	_, err := r.coll.UpdateOne(ctx, bson.M{"_id": note.ID}, bson.M{
		"$set": bson.M{
			"content":    note.Content,
			"updated_at": note.UpdatedAt,
		},
	})
	return err
}

func (r *NoteRepo) Delete(ctx context.Context, id, userID primitive.ObjectID) error {
	result, err := r.coll.DeleteOne(ctx, bson.M{"_id": id, "user_id": userID})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil
}
