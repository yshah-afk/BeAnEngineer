package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Progress struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID      primitive.ObjectID `json:"userId" bson:"user_id"`
	LessonSlug  string             `json:"lessonSlug" bson:"lesson_slug"`
	TrackSlug   string             `json:"trackSlug" bson:"track_slug"`
	ModuleSlug  string             `json:"moduleSlug" bson:"module_slug"`
	Status      string             `json:"status" bson:"status"`
	CompletedAt *time.Time         `json:"completedAt,omitempty" bson:"completed_at,omitempty"`
	QuizScores  []QuizScore        `json:"quizScores" bson:"quiz_scores"`
	TimeSpent   int                `json:"timeSpent" bson:"time_spent"`
	CreatedAt   time.Time          `json:"createdAt" bson:"created_at"`
	UpdatedAt   time.Time          `json:"updatedAt" bson:"updated_at"`
}

type UpdateProgressRequest struct {
	Status         string `json:"status" binding:"required,oneof=not_started in_progress completed"`
	TrackSlug      string `json:"trackSlug" binding:"required"`
	ModuleSlug     string `json:"moduleSlug" binding:"required"`
	TimeSpentDelta int    `json:"timeSpentDelta"`
}

type ProgressResponse struct {
	LessonSlug  string      `json:"lessonSlug"`
	Status      string      `json:"status"`
	CompletedAt *time.Time  `json:"completedAt,omitempty"`
	TimeSpent   int         `json:"timeSpent"`
}

type StreakResponse struct {
	Current      int              `json:"current"`
	Longest      int              `json:"longest"`
	LastActivity time.Time        `json:"lastActivity"`
	History      []StreakDayEntry `json:"history"`
}

type StreakDayEntry struct {
	Date             string `json:"date"`
	LessonsCompleted int    `json:"lessonsCompleted"`
}
