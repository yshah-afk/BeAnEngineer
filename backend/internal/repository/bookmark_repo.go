package repository

import (
	"context"
	"time"

	"github.com/mastery-hub/backend/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BookmarkRepo struct {
	coll *mongo.Collection
}

func NewBookmarkRepo(db *mongo.Database) *BookmarkRepo {
	return &BookmarkRepo{coll: db.Collection("bookmarks")}
}

func (r *BookmarkRepo) FindByUser(ctx context.Context, userID primitive.ObjectID) ([]models.Bookmark, error) {
	cursor, err := r.coll.Find(ctx, bson.M{"user_id": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []models.Bookmark
	if err := cursor.All(ctx, &results); err != nil {
		return nil, err
	}
	return results, nil
}

func (r *BookmarkRepo) Create(ctx context.Context, bookmark *models.Bookmark) error {
	now := time.Now()
	bookmark.CreatedAt = now
	bookmark.UpdatedAt = now
	result, err := r.coll.InsertOne(ctx, bookmark)
	if err != nil {
		return err
	}
	bookmark.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (r *BookmarkRepo) Delete(ctx context.Context, id, userID primitive.ObjectID) error {
	result, err := r.coll.DeleteOne(ctx, bson.M{"_id": id, "user_id": userID})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil
}

func (r *BookmarkRepo) Exists(ctx context.Context, userID primitive.ObjectID, lessonSlug string) (bool, error) {
	count, err := r.coll.CountDocuments(ctx, bson.M{
		"user_id":     userID,
		"lesson_slug": lessonSlug,
	})
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *BookmarkRepo) FindByID(ctx context.Context, id primitive.ObjectID) (*models.Bookmark, error) {
	var b models.Bookmark
	err := r.coll.FindOne(ctx, bson.M{"_id": id}).Decode(&b)
	if err != nil {
		return nil, err
	}
	return &b, nil
}
