package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mastery-hub/backend/internal/config"
	"github.com/mastery-hub/backend/internal/handlers"
	"github.com/mastery-hub/backend/internal/middleware"
	"github.com/mastery-hub/backend/internal/repository"
	"github.com/mastery-hub/backend/internal/services"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		slog.Error("failed to load config", "error", err)
		os.Exit(1)
	}

	setupLogger(cfg.LogLevel)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db, err := repository.Connect(ctx, cfg.MongoURI, cfg.MongoDBName)
	if err != nil {
		slog.Error("failed to connect to MongoDB", "error", err)
		os.Exit(1)
	}

	if err := repository.EnsureIndexes(ctx, db); err != nil {
		slog.Warn("failed to ensure indexes", "error", err)
	}

	userRepo := repository.NewUserRepo(db)
	trackRepo := repository.NewTrackRepo(db)
	progressRepo := repository.NewProgressRepo(db)
	bookmarkRepo := repository.NewBookmarkRepo(db)
	noteRepo := repository.NewNoteRepo(db)
	flashcardRepo := repository.NewFlashcardRepo(db)

	authService := services.NewAuthService(userRepo, cfg)
	lessonService := services.NewLessonService(trackRepo, cfg.ContentDir)
	progressService := services.NewProgressService(progressRepo, userRepo)
	searchService := services.NewSearchService(db)
	tutorService := services.NewTutorService(cfg)
	playgroundService := services.NewPlaygroundService(cfg)
	flashcardService := services.NewFlashcardService(flashcardRepo)

	authHandler := handlers.NewAuthHandler(authService)
	tracksHandler := handlers.NewTracksHandler(trackRepo)
	lessonsHandler := handlers.NewLessonsHandler(lessonService)
	progressHandler := handlers.NewProgressHandler(progressService, lessonService)
	bookmarksHandler := handlers.NewBookmarksHandler(bookmarkRepo)
	notesHandler := handlers.NewNotesHandler(noteRepo)
	searchHandler := handlers.NewSearchHandler(searchService)
	tutorHandler := handlers.NewTutorHandler(tutorService)
	playgroundHandler := handlers.NewPlaygroundHandler(playgroundService)
	flashcardsHandler := handlers.NewFlashcardsHandler(flashcardService)
	adminHandler := handlers.NewAdminHandler(trackRepo, flashcardRepo)

	router := setupRouter(cfg, db, authHandler, tracksHandler, lessonsHandler, progressHandler,
		bookmarksHandler, notesHandler, searchHandler, tutorHandler, playgroundHandler,
		flashcardsHandler, adminHandler)

	srv := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 60 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		slog.Info("server starting", "port", cfg.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("server failed", "error", err)
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	slog.Info("shutting down server...")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		slog.Error("server forced to shutdown", "error", err)
	}

	if client := db.Client(); client != nil {
		if err := client.Disconnect(shutdownCtx); err != nil {
			slog.Error("mongodb disconnect error", "error", err)
		}
	}

	slog.Info("server stopped")
}

