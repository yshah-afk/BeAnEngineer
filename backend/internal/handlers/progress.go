package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mastery-hub/backend/internal/models"
	"github.com/mastery-hub/backend/internal/services"
)

type ProgressHandler struct {
	progressService *services.ProgressService
	lessonService   *services.LessonService
}

func NewProgressHandler(progressService *services.ProgressService, lessonService *services.LessonService) *ProgressHandler {
	return &ProgressHandler{progressService: progressService, lessonService: lessonService}
}

func (h *ProgressHandler) GetUserProgress(c *gin.Context) {
	userID, _ := c.Get("user_id")
	trackSlug := c.Query("track_slug")

	progress, err := h.progressService.GetUserProgress(c.Request.Context(), userID.(string), trackSlug)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch progress"})
		return
	}

	c.JSON(http.StatusOK, progress)
}

func (h *ProgressHandler) UpdateProgress(c *gin.Context) {
	userID, _ := c.Get("user_id")
	lessonSlug := c.Param("lessonSlug")

	var req models.UpdateProgressRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	progress, err := h.progressService.UpdateProgress(c.Request.Context(), userID.(string), lessonSlug, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update progress"})
		return
	}

	c.JSON(http.StatusOK, models.ProgressResponse{
		LessonSlug:  progress.LessonSlug,
		Status:      progress.Status,
		CompletedAt: progress.CompletedAt,
		TimeSpent:   progress.TimeSpent,
	})
}

func (h *ProgressHandler) SubmitQuiz(c *gin.Context) {
	userID, _ := c.Get("user_id")
	lessonSlug := c.Param("lessonSlug")

	var req models.QuizSubmission
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	questions, err := h.lessonService.GetQuiz(c.Request.Context(), req.TrackSlug, req.ModuleSlug)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "quiz not found"})
		return
	}

	score, total, results := h.lessonService.GradeQuiz(questions, req.Answers)

	quizScore := models.QuizScore{
		Score:       score,
		Total:       total,
		SubmittedAt: time.Now(),
	}
	if err := h.progressService.SubmitQuiz(c.Request.Context(), userID.(string), lessonSlug, quizScore); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save quiz score"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"score":   score,
		"total":   total,
		"results": results,
	})
}

func (h *ProgressHandler) GetStreaks(c *gin.Context) {
	userID, _ := c.Get("user_id")

	streaks, err := h.progressService.GetStreaks(c.Request.Context(), userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch streaks"})
		return
	}

	c.JSON(http.StatusOK, streaks)
}
