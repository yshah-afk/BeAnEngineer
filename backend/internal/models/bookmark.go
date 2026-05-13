package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Bookmark struct {
	ID         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID     primitive.ObjectID `json:"userId" bson:"user_id"`
	LessonSlug string             `json:"lessonSlug" bson:"lesson_slug"`
	CreatedAt  time.Time          `json:"createdAt" bson:"created_at"`
	UpdatedAt  time.Time          `json:"updatedAt" bson:"updated_at"`
}

type CreateBookmarkRequest struct {
	LessonSlug string `json:"lessonSlug" binding:"required"`
}

type BookmarkResponse struct {
	ID         string    `json:"id"`
	LessonSlug string    `json:"lessonSlug"`
	CreatedAt  time.Time `json:"createdAt"`
}
