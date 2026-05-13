package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/mastery-hub/backend/internal/config"
	"github.com/mastery-hub/backend/internal/models"
	"github.com/mastery-hub/backend/internal/repository"
	"github.com/mastery-hub/backend/internal/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ErrUserExists       = errors.New("user already exists")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUserNotFound     = errors.New("user not found")
)

type AuthService struct {
	userRepo *repository.UserRepo
	cfg      *config.Config
}

func NewAuthService(userRepo *repository.UserRepo, cfg *config.Config) *AuthService {
	return &AuthService{userRepo: userRepo, cfg: cfg}
}

func (s *AuthService) Register(ctx context.Context, req models.RegisterRequest) (*models.User, string, error) {
	if !utils.ValidatePassword(req.Password) {
		return nil, "", errors.New("password must be at least 8 characters with uppercase, lowercase, and a number")
	}

	existing, _ := s.userRepo.FindByEmail(ctx, req.Email)
	if existing != nil {
		return nil, "", ErrUserExists
	}

	hash, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, "", fmt.Errorf("hashing password: %w", err)
	}

	user := &models.User{
		Email:        req.Email,
		Name:         req.Name,
		PasswordHash: hash,
		AuthProvider: "local",
		Role:         "learner",
		Streak:       models.Streak{},
	}

	if err := s.userRepo.Create(ctx, user); err != nil {
		return nil, "", fmt.Errorf("creating user: %w", err)
	}

	token, err := utils.GenerateAccessToken(user.ID.Hex(), user.Role, s.cfg.JWTSecret, s.cfg.JWTAccessTTL)
	if err != nil {
		return nil, "", fmt.Errorf("generating token: %w", err)
	}

	return user, token, nil
}

func (s *AuthService) Login(ctx context.Context, req models.LoginRequest) (*models.User, string, error) {
	user, err := s.userRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		return nil, "", ErrInvalidCredentials
	}

	if !utils.CheckPassword(req.Password, user.PasswordHash) {
		return nil, "", ErrInvalidCredentials
	}

	token, err := utils.GenerateAccessToken(user.ID.Hex(), user.Role, s.cfg.JWTSecret, s.cfg.JWTAccessTTL)
	if err != nil {
		return nil, "", fmt.Errorf("generating token: %w", err)
	}

	return user, token, nil
}

func (s *AuthService) GenerateRefreshToken(userID string) (string, error) {
	return utils.GenerateRefreshToken(userID, s.cfg.JWTSecret, s.cfg.JWTRefreshTTL)
}

func (s *AuthService) RefreshToken(ctx context.Context, refreshToken string) (string, error) {
	claims, err := utils.ValidateToken(refreshToken, s.cfg.JWTSecret)
	if err != nil {
		return "", err
	}

	oid, err := primitive.ObjectIDFromHex(claims.Subject)
	if err != nil {
		return "", ErrUserNotFound
	}

	user, err := s.userRepo.FindByID(ctx, oid)
	if err != nil {
		return "", ErrUserNotFound
	}

	return utils.GenerateAccessToken(user.ID.Hex(), user.Role, s.cfg.JWTSecret, s.cfg.JWTAccessTTL)
}

func (s *AuthService) GetCurrentUser(ctx context.Context, userID string) (*models.User, error) {
	oid, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, ErrUserNotFound
	}
	return s.userRepo.FindByID(ctx, oid)
}

func (s *AuthService) BackendURL() string {
	return fmt.Sprintf("http://localhost:%s", s.cfg.Port)
}

func (s *AuthService) FrontendURL() string {
	return s.cfg.FrontendURL
}

func (s *AuthService) GitHubAuthURL() string {
	params := url.Values{
		"client_id":    {s.cfg.GitHubClientID},
		"redirect_uri": {s.BackendURL() + "/api/auth/github/callback"},
		"scope":        {"user:email"},
	}
	return "https://github.com/login/oauth/authorize?" + params.Encode()
}

