package services

import (
	"context"
	"errors"
	"math"
	"time"

	"github.com/mastery-hub/backend/internal/models"
	"github.com/mastery-hub/backend/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type FlashcardService struct {
	repo *repository.FlashcardRepo
}

func NewFlashcardService(repo *repository.FlashcardRepo) *FlashcardService {
	return &FlashcardService{repo: repo}
}

func (s *FlashcardService) GetByModule(ctx context.Context, moduleSlug string) ([]models.Flashcard, error) {
	return s.repo.FindByModule(ctx, moduleSlug)
}

func (s *FlashcardService) GetDueCards(ctx context.Context, userID string, limit int64) (*models.DueCardsResponse, error) {
	oid, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}

	dueCount, err := s.repo.CountDueForUser(ctx, oid)
	if err != nil {
		return nil, err
	}

	progressList, err := s.repo.FindDueForUser(ctx, oid, limit)
	if err != nil {
		return nil, err
	}

	cards := make([]models.FlashcardWithProgress, 0, len(progressList))
	for _, p := range progressList {
		card, err := s.repo.FindByID(ctx, p.FlashcardID)
		if err != nil {
			continue
		}
		cards = append(cards, models.FlashcardWithProgress{
			ID:           card.ID.Hex(),
			Question:     card.Question,
			Answer:       card.Answer,
			ModuleSlug:   card.ModuleSlug,
			Difficulty:   card.Difficulty,
			EaseFactor:   p.EaseFactor,
			IntervalDays: p.IntervalDays,
			Repetitions:  p.Repetitions,
		})
	}

	return &models.DueCardsResponse{
		DueCount: int(dueCount),
		Cards:    cards,
	}, nil
}

func (s *FlashcardService) ReviewCard(ctx context.Context, userID, flashcardID string, quality int) (*models.ReviewResponse, error) {
	uid, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}
	fid, err := primitive.ObjectIDFromHex(flashcardID)
	if err != nil {
		return nil, err
	}

	_, err = s.repo.FindByID(ctx, fid)
	if err != nil {
		return nil, errors.New("flashcard not found")
	}

	progress, err := s.repo.FindProgress(ctx, uid, fid)
	if errors.Is(err, mongo.ErrNoDocuments) {
		progress = &models.UserFlashcardProgress{
			UserID:       uid,
			FlashcardID:  fid,
			EaseFactor:   2.5,
			IntervalDays: 1,
			Repetitions:  0,
		}
	} else if err != nil {
		return nil, err
	}

	progress = sm2(progress, quality)

	if err := s.repo.UpsertProgress(ctx, progress); err != nil {
		return nil, err
	}

	return &models.ReviewResponse{
		FlashcardID:  flashcardID,
		EaseFactor:   progress.EaseFactor,
		IntervalDays: progress.IntervalDays,
		NextReview:   progress.NextReview,
		Repetitions:  progress.Repetitions,
	}, nil
}

func (s *FlashcardService) InitializeUserCards(ctx context.Context, userID string, cards []models.Flashcard) error {
	uid, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	now := time.Now()
	for _, card := range cards {
		p := &models.UserFlashcardProgress{
			UserID:       uid,
			FlashcardID:  card.ID,
			EaseFactor:   2.5,
			IntervalDays: 1,
			NextReview:   now,
			Repetitions:  0,
		}
		if err := s.repo.UpsertProgress(ctx, p); err != nil {
			return err
		}
	}
	return nil
}

// SM-2 algorithm: calculates the next review date based on quality (0-5).
// quality < 3: reset repetitions and start over.
// quality >= 3: increase interval based on ease factor.
func sm2(p *models.UserFlashcardProgress, quality int) *models.UserFlashcardProgress {
	if quality < 3 {
		p.Repetitions = 0
		p.IntervalDays = 1
	} else {
		switch p.Repetitions {
		case 0:
			p.IntervalDays = 1
		case 1:
			p.IntervalDays = 6
		default:
			p.IntervalDays = int(math.Round(float64(p.IntervalDays) * p.EaseFactor))
		}
		p.Repetitions++
	}

	p.EaseFactor = p.EaseFactor + (0.1 - float64(5-quality)*(0.08+float64(5-quality)*0.02))
	if p.EaseFactor < 1.3 {
		p.EaseFactor = 1.3
	}

	p.NextReview = time.Now().AddDate(0, 0, p.IntervalDays)
	return p
}
