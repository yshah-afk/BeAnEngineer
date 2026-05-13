package models

import "time"

type PaginationQuery struct {
	Page    int `form:"page,default=1" binding:"min=1"`
	PerPage int `form:"per_page,default=20" binding:"min=1,max=100"`
}

func (p PaginationQuery) Skip() int64 {
	return int64((p.Page - 1) * p.PerPage)
}

func (p PaginationQuery) Limit() int64 {
	return int64(p.PerPage)
}

type PaginatedResponse struct {
	Total   int64       `json:"total"`
	Page    int         `json:"page"`
	PerPage int         `json:"perPage"`
	Items   interface{} `json:"items"`
}

type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewAPIError(code int, message string) APIError {
	return APIError{Code: code, Message: message}
}

type QuizQuestion struct {
	ID      string   `json:"id" bson:"id"`
	Text    string   `json:"text" bson:"text"`
	Options []string `json:"options" bson:"options"`
	Type    string   `json:"type" bson:"type"`
	Answer  string   `json:"-" bson:"answer"`
}

type QuizSubmission struct {
	TrackSlug  string       `json:"trackSlug" binding:"required"`
	ModuleSlug string       `json:"moduleSlug" binding:"required"`
	Answers    []QuizAnswer `json:"answers" binding:"required,dive"`
}

type QuizAnswer struct {
	QuestionID string `json:"questionId" binding:"required"`
	Selected   string `json:"selected" binding:"required"`
}

type QuizResult struct {
	QuestionID    string `json:"questionId"`
	Correct       bool   `json:"correct"`
	CorrectAnswer string `json:"correctAnswer,omitempty"`
}

type QuizScore struct {
	Score       int       `json:"score" bson:"score"`
	Total       int       `json:"total" bson:"total"`
	SubmittedAt time.Time `json:"submittedAt" bson:"submitted_at"`
}
