package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mastery-hub/backend/internal/services"
)

type PlaygroundHandler struct {
	playgroundService *services.PlaygroundService
}

func NewPlaygroundHandler(playgroundService *services.PlaygroundService) *PlaygroundHandler {
	return &PlaygroundHandler{playgroundService: playgroundService}
}

func (h *PlaygroundHandler) Run(c *gin.Context) {
	var req services.RunRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.playgroundService.Run(c.Request.Context(), req)
	if err != nil {
		if err.Error() == "code exceeds maximum size of 64KB" {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "execution failed"})
		return
	}

	if result.ExitCode == 124 {
		c.JSON(http.StatusRequestTimeout, result)
		return
	}

	c.JSON(http.StatusOK, result)
}
