package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Lesson struct {
	Slug             string `json:"slug" bson:"slug"`
	Title            string `json:"title" bson:"title"`
	Order            int    `json:"order" bson:"order"`
	ContentPath      string `json:"contentPath,omitempty" bson:"content_path"`
	HasQuiz          bool   `json:"hasQuiz" bson:"has_quiz"`
	EstimatedMinutes int    `json:"estimatedMinutes" bson:"estimated_minutes"`
}

type Module struct {
	Slug        string   `json:"slug" bson:"slug"`
	Title       string   `json:"title" bson:"title"`
	Description string   `json:"description" bson:"description"`
	Order       int      `json:"order" bson:"order"`
	Lessons     []Lesson `json:"lessons" bson:"lessons"`
}

type Track struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Slug        string             `json:"slug" bson:"slug"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
	Icon        string             `json:"icon" bson:"icon"`
	Order       int                `json:"order" bson:"order"`
	Modules     []Module           `json:"modules" bson:"modules"`
}

type TrackSummary struct {
	Slug           string `json:"slug"`
	Title          string `json:"title"`
	Description    string `json:"description"`
	Icon           string `json:"icon"`
	Order          int    `json:"order"`
	ModuleCount    int    `json:"moduleCount"`
	LessonCount    int    `json:"lessonCount"`
	EstimatedHours int    `json:"estimatedHours"`
}

type LessonDetail struct {
	Slug             string      `json:"slug"`
	Title            string      `json:"title"`
	TrackSlug        string      `json:"trackSlug"`
	ModuleSlug       string      `json:"moduleSlug"`
	EstimatedMinutes int         `json:"estimatedMinutes"`
	HasQuiz          bool        `json:"hasQuiz"`
	Content          string      `json:"content"`
	PrevLesson       *LessonLink `json:"prevLesson"`
	NextLesson       *LessonLink `json:"nextLesson"`
}

type LessonLink struct {
	Slug  string `json:"slug"`
	Title string `json:"title"`
}

type AdminLessonRequest struct {
	Slug             string `json:"slug" binding:"required"`
	Title            string `json:"title" binding:"required"`
	TrackSlug        string `json:"trackSlug" binding:"required"`
	ModuleSlug       string `json:"moduleSlug" binding:"required"`
	ContentPath      string `json:"contentPath"`
	HasQuiz          bool   `json:"hasQuiz"`
	EstimatedMinutes int    `json:"estimatedMinutes"`
}
