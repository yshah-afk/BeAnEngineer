package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Note struct {
	ID         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID     primitive.ObjectID `json:"userId" bson:"user_id"`
	LessonSlug string             `json:"lessonSlug" bson:"lesson_slug"`
	Content    string             `json:"content" bson:"content"`
	CreatedAt  time.Time          `json:"createdAt" bson:"created_at"`
	UpdatedAt  time.Time          `json:"updatedAt" bson:"updated_at"`
}

type CreateNoteRequest struct {
	LessonSlug string `json:"lessonSlug" binding:"required"`
	Content    string `json:"content" binding:"required,min=1"`
}

type UpdateNoteRequest struct {
	Content string `json:"content" binding:"required,min=1"`
}

type NoteResponse struct {
	ID         string    `json:"id"`
	LessonSlug string    `json:"lessonSlug"`
	Content    string    `json:"content"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
