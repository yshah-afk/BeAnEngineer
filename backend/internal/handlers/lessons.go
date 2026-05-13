package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mastery-hub/backend/internal/services"
	"go.mongodb.org/mongo-driver/mongo"
)

type LessonsHandler struct {
	lessonService *services.LessonService
}

func NewLessonsHandler(lessonService *services.LessonService) *LessonsHandler {
	return &LessonsHandler{lessonService: lessonService}
}

func (h *LessonsHandler) GetLesson(c *gin.Context) {
	trackSlug := c.Param("trackSlug")
	moduleSlug := c.Param("moduleSlug")
	lessonSlug := c.Param("lessonSlug")

	lesson, err := h.lessonService.GetLesson(c.Request.Context(), trackSlug, moduleSlug, lessonSlug)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "lesson not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch lesson"})
		return
	}

	c.JSON(http.StatusOK, lesson)
}

func (h *LessonsHandler) GetLessonBySlug(c *gin.Context) {
	slug := c.Param("slug")

	lesson, err := h.lessonService.GetLessonBySlug(c.Request.Context(), slug)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "lesson not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch lesson"})
		return
	}

	c.JSON(http.StatusOK, lesson)
}

func (h *LessonsHandler) GetQuiz(c *gin.Context) {
	trackSlug := c.Param("trackSlug")
	moduleSlug := c.Param("moduleSlug")

	questions, err := h.lessonService.GetQuiz(c.Request.Context(), trackSlug, moduleSlug)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "quiz not found"})
		return
	}

	sanitized := make([]gin.H, 0, len(questions))
	for _, q := range questions {
		sanitized = append(sanitized, gin.H{
			"id":      q.ID,
			"text":    q.Text,
			"options": q.Options,
			"type":    q.Type,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"module_slug": moduleSlug,
		"questions":   sanitized,
	})
}
