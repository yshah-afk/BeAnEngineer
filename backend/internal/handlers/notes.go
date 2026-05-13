package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mastery-hub/backend/internal/models"
	"github.com/mastery-hub/backend/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type NotesHandler struct {
	noteRepo *repository.NoteRepo
}

func NewNotesHandler(noteRepo *repository.NoteRepo) *NotesHandler {
	return &NotesHandler{noteRepo: noteRepo}
}

func (h *NotesHandler) List(c *gin.Context) {
	userID, _ := c.Get("user_id")
	oid, err := primitive.ObjectIDFromHex(userID.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	lessonSlug := c.Query("lesson_slug")
	notes, err := h.noteRepo.FindByUser(c.Request.Context(), oid, lessonSlug)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch notes"})
		return
	}

	responses := make([]models.NoteResponse, 0, len(notes))
	for _, n := range notes {
		responses = append(responses, models.NoteResponse{
			ID:         n.ID.Hex(),
			LessonSlug: n.LessonSlug,
			Content:    n.Content,
			CreatedAt:  n.CreatedAt,
			UpdatedAt:  n.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, responses)
}

func (h *NotesHandler) Create(c *gin.Context) {
	userID, _ := c.Get("user_id")
	oid, err := primitive.ObjectIDFromHex(userID.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	var req models.CreateNoteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	note := &models.Note{
		UserID:     oid,
		LessonSlug: req.LessonSlug,
		Content:    req.Content,
	}

	if err := h.noteRepo.Create(c.Request.Context(), note); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create note"})
		return
	}

	c.JSON(http.StatusCreated, models.NoteResponse{
		ID:         note.ID.Hex(),
		LessonSlug: note.LessonSlug,
		Content:    note.Content,
		CreatedAt:  note.CreatedAt,
		UpdatedAt:  note.UpdatedAt,
	})
}

func (h *NotesHandler) Update(c *gin.Context) {
	userID, _ := c.Get("user_id")
	uid, _ := primitive.ObjectIDFromHex(userID.(string))

	noteID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid note id"})
		return
	}

	existing, err := h.noteRepo.FindByID(c.Request.Context(), noteID)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "note not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch note"})
		return
	}

	if existing.UserID != uid {
		c.JSON(http.StatusForbidden, gin.H{"error": "not owner"})
		return
	}

	var req models.UpdateNoteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	existing.Content = req.Content
	if err := h.noteRepo.Update(c.Request.Context(), existing); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update note"})
		return
	}

	c.JSON(http.StatusOK, models.NoteResponse{
		ID:         existing.ID.Hex(),
		LessonSlug: existing.LessonSlug,
		Content:    existing.Content,
		CreatedAt:  existing.CreatedAt,
		UpdatedAt:  existing.UpdatedAt,
	})
}

func (h *NotesHandler) Delete(c *gin.Context) {
	userID, _ := c.Get("user_id")
	uid, _ := primitive.ObjectIDFromHex(userID.(string))

	noteID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid note id"})
		return
	}

	existing, err := h.noteRepo.FindByID(c.Request.Context(), noteID)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "note not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch note"})
		return
	}

	if existing.UserID != uid {
		c.JSON(http.StatusForbidden, gin.H{"error": "not owner"})
		return
	}

	if err := h.noteRepo.Delete(c.Request.Context(), noteID, uid); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete note"})
		return
	}

	c.Status(http.StatusNoContent)
}