func setupRouter(
	cfg *config.Config,
	db *mongo.Database,
	authHandler *handlers.AuthHandler,
	tracksHandler *handlers.TracksHandler,
	lessonsHandler *handlers.LessonsHandler,
	progressHandler *handlers.ProgressHandler,
	bookmarksHandler *handlers.BookmarksHandler,
	notesHandler *handlers.NotesHandler,
	searchHandler *handlers.SearchHandler,
	tutorHandler *handlers.TutorHandler,
	playgroundHandler *handlers.PlaygroundHandler,
	flashcardsHandler *handlers.FlashcardsHandler,
	adminHandler *handlers.AdminHandler,
) *gin.Engine {
	router := gin.New()

	logger := slog.Default()
	router.Use(middleware.Logging(logger))
	router.Use(gin.Recovery())
	router.Use(middleware.CORS(cfg.FrontendURL))
	router.Use(middleware.GeneralLimiter.Middleware())

	router.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	router.GET("/readyz", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), 2*time.Second)
		defer cancel()
		if err := db.Client().Ping(ctx, nil); err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{"status": "not ready", "error": "mongodb unreachable"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "ready"})
	})

	api := router.Group("/api")

	auth := api.Group("/auth")
	auth.Use(middleware.AuthLimiter.Middleware())
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
		auth.POST("/refresh", authHandler.Refresh)
		auth.GET("/providers", authHandler.OAuthProviders)
		auth.GET("/github", authHandler.GitHubRedirect)
		auth.GET("/github/callback", authHandler.GitHubCallback)
		auth.GET("/google", authHandler.GoogleRedirect)
		auth.GET("/google/callback", authHandler.GoogleCallback)
		auth.GET("/me", middleware.RequireAuth(cfg.JWTSecret), authHandler.Me)
	}

	tracks := api.Group("/tracks")
	{
		tracks.GET("", tracksHandler.List)
		tracks.GET("/:slug", tracksHandler.GetBySlug)
	}

	lessons := api.Group("/lessons")
	{
		lessons.GET("/:slug", lessonsHandler.GetLessonBySlug)
	}

	content := api.Group("/content")
	{
		content.GET("/:trackSlug/:moduleSlug/lessons/:lessonSlug", lessonsHandler.GetLesson)
		content.GET("/:trackSlug/:moduleSlug/quiz", lessonsHandler.GetQuiz)
	}

	progress := api.Group("/progress")
	progress.Use(middleware.RequireAuth(cfg.JWTSecret))
	{
		progress.GET("", progressHandler.GetUserProgress)
		progress.PUT("/:lessonSlug", progressHandler.UpdateProgress)
		progress.POST("/:lessonSlug/quiz", progressHandler.SubmitQuiz)
		progress.GET("/streaks", progressHandler.GetStreaks)
	}

	bookmarks := api.Group("/bookmarks")
	bookmarks.Use(middleware.RequireAuth(cfg.JWTSecret))
	{
		bookmarks.GET("", bookmarksHandler.List)
		bookmarks.POST("", bookmarksHandler.Create)
		bookmarks.DELETE("/:id", bookmarksHandler.Delete)
	}

	notes := api.Group("/notes")
	notes.Use(middleware.RequireAuth(cfg.JWTSecret))
	{
		notes.GET("", notesHandler.List)
		notes.POST("", notesHandler.Create)
		notes.PUT("/:id", notesHandler.Update)
		notes.DELETE("/:id", notesHandler.Delete)
	}

	search := api.Group("/search")
	{
		search.GET("", searchHandler.Search)
	}

	tutor := api.Group("/tutor")
	tutor.Use(middleware.RequireAuth(cfg.JWTSecret))
	tutor.Use(middleware.TutorLimiter.Middleware())
	{
		tutor.POST("/chat", tutorHandler.Chat)
	}

	playground := api.Group("/playground")
	playground.Use(middleware.RequireAuth(cfg.JWTSecret))
	playground.Use(middleware.PlaygroundLimiter.Middleware())
	{
		playground.POST("/run", playgroundHandler.Run)
	}

	flashcards := api.Group("/flashcards")
	flashcards.Use(middleware.RequireAuth(cfg.JWTSecret))
	{
		flashcards.GET("", flashcardsHandler.List)
		flashcards.GET("/review", flashcardsHandler.GetDueCards)
		flashcards.POST("/:id/review", flashcardsHandler.Review)
	}

	admin := api.Group("/admin")
	admin.Use(middleware.RequireAdmin(cfg.JWTSecret))
	{
		admin.GET("/lessons", adminHandler.ListLessons)
		admin.POST("/lessons", adminHandler.CreateLesson)
		admin.PUT("/lessons/:slug", adminHandler.UpdateLesson)
		admin.DELETE("/lessons/:slug", adminHandler.DeleteLesson)
		admin.POST("/seed", adminHandler.Seed)
	}

	return router
}

func setupLogger(level string) {
	var logLevel slog.Level
	switch level {
	case "debug":
		logLevel = slog.LevelDebug
	case "warn":
		logLevel = slog.LevelWarn
	case "error":
		logLevel = slog.LevelError
	default:
		logLevel = slog.LevelInfo
	}

	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: logLevel})
	slog.SetDefault(slog.New(handler))
}
