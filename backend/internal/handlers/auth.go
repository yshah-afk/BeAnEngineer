package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mastery-hub/backend/internal/models"
	"github.com/mastery-hub/backend/internal/services"
)

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req models.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, token, err := h.authService.Register(c.Request.Context(), req)
	if err != nil {
		if errors.Is(err, services.ErrUserExists) {
			c.JSON(http.StatusConflict, gin.H{"error": "email already exists"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	refreshToken, err := h.authService.GenerateRefreshToken(user.ID.Hex())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate refresh token"})
		return
	}

	h.setRefreshCookie(c, refreshToken)

	c.JSON(http.StatusCreated, models.AuthResponse{
		User: models.UserResponse{
			ID:    user.ID.Hex(),
			Email: user.Email,
			Name:  user.Name,
			Role:  user.Role,
		},
		AccessToken: token,
		ExpiresIn:   h.authService.AccessTTLSeconds(),
	})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, token, err := h.authService.Login(c.Request.Context(), req)
	if err != nil {
		if errors.Is(err, services.ErrInvalidCredentials) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "login failed"})
		return
	}

	refreshToken, err := h.authService.GenerateRefreshToken(user.ID.Hex())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate refresh token"})
		return
	}

	h.setRefreshCookie(c, refreshToken)

	c.JSON(http.StatusOK, models.AuthResponse{
		User: models.UserResponse{
			ID:    user.ID.Hex(),
			Email: user.Email,
			Name:  user.Name,
			Role:  user.Role,
		},
		AccessToken: token,
		ExpiresIn:   h.authService.AccessTTLSeconds(),
	})
}

func (h *AuthHandler) Refresh(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "missing refresh token"})
		return
	}

	accessToken, err := h.authService.RefreshToken(c.Request.Context(), refreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired refresh token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"accessToken": accessToken,
		"expiresIn":   h.authService.AccessTTLSeconds(),
	})
}

func (h *AuthHandler) OAuthProviders(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"github": h.authService.IsGitHubConfigured(),
		"google": h.authService.IsGoogleConfigured(),
	})
}

func (h *AuthHandler) GitHubRedirect(c *gin.Context) {
	if !h.authService.IsGitHubConfigured() {
		c.JSON(http.StatusNotImplemented, gin.H{"error": "GitHub OAuth is not configured"})
		return
	}
	c.Redirect(http.StatusFound, h.authService.GitHubAuthURL())
}

func (h *AuthHandler) GitHubCallback(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing code parameter"})
		return
	}

	user, token, err := h.authService.GitHubCallback(c.Request.Context(), code)
	if err != nil {
		c.Redirect(http.StatusFound, h.authService.FrontendURL()+"/login?error=github_failed")
		return
	}

	refreshToken, _ := h.authService.GenerateRefreshToken(user.ID.Hex())
	h.setRefreshCookie(c, refreshToken)

	c.Redirect(http.StatusFound, h.authService.FrontendURL()+"/auth/callback/github?access_token="+token)
}

func (h *AuthHandler) GoogleRedirect(c *gin.Context) {
	if !h.authService.IsGoogleConfigured() {
		c.JSON(http.StatusNotImplemented, gin.H{"error": "Google OAuth is not configured"})
		return
	}
	c.Redirect(http.StatusFound, h.authService.GoogleAuthURL())
}

func (h *AuthHandler) GoogleCallback(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing code parameter"})
		return
	}

	user, token, err := h.authService.GoogleCallback(c.Request.Context(), code)
	if err != nil {
		c.Redirect(http.StatusFound, h.authService.FrontendURL()+"/login?error=google_failed")
		return
	}

	refreshToken, _ := h.authService.GenerateRefreshToken(user.ID.Hex())
	h.setRefreshCookie(c, refreshToken)

	c.Redirect(http.StatusFound, h.authService.FrontendURL()+"/auth/callback/google?access_token="+token)
}

func (h *AuthHandler) Me(c *gin.Context) {
	userID, _ := c.Get("user_id")
	user, err := h.authService.GetCurrentUser(c.Request.Context(), userID.(string))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, models.UserProfileResponse{
		ID:        user.ID.Hex(),
		Email:     user.Email,
		Name:      user.Name,
		AvatarURL: user.AvatarURL,
		Role:      user.Role,
		Streak:    user.Streak,
	})
}

func (h *AuthHandler) setRefreshCookie(c *gin.Context, token string) {
	c.SetSameSite(http.SameSiteStrictMode)
	c.SetCookie(
		"refresh_token",
		token,
		int(h.authService.RefreshTTL().Seconds()),
		"/",
		"",
		true,
		true,
	)
}
