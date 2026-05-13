package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mastery-hub/backend/internal/services"
)

type SearchHandler struct {
	searchService *services.SearchService
}

func NewSearchHandler(searchService *services.SearchService) *SearchHandler {
	return &SearchHandler{searchService: searchService}
}

func (h *SearchHandler) Search(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing query parameter 'q'"})
		return
	}

	searchType := c.DefaultQuery("type", "fulltext")

	result, err := h.searchService.Search(c.Request.Context(), query, searchType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "search failed"})
		return
	}

	c.JSON(http.StatusOK, result)
}
