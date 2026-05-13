package services

import (
	"context"
	"errors"
	"time"

	"github.com/mastery-hub/backend/internal/models"
	"github.com/mastery-hub/backend/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProgressService struct {
	progressRepo *repository.ProgressRepo
	userRepo     *repository.UserRepo
}

func NewProgressService(progressRepo *repository.ProgressRepo, userRepo *repository.UserRepo) *ProgressService {
	return &ProgressService{progressRepo: progressRepo, userRepo: userRepo}
}

func (s *ProgressService) GetUserProgress(ctx context.Context, userID string, trackSlug string) ([]models.Progress, error) {
	oid, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}
	return s.progressRepo.FindByUser(ctx, oid, trackSlug)
}

func (s *ProgressService) UpdateProgress(ctx context.Context, userID, lessonSlug string, req models.UpdateProgressRequest) (*models.Progress, error) {
	oid, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}

	existing, err := s.progressRepo.FindByUserAndLesson(ctx, oid, lessonSlug)
	if errors.Is(err, mongo.ErrNoDocuments) {
		existing = &models.Progress{
			UserID:     oid,
			LessonSlug: lessonSlug,
		}
	} else if err != nil {
		return nil, err
	}

	existing.Status = req.Status
	existing.TrackSlug = req.TrackSlug
	existing.ModuleSlug = req.ModuleSlug
	existing.TimeSpent += req.TimeSpentDelta

	if req.Status == "completed" && existing.CompletedAt == nil {
		now := time.Now()
		existing.CompletedAt = &now
		s.updateStreak(ctx, oid)
	}

	if err := s.progressRepo.Upsert(ctx, existing); err != nil {
		return nil, err
	}
	return existing, nil
}

func (s *ProgressService) SubmitQuiz(ctx context.Context, userID, lessonSlug string, score models.QuizScore) error {
	oid, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}
	return s.progressRepo.AddQuizScore(ctx, oid, lessonSlug, score)
}

func (s *ProgressService) GetStreaks(ctx context.Context, userID string) (*models.StreakResponse, error) {
	oid, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}

	user, err := s.userRepo.FindByID(ctx, oid)
	if err != nil {
		return nil, err
	}

	history := make([]models.StreakDayEntry, 0, 7)
	for i := 0; i < 7; i++ {
		date := time.Now().AddDate(0, 0, -i)
		count, _ := s.progressRepo.CountCompletedOnDate(ctx, oid, date)
		history = append(history, models.StreakDayEntry{
			Date:             date.Format("2006-01-02"),
			LessonsCompleted: int(count),
		})
	}

	return &models.StreakResponse{
		Current:      user.Streak.Current,
		Longest:      user.Streak.Longest,
		LastActivity: user.Streak.LastActivity,
		History:      history,
	}, nil
}

func (s *ProgressService) updateStreak(ctx context.Context, userID primitive.ObjectID) {
	user, err := s.userRepo.FindByID(ctx, userID)
	if err != nil {
		return
	}

	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	lastActivity := time.Date(
		user.Streak.LastActivity.Year(),
		user.Streak.LastActivity.Month(),
		user.Streak.LastActivity.Day(),
		0, 0, 0, 0, time.UTC,
	)

	daysSinceLast := int(today.Sub(lastActivity).Hours() / 24)

	switch {
	case daysSinceLast == 0:
		// Already active today
	case daysSinceLast == 1:
		user.Streak.Current++
	default:
		user.Streak.Current = 1
	}

	if user.Streak.Current > user.Streak.Longest {
		user.Streak.Longest = user.Streak.Current
	}
	user.Streak.LastActivity = now

	_ = s.userRepo.Update(ctx, user)
}
