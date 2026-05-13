package middleware

import (
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
	"crypto/rand"
	"encoding/hex"
)

func generateRequestID() string {
	b := make([]byte, 8)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}

func Logging(logger *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := generateRequestID()
		c.Set("request_id", requestID)
		c.Header("X-Request-ID", requestID)

		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method

		c.Next()

		duration := time.Since(start)
		status := c.Writer.Status()

		logAttrs := []slog.Attr{
			slog.String("request_id", requestID),
			slog.String("method", method),
			slog.String("path", path),
			slog.Int("status", status),
			slog.Duration("duration", duration),
			slog.String("client_ip", c.ClientIP()),
		}

		msg := "request completed"
		switch {
		case status >= 500:
			logger.LogAttrs(c.Request.Context(), slog.LevelError, msg, logAttrs...)
		case status >= 400:
			logger.LogAttrs(c.Request.Context(), slog.LevelWarn, msg, logAttrs...)
		default:
			logger.LogAttrs(c.Request.Context(), slog.LevelInfo, msg, logAttrs...)
		}
	}
}
