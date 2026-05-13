package handlers

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mastery-hub/backend/internal/services"
)

type TutorHandler struct {
	tutorService *services.TutorService
}

func NewTutorHandler(tutorService *services.TutorService) *TutorHandler {
	return &TutorHandler{tutorService: tutorService}
}

func (h *TutorHandler) Chat(c *gin.Context) {
	var req services.TutorRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")
	c.Header("X-Accel-Buffering", "no")

	flusher, ok := c.Writer.(http.Flusher)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "streaming not supported"})
		return
	}

	writer := func(token string) error {
		data, _ := json.Marshal(map[string]string{"token": token})
		_, err := fmt.Fprintf(c.Writer, "data: %s\n\n", data)
		if err != nil {
			return err
		}
		flusher.Flush()
		return nil
	}

	doneFn := func(promptTokens, completionTokens int) {
		data, _ := json.Marshal(map[string]interface{}{
			"done": true,
			"usage": map[string]int{
				"prompt_tokens":     promptTokens,
				"completion_tokens": completionTokens,
			},
		})
		fmt.Fprintf(c.Writer, "data: %s\n\n", data)
		flusher.Flush()
	}

	if err := h.tutorService.StreamChat(c.Request.Context(), req, writer, doneFn); err != nil {
		slog.Error("tutor stream error", "error", err)
		data, _ := json.Marshal(map[string]string{"error": "AI provider unavailable"})
		fmt.Fprintf(c.Writer, "data: %s\n\n", data)
		flusher.Flush()
	}
}