func (s *AuthService) GitHubCallback(ctx context.Context, code string) (*models.User, string, error) {
	tokenResp, err := s.exchangeGitHubCode(code)
	if err != nil {
		return nil, "", fmt.Errorf("github token exchange: %w", err)
	}

	ghUser, err := s.fetchGitHubUser(tokenResp.AccessToken)
	if err != nil {
		return nil, "", fmt.Errorf("github user fetch: %w", err)
	}

	user, err := s.userRepo.FindByProviderID(ctx, "github", fmt.Sprintf("%d", ghUser.ID))
	if errors.Is(err, mongo.ErrNoDocuments) {
		user = &models.User{
			Email:        ghUser.Email,
			Name:         ghUser.Name,
			AvatarURL:    ghUser.AvatarURL,
			AuthProvider: "github",
			ProviderID:   fmt.Sprintf("%d", ghUser.ID),
			Role:         "learner",
		}
		if user.Name == "" {
			user.Name = ghUser.Login
		}
		if err := s.userRepo.Create(ctx, user); err != nil {
			return nil, "", err
		}
	} else if err != nil {
		return nil, "", err
	} else {
		user.AvatarURL = ghUser.AvatarURL
		user.Name = ghUser.Name
		if user.Name == "" {
			user.Name = ghUser.Login
		}
		_ = s.userRepo.Update(ctx, user)
	}

	token, err := utils.GenerateAccessToken(user.ID.Hex(), user.Role, s.cfg.JWTSecret, s.cfg.JWTAccessTTL)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

type githubTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

type githubUser struct {
	ID        int    `json:"id"`
	Login     string `json:"login"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	AvatarURL string `json:"avatar_url"`
}

func (s *AuthService) exchangeGitHubCode(code string) (*githubTokenResponse, error) {
	data := url.Values{
		"client_id":     {s.cfg.GitHubClientID},
		"client_secret": {s.cfg.GitHubClientSecret},
		"code":          {code},
	}

	req, err := http.NewRequest("POST", "https://github.com/login/oauth/access_token", strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var tokenResp githubTokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
		return nil, err
	}
	return &tokenResp, nil
}

func (s *AuthService) fetchGitHubUser(accessToken string) (*githubUser, error) {
	req, err := http.NewRequest("GET", "https://api.github.com/user", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var user githubUser
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, err
	}

	if user.Email == "" {
		email, _ := s.fetchGitHubEmail(accessToken)
		user.Email = email
	}

	return &user, nil
}

func (s *AuthService) fetchGitHubEmail(accessToken string) (string, error) {
	req, err := http.NewRequest("GET", "https://api.github.com/user/emails", nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var emails []struct {
		Email    string `json:"email"`
		Primary  bool   `json:"primary"`
		Verified bool   `json:"verified"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&emails); err != nil {
		return "", err
	}

	for _, e := range emails {
		if e.Primary && e.Verified {
			return e.Email, nil
		}
	}
	if len(emails) > 0 {
		return emails[0].Email, nil
	}
	return "", errors.New("no email found")
}

func (s *AuthService) GoogleAuthURL() string {
	params := url.Values{
		"client_id":     {s.cfg.GoogleClientID},
		"redirect_uri":  {s.BackendURL() + "/api/auth/google/callback"},
		"response_type": {"code"},
		"scope":         {"openid email profile"},
	}
	return "https://accounts.google.com/o/oauth2/v2/auth?" + params.Encode()
}

func (s *AuthService) GoogleCallback(ctx context.Context, code string) (*models.User, string, error) {
	tokenResp, err := s.exchangeGoogleCode(code)
	if err != nil {
		return nil, "", fmt.Errorf("google token exchange: %w", err)
	}

	gUser, err := s.fetchGoogleUser(tokenResp.AccessToken)
	if err != nil {
		return nil, "", fmt.Errorf("google user fetch: %w", err)
	}

	user, err := s.userRepo.FindByProviderID(ctx, "google", gUser.Sub)
	if errors.Is(err, mongo.ErrNoDocuments) {
		user = &models.User{
			Email:        gUser.Email,
			Name:         gUser.Name,
			AvatarURL:    gUser.Picture,
			AuthProvider: "google",
			ProviderID:   gUser.Sub,
			Role:         "learner",
		}
		if err := s.userRepo.Create(ctx, user); err != nil {
			return nil, "", err
		}
	} else if err != nil {
		return nil, "", err
	} else {
		user.AvatarURL = gUser.Picture
		user.Name = gUser.Name
		_ = s.userRepo.Update(ctx, user)
	}

	token, err := utils.GenerateAccessToken(user.ID.Hex(), user.Role, s.cfg.JWTSecret, s.cfg.JWTAccessTTL)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

type googleTokenResponse struct {
	AccessToken string `json:"access_token"`
	IDToken     string `json:"id_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

type googleUser struct {
	Sub     string `json:"sub"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
}

func (s *AuthService) exchangeGoogleCode(code string) (*googleTokenResponse, error) {
	data := url.Values{
		"code":          {code},
		"client_id":     {s.cfg.GoogleClientID},
		"client_secret": {s.cfg.GoogleClientSecret},
		"redirect_uri":  {s.BackendURL() + "/api/auth/google/callback"},
		"grant_type":    {"authorization_code"},
	}

	resp, err := http.PostForm("https://oauth2.googleapis.com/token", data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var tokenResp googleTokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
		return nil, err
	}
	return &tokenResp, nil
}

func (s *AuthService) fetchGoogleUser(accessToken string) (*googleUser, error) {
	resp, err := http.Get("https://www.googleapis.com/oauth2/v3/userinfo?access_token=" + accessToken)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var user googleUser
	if err := json.Unmarshal(body, &user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *AuthService) IsGitHubConfigured() bool {
	return s.cfg.GitHubClientID != "" && s.cfg.GitHubClientSecret != ""
}

func (s *AuthService) IsGoogleConfigured() bool {
	return s.cfg.GoogleClientID != "" && s.cfg.GoogleClientSecret != ""
}

func (s *AuthService) AccessTTLSeconds() int {
	return int(s.cfg.JWTAccessTTL.Seconds())
}

func (s *AuthService) RefreshTTL() time.Duration {
	return s.cfg.JWTRefreshTTL
}
