package services

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"strings"
	"time"

	"github.com/mastery-hub/backend/internal/config"
)

type RunRequest struct {
	Language string `json:"language" binding:"required,oneof=go python javascript"`
	Code     string `json:"code" binding:"required"`
	Stdin    string `json:"stdin"`
}

type RunResponse struct {
	Stdout          string `json:"stdout"`
	Stderr          string `json:"stderr"`
	ExitCode        int    `json:"exitCode"`
	ExecutionTimeMs int64  `json:"executionTimeMs"`
}

var languageImages = map[string]string{
	"go":         "golang:1.24-alpine",
	"python":     "python:3.12-alpine",
	"javascript": "node:22-alpine",
}

var languageCommands = map[string][]string{
	"go":         {"sh", "-c", "echo \"$CODE\" > /tmp/main.go && cd /tmp && go run main.go"},
	"python":     {"sh", "-c", "echo \"$CODE\" | python3"},
	"javascript": {"sh", "-c", "echo \"$CODE\" | node"},
}

type PlaygroundService struct {
	cfg *config.Config
}

func NewPlaygroundService(cfg *config.Config) *PlaygroundService {
	return &PlaygroundService{cfg: cfg}
}

func (s *PlaygroundService) Run(ctx context.Context, req RunRequest) (*RunResponse, error) {
	image, ok := languageImages[req.Language]
	if !ok {
		return nil, fmt.Errorf("unsupported language: %s", req.Language)
	}

	if len(req.Code) > 64*1024 {
		return nil, fmt.Errorf("code exceeds maximum size of 64KB")
	}

	ctx, cancel := context.WithTimeout(ctx, s.cfg.SandboxTimeout)
	defer cancel()

	args := []string{
		"run", "--rm",
		"--network=none",
		"--memory=" + s.cfg.SandboxMemory,
		"--cpus=0.5",
		"--user=1000:1000",
		"--read-only",
		"--tmpfs=/tmp:rw,noexec,nosuid,size=64m",
		"-e", "CODE=" + req.Code,
	}

	if req.Stdin != "" {
		args = append(args, "-i")
	}

	args = append(args, image)
	args = append(args, languageCommands[req.Language]...)

	cmd := exec.CommandContext(ctx, "docker", args...)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if req.Stdin != "" {
		cmd.Stdin = strings.NewReader(req.Stdin)
	}

	start := time.Now()
	err := cmd.Run()
	elapsed := time.Since(start)

	exitCode := 0
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			exitCode = exitErr.ExitCode()
		} else if ctx.Err() == context.DeadlineExceeded {
			return &RunResponse{
				Stderr:          "execution timed out",
				ExitCode:        124,
				ExecutionTimeMs: elapsed.Milliseconds(),
			}, nil
		} else {
			return nil, fmt.Errorf("docker exec failed: %w", err)
		}
	}

	return &RunResponse{
		Stdout:          stdout.String(),
		Stderr:          stderr.String(),
		ExitCode:        exitCode,
		ExecutionTimeMs: elapsed.Milliseconds(),
	}, nil
}
