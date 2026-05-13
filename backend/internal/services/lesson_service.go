package services

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/mastery-hub/backend/internal/models"
	"github.com/mastery-hub/backend/internal/repository"
)

type LessonService struct {
	trackRepo  *repository.TrackRepo
	contentDir string
}

func NewLessonService(trackRepo *repository.TrackRepo, contentDir string) *LessonService {
	return &LessonService{trackRepo: trackRepo, contentDir: contentDir}
}

func (s *LessonService) GetLesson(ctx context.Context, trackSlug, moduleSlug, lessonSlug string) (*models.LessonDetail, error) {
	track, _, lesson, err := s.trackRepo.FindLessonBySlug(ctx, trackSlug, moduleSlug, lessonSlug)
	if err != nil {
		return nil, err
	}

	content, err := s.readContent(lesson.ContentPath)
	if err != nil {
		content = ""
	}

	prev, next := s.findAdjacentLessons(track, moduleSlug, lessonSlug)

	return &models.LessonDetail{
		Slug:             lesson.Slug,
		Title:            lesson.Title,
		TrackSlug:        trackSlug,
		ModuleSlug:       moduleSlug,
		EstimatedMinutes: lesson.EstimatedMinutes,
		HasQuiz:          lesson.HasQuiz,
		Content:          content,
		PrevLesson:       prev,
		NextLesson:       next,
	}, nil
}

func (s *LessonService) GetLessonBySlug(ctx context.Context, lessonSlug string) (*models.LessonDetail, error) {
	track, mod, lesson, err := s.trackRepo.FindLessonBySlugFlat(ctx, lessonSlug)
	if err != nil {
		return nil, err
	}

	content, err := s.readContent(lesson.ContentPath)
	if err != nil {
		content = ""
	}

	prev, next := s.findAdjacentLessons(track, mod.Slug, lessonSlug)

	return &models.LessonDetail{
		Slug:             lesson.Slug,
		Title:            lesson.Title,
		TrackSlug:        track.Slug,
		ModuleSlug:       mod.Slug,
		EstimatedMinutes: lesson.EstimatedMinutes,
		HasQuiz:          lesson.HasQuiz,
		Content:          content,
		PrevLesson:       prev,
		NextLesson:       next,
	}, nil
}

func (s *LessonService) GetQuiz(ctx context.Context, trackSlug, moduleSlug string) ([]models.QuizQuestion, error) {
	quizPath := filepath.Join(s.contentDir, trackSlug, moduleSlug, "quiz.json")
	data, err := os.ReadFile(quizPath)
	if err != nil {
		return nil, fmt.Errorf("quiz not found: %w", err)
	}

	var questions []models.QuizQuestion
	if err := json.Unmarshal(data, &questions); err != nil {
		return nil, fmt.Errorf("invalid quiz format: %w", err)
	}
	return questions, nil
}

func (s *LessonService) GradeQuiz(questions []models.QuizQuestion, answers []models.QuizAnswer) (int, int, []models.QuizResult) {
	answerMap := make(map[string]string, len(answers))
	for _, a := range answers {
		answerMap[a.QuestionID] = a.Selected
	}

	var score int
	results := make([]models.QuizResult, 0, len(questions))
	for _, q := range questions {
		selected := answerMap[q.ID]
		correct := selected == q.Answer
		if correct {
			score++
		}
		result := models.QuizResult{
			QuestionID: q.ID,
			Correct:    correct,
		}
		if !correct {
			result.CorrectAnswer = q.Answer
		}
		results = append(results, result)
	}
	return score, len(questions), results
}

func (s *LessonService) readContent(contentPath string) (string, error) {
	if contentPath == "" {
		return "", nil
	}
	fullPath := filepath.Join(s.contentDir, contentPath)
	data, err := os.ReadFile(fullPath)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (s *LessonService) findAdjacentLessons(track *models.Track, moduleSlug, lessonSlug string) (*models.LessonLink, *models.LessonLink) {
	type flatLesson struct {
		slug  string
		title string
	}

	var allLessons []flatLesson
	var currentIdx int = -1

	for _, mod := range track.Modules {
		for _, les := range mod.Lessons {
			if mod.Slug == moduleSlug && les.Slug == lessonSlug {
				currentIdx = len(allLessons)
			}
			allLessons = append(allLessons, flatLesson{slug: les.Slug, title: les.Title})
		}
	}

	var prev, next *models.LessonLink
	if currentIdx > 0 {
		prev = &models.LessonLink{Slug: allLessons[currentIdx-1].slug, Title: allLessons[currentIdx-1].title}
	}
	if currentIdx >= 0 && currentIdx < len(allLessons)-1 {
		next = &models.LessonLink{Slug: allLessons[currentIdx+1].slug, Title: allLessons[currentIdx+1].title}
	}
	return prev, next
}
