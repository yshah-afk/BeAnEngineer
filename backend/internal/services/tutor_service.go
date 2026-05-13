package services

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/mastery-hub/backend/internal/config"
)

type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type TutorRequest struct {
	Message    string        `json:"message" binding:"required"`
	LessonSlug string        `json:"lessonSlug"`
	History    []ChatMessage `json:"history"`
}

func (r TutorRequest) GetLessonSlug() string {
	return r.LessonSlug
}

type TutorService struct {
	cfg *config.Config
}

func NewTutorService(cfg *config.Config) *TutorService {
	return &TutorService{cfg: cfg}
}

func (s *TutorService) StreamChat(ctx context.Context, req TutorRequest, writer func(token string) error, done func(promptTokens, completionTokens int)) error {
	if s.cfg.OpenAIAPIKey != "" {
		return s.streamOpenAI(ctx, req, writer, done)
	}
	return s.streamOllama(ctx, req, writer, done)
}

func (s *TutorService) buildMessages(req TutorRequest) []ChatMessage {
	systemPrompt := "You are an expert programming tutor for the AI & Full-Stack Mastery Hub learning platform. " +
		"Provide clear, concise explanations with code examples when helpful. " +
		"Focus on building understanding, not just giving answers."

	if req.LessonSlug != "" {
		systemPrompt += fmt.Sprintf(" The student is currently studying lesson '%s'.", req.LessonSlug)
	}

	messages := []ChatMessage{{Role: "system", Content: systemPrompt}}
	messages = append(messages, req.History...)
	messages = append(messages, ChatMessage{Role: "user", Content: req.Message})
	return messages
}

func (s *TutorService) streamOllama(ctx context.Context, req TutorRequest, writer func(string) error, done func(int, int)) error {
	messages := s.buildMessages(req)

	body, err := json.Marshal(map[string]interface{}{
		"model":    s.cfg.OllamaModel,
		"messages": messages,
		"stream":   true,
	})
	if err != nil {
		return err
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", s.cfg.OllamaBaseURL+"/api/chat", bytes.NewReader(body))
	if err != nil {
		return err
	}
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return fmt.Errorf("ollama unavailable: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		respBody, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("ollama error %d: %s", resp.StatusCode, string(respBody))
	}

	scanner := bufio.NewScanner(resp.Body)
	var totalTokens int
	for scanner.Scan() {
		var chunk struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
			Done bool `json:"done"`
		}
		if err := json.Unmarshal(scanner.Bytes(), &chunk); err != nil {
			continue
		}
		if chunk.Message.Content != "" {
			totalTokens++
			if err := writer(chunk.Message.Content); err != nil {
				return err
			}
		}
		if chunk.Done {
			break
		}
	}

	done(0, totalTokens)
	return scanner.Err()
}

func (s *TutorService) streamOpenAI(ctx context.Context, req TutorRequest, writer func(string) error, done func(int, int)) error {
	messages := s.buildMessages(req)

	body, err := json.Marshal(map[string]interface{}{
		"model":    s.cfg.OllamaModel,
		"messages": messages,
		"stream":   true,
	})
	if err != nil {
		return err
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", s.cfg.OpenAIBaseURL+"/chat/completions", bytes.NewReader(body))
	if err != nil {
		return err
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Bearer "+s.cfg.OpenAIAPIKey)

	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return fmt.Errorf("openai unavailable: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		respBody, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("openai error %d: %s", resp.StatusCode, string(respBody))
	}

	scanner := bufio.NewScanner(resp.Body)
	var promptTokens, completionTokens int
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || line == "data: [DONE]" {
			continue
		}
		if len(line) > 6 && line[:6] == "data: " {
			line = line[6:]
		}

		var chunk struct {
			Choices []struct {
				Delta struct {
					Content string `json:"content"`
				} `json:"delta"`
			} `json:"choices"`
			Usage *struct {
				PromptTokens     int `json:"prompt_tokens"`
				CompletionTokens int `json:"completion_tokens"`
			} `json:"usage"`
		}
		if err := json.Unmarshal([]byte(line), &chunk); err != nil {
			continue
		}

		if len(chunk.Choices) > 0 && chunk.Choices[0].Delta.Content != "" {
			completionTokens++
			if err := writer(chunk.Choices[0].Delta.Content); err != nil {
				return err
			}
		}
		if chunk.Usage != nil {
			promptTokens = chunk.Usage.PromptTokens
			completionTokens = chunk.Usage.CompletionTokens
		}
	}

	done(promptTokens, completionTokens)
	return scanner.Err()
}
