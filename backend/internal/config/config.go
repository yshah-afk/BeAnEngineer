package config

import (
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoURI   string
	MongoDBName string
	Port       string

	JWTSecret     string
	JWTAccessTTL  time.Duration
	JWTRefreshTTL time.Duration

	GitHubClientID     string
	GitHubClientSecret string
	GoogleClientID     string
	GoogleClientSecret string

	FrontendURL string

	OllamaBaseURL string
	OllamaModel   string
	OpenAIAPIKey  string
	OpenAIBaseURL string

	SandboxTimeout time.Duration
	SandboxMemory  string
	LogLevel       string
	ContentDir     string
}

func Load() (*Config, error) {
	_ = godotenv.Load()

	cfg := &Config{
		MongoURI:           getEnv("MONGO_URI", "mongodb://localhost:27017"),
		MongoDBName:        getEnv("MONGO_DB_NAME", "mastery_hub"),
		Port:               getEnv("PORT", "8080"),
		JWTSecret:          getEnv("JWT_SECRET", ""),
		GitHubClientID:     getEnv("GITHUB_CLIENT_ID", ""),
		GitHubClientSecret: getEnv("GITHUB_CLIENT_SECRET", ""),
		GoogleClientID:     getEnv("GOOGLE_CLIENT_ID", ""),
		GoogleClientSecret: getEnv("GOOGLE_CLIENT_SECRET", ""),
		FrontendURL:        getEnv("FRONTEND_URL", "http://localhost:5173"),
		OllamaBaseURL:      getEnv("OLLAMA_BASE_URL", "http://localhost:11434"),
		OllamaModel:        getEnv("OLLAMA_MODEL", "llama3"),
		OpenAIAPIKey:       getEnv("OPENAI_API_KEY", ""),
		OpenAIBaseURL:      getEnv("OPENAI_BASE_URL", "https://api.openai.com/v1"),
		SandboxMemory:      getEnv("SANDBOX_MEMORY", "128m"),
		LogLevel:           getEnv("LOG_LEVEL", "info"),
		ContentDir:         getEnv("CONTENT_DIR", "../content"),
	}

	var err error
	cfg.JWTAccessTTL, err = time.ParseDuration(getEnv("JWT_ACCESS_TTL", "15m"))
	if err != nil {
		cfg.JWTAccessTTL = 15 * time.Minute
	}

	cfg.JWTRefreshTTL, err = time.ParseDuration(getEnv("JWT_REFRESH_TTL", "168h"))
	if err != nil {
		cfg.JWTRefreshTTL = 168 * time.Hour
	}

	cfg.SandboxTimeout, err = time.ParseDuration(getEnv("SANDBOX_TIMEOUT", "30s"))
	if err != nil {
		cfg.SandboxTimeout = 30 * time.Second
	}

	return cfg, nil
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
