package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mastery-hub/backend/internal/models"
	"github.com/mastery-hub/backend/internal/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type TracksHandler struct {
	trackRepo *repository.TrackRepo
}

func NewTracksHandler(trackRepo *repository.TrackRepo) *TracksHandler {
	return &TracksHandler{trackRepo: trackRepo}
}

func (h *TracksHandler) List(c *gin.Context) {
	tracks, err := h.trackRepo.FindAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch tracks"})
		return
	}

	summaries := make([]models.TrackSummary, 0, len(tracks))
	for _, t := range tracks {
		lessonCount := 0
		totalMinutes := 0
		for _, m := range t.Modules {
			lessonCount += len(m.Lessons)
			for _, l := range m.Lessons {
				totalMinutes += l.EstimatedMinutes
			}
		}
		summaries = append(summaries, models.TrackSummary{
			Slug:           t.Slug,
			Title:          t.Title,
			Description:    t.Description,
			Icon:           t.Icon,
			Order:          t.Order,
			ModuleCount:    len(t.Modules),
			LessonCount:    lessonCount,
			EstimatedHours: totalMinutes / 60,
		})
	}

	c.JSON(http.StatusOK, summaries)
}

type trackDetailResponse struct {
	ID             string           `json:"id"`
	Slug           string           `json:"slug"`
	Title          string           `json:"title"`
	Description    string           `json:"description"`
	Icon           string           `json:"icon"`
	ModuleCount    int              `json:"moduleCount"`
	LessonCount    int              `json:"lessonCount"`
	EstimatedHours int              `json:"estimatedHours"`
	Modules        []models.Module  `json:"modules"`
}

func (h *TracksHandler) GetBySlug(c *gin.Context) {
	slug := c.Param("slug")
	track, err := h.trackRepo.FindBySlug(c.Request.Context(), slug)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "track not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch track"})
		return
	}

	lessonCount := 0
	totalMinutes := 0
	for _, m := range track.Modules {
		lessonCount += len(m.Lessons)
		for _, l := range m.Lessons {
			totalMinutes += l.EstimatedMinutes
		}
	}

	c.JSON(http.StatusOK, trackDetailResponse{
		ID:             track.ID.Hex(),
		Slug:           track.Slug,
		Title:          track.Title,
		Description:    track.Description,
		Icon:           track.Icon,
		ModuleCount:    len(track.Modules),
		LessonCount:    lessonCount,
		EstimatedHours: totalMinutes / 60,
		Modules:        track.Modules,
	})
}
