package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mastery-hub/backend/internal/models"
	"github.com/mastery-hub/backend/internal/services"
)

type FlashcardsHandler struct {
	flashcardService *services.FlashcardService
}

func NewFlashcardsHandler(flashcardService *services.FlashcardService) *FlashcardsHandler {
	return &FlashcardsHandler{flashcardService: flashcardService}
}

func (h *FlashcardsHandler) List(c *gin.Context) {
	moduleSlug := c.Query("module")

	cards, err := h.flashcardService.GetByModule(c.Request.Context(), moduleSlug)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch flashcards"})
		return
	}

	c.JSON(http.StatusOK, cards)
}

func (h *FlashcardsHandler) GetDueCards(c *gin.Context) {
	userID, _ := c.Get("user_id")
	limitStr := c.DefaultQuery("limit", "20")
	limit, err := strconv.ParseInt(limitStr, 10, 64)
	if err != nil || limit <= 0 {
		limit = 20
	}

	result, err := h.flashcardService.GetDueCards(c.Request.Context(), userID.(string), limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch due cards"})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *FlashcardsHandler) Review(c *gin.Context) {
	userID, _ := c.Get("user_id")
	flashcardID := c.Param("id")

	var req models.ReviewRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.flashcardService.ReviewCard(c.Request.Context(), userID.(string), flashcardID, req.Quality)
	if err != nil {
		if err.Error() == "flashcard not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "flashcard not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to review flashcard"})
		return
	}

	c.JSON(http.StatusOK, result)
}
