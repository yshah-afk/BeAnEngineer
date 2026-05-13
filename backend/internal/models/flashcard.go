package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Flashcard struct {
	ID         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	ModuleSlug string             `json:"moduleSlug" bson:"module_slug"`
	Question   string             `json:"question" bson:"question"`
	Answer     string             `json:"answer" bson:"answer"`
	Tags       []string           `json:"tags" bson:"tags"`
	Difficulty string             `json:"difficulty" bson:"difficulty"`
}

type UserFlashcardProgress struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID      primitive.ObjectID `json:"userId" bson:"user_id"`
	FlashcardID primitive.ObjectID `json:"flashcardId" bson:"flashcard_id"`
	EaseFactor  float64            `json:"easeFactor" bson:"ease_factor"`
	IntervalDays int               `json:"intervalDays" bson:"interval_days"`
	NextReview  time.Time          `json:"nextReview" bson:"next_review"`
	Repetitions int                `json:"repetitions" bson:"repetitions"`
	CreatedAt   time.Time          `json:"createdAt" bson:"created_at"`
	UpdatedAt   time.Time          `json:"updatedAt" bson:"updated_at"`
}

type ReviewRequest struct {
	Quality int `json:"quality" binding:"required,min=0,max=5"`
}

type ReviewResponse struct {
	FlashcardID  string    `json:"flashcardId"`
	EaseFactor   float64   `json:"easeFactor"`
	IntervalDays int       `json:"intervalDays"`
	NextReview   time.Time `json:"nextReview"`
	Repetitions  int       `json:"repetitions"`
}

type FlashcardWithProgress struct {
	ID           string  `json:"id"`
	Question     string  `json:"question"`
	Answer       string  `json:"answer"`
	ModuleSlug   string  `json:"moduleSlug"`
	Difficulty   string  `json:"difficulty"`
	EaseFactor   float64 `json:"easeFactor"`
	IntervalDays int     `json:"intervalDays"`
	Repetitions  int     `json:"repetitions"`
}

type DueCardsResponse struct {
	DueCount int                     `json:"dueCount"`
	Cards    []FlashcardWithProgress `json:"cards"`
}
