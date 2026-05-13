package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mastery-hub/backend/internal/models"
	"github.com/mastery-hub/backend/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BookmarksHandler struct {
	bookmarkRepo *repository.BookmarkRepo
}

func NewBookmarksHandler(bookmarkRepo *repository.BookmarkRepo) *BookmarksHandler {
	return &BookmarksHandler{bookmarkRepo: bookmarkRepo}
}

func (h *BookmarksHandler) List(c *gin.Context) {
	userID, _ := c.Get("user_id")
	oid, err := primitive.ObjectIDFromHex(userID.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	bookmarks, err := h.bookmarkRepo.FindByUser(c.Request.Context(), oid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch bookmarks"})
		return
	}

	responses := make([]models.BookmarkResponse, 0, len(bookmarks))
	for _, b := range bookmarks {
		responses = append(responses, models.BookmarkResponse{
			ID:         b.ID.Hex(),
			LessonSlug: b.LessonSlug,
			CreatedAt:  b.CreatedAt,
		})
	}

	c.JSON(http.StatusOK, responses)
}

func (h *BookmarksHandler) Create(c *gin.Context) {
	userID, _ := c.Get("user_id")
	oid, err := primitive.ObjectIDFromHex(userID.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	var req models.CreateBookmarkRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	exists, err := h.bookmarkRepo.Exists(c.Request.Context(), oid, req.LessonSlug)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to check bookmark"})
		return
	}
	if exists {
		c.JSON(http.StatusConflict, gin.H{"error": "already bookmarked"})
		return
	}

	bookmark := &models.Bookmark{
		UserID:     oid,
		LessonSlug: req.LessonSlug,
	}

	if err := h.bookmarkRepo.Create(c.Request.Context(), bookmark); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create bookmark"})
		return
	}

	c.JSON(http.StatusCreated, models.BookmarkResponse{
		ID:         bookmark.ID.Hex(),
		LessonSlug: bookmark.LessonSlug,
		CreatedAt:  bookmark.CreatedAt,
	})
}

func (h *BookmarksHandler) Delete(c *gin.Context) {
	userID, _ := c.Get("user_id")
	uid, _ := primitive.ObjectIDFromHex(userID.(string))

	bookmarkID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid bookmark id"})
		return
	}

	bookmark, err := h.bookmarkRepo.FindByID(c.Request.Context(), bookmarkID)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "bookmark not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch bookmark"})
		return
	}

	if bookmark.UserID != uid {
		c.JSON(http.StatusForbidden, gin.H{"error": "not owner"})
		return
	}

	if err := h.bookmarkRepo.Delete(c.Request.Context(), bookmarkID, uid); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete bookmark"})
		return
	}

	c.Status(http.StatusNoContent)
}
