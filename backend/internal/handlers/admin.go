package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mastery-hub/backend/internal/models"
	"github.com/mastery-hub/backend/internal/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type AdminHandler struct {
	trackRepo     *repository.TrackRepo
	flashcardRepo *repository.FlashcardRepo
}

func NewAdminHandler(trackRepo *repository.TrackRepo, flashcardRepo *repository.FlashcardRepo) *AdminHandler {
	return &AdminHandler{trackRepo: trackRepo, flashcardRepo: flashcardRepo}
}

func (h *AdminHandler) ListLessons(c *gin.Context) {
	var query models.PaginationQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		query = models.PaginationQuery{Page: 1, PerPage: 20}
	}
	trackSlugFilter := c.Query("track_slug")

	tracks, err := h.trackRepo.FindAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch tracks"})
		return
	}

	type lessonItem struct {
		Slug       string `json:"slug"`
		Title      string `json:"title"`
		TrackSlug  string `json:"trackSlug"`
		ModuleSlug string `json:"moduleSlug"`
		HasQuiz    bool   `json:"hasQuiz"`
	}

	var allLessons []lessonItem
	for _, t := range tracks {
		if trackSlugFilter != "" && t.Slug != trackSlugFilter {
			continue
		}
		for _, m := range t.Modules {
			for _, l := range m.Lessons {
				allLessons = append(allLessons, lessonItem{
					Slug:       l.Slug,
					Title:      l.Title,
					TrackSlug:  t.Slug,
					ModuleSlug: m.Slug,
					HasQuiz:    l.HasQuiz,
				})
			}
		}
	}

	total := len(allLessons)
	start := int(query.Skip())
	if start > total {
		start = total
	}
	end := start + query.PerPage
	if end > total {
		end = total
	}

	c.JSON(http.StatusOK, models.PaginatedResponse{
		Total:   int64(total),
		Page:    query.Page,
		PerPage: query.PerPage,
		Items:   allLessons[start:end],
	})
}

func (h *AdminHandler) CreateLesson(c *gin.Context) {
	var req models.AdminLessonRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	lesson := models.Lesson{
		Slug:             req.Slug,
		Title:            req.Title,
		ContentPath:      req.ContentPath,
		HasQuiz:          req.HasQuiz,
		EstimatedMinutes: req.EstimatedMinutes,
	}

	if err := h.trackRepo.AddLesson(c.Request.Context(), req.TrackSlug, req.ModuleSlug, lesson); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create lesson"})
		return
	}

	c.JSON(http.StatusCreated, lesson)
}

func (h *AdminHandler) UpdateLesson(c *gin.Context) {
	slug := c.Param("slug")

	var req models.AdminLessonRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	lesson := models.Lesson{
		Slug:             slug,
		Title:            req.Title,
		ContentPath:      req.ContentPath,
		HasQuiz:          req.HasQuiz,
		EstimatedMinutes: req.EstimatedMinutes,
	}

	if err := h.trackRepo.UpdateLesson(c.Request.Context(), req.TrackSlug, req.ModuleSlug, lesson); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update lesson"})
		return
	}

	c.JSON(http.StatusOK, lesson)
}

func (h *AdminHandler) DeleteLesson(c *gin.Context) {
	slug := c.Param("slug")
	trackSlug := c.Query("track_slug")
	moduleSlug := c.Query("module_slug")

	if trackSlug == "" || moduleSlug == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "track_slug and module_slug are required"})
		return
	}

	if err := h.trackRepo.RemoveLesson(c.Request.Context(), trackSlug, moduleSlug, slug); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete lesson"})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *AdminHandler) Seed(c *gin.Context) {
	tracks, err := h.trackRepo.FindAll(c.Request.Context())
	if err != nil && err != mongo.ErrNoDocuments {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to check existing data"})
		return
	}

	trackCount := len(tracks)
	moduleCount := 0
	lessonCount := 0
	for _, t := range tracks {
		moduleCount += len(t.Modules)
		for _, m := range t.Modules {
			lessonCount += len(m.Lessons)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"tracks_seeded":     trackCount,
		"modules_seeded":    moduleCount,
		"lessons_seeded":    lessonCount,
		"flashcards_seeded": 0,
	})
}
