package middleware

import (
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type bucket struct {
	tokens    float64
	lastCheck time.Time
}

type RateLimiter struct {
	mu       sync.Mutex
	buckets  map[string]*bucket
	rate     float64
	capacity float64
	window   time.Duration
}

func NewRateLimiter(requestsPerMinute int) *RateLimiter {
	rl := &RateLimiter{
		buckets:  make(map[string]*bucket),
		rate:     float64(requestsPerMinute) / 60.0,
		capacity: float64(requestsPerMinute),
		window:   time.Minute,
	}
	go rl.cleanup()
	return rl
}

func (rl *RateLimiter) cleanup() {
	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()
	for range ticker.C {
		rl.mu.Lock()
		cutoff := time.Now().Add(-10 * time.Minute)
		for key, b := range rl.buckets {
			if b.lastCheck.Before(cutoff) {
				delete(rl.buckets, key)
			}
		}
		rl.mu.Unlock()
	}
}

func (rl *RateLimiter) allow(key string) (bool, float64) {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	b, exists := rl.buckets[key]
	if !exists {
		rl.buckets[key] = &bucket{tokens: rl.capacity - 1, lastCheck: now}
		return true, rl.capacity - 1
	}

	elapsed := now.Sub(b.lastCheck).Seconds()
	b.tokens += elapsed * rl.rate
	if b.tokens > rl.capacity {
		b.tokens = rl.capacity
	}
	b.lastCheck = now

	if b.tokens < 1 {
		retryAfter := (1 - b.tokens) / rl.rate
		return false, retryAfter
	}

	b.tokens--
	return true, b.tokens
}

func (rl *RateLimiter) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.ClientIP()
		if userID, exists := c.Get("user_id"); exists {
			key = userID.(string)
		}

		allowed, val := rl.allow(key)
		if !allowed {
			c.Header("Retry-After", strconv.FormatFloat(val, 'f', 0, 64))
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error":       "rate limit exceeded",
				"retry_after": val,
			})
			return
		}
		c.Next()
	}
}

var (
	AuthLimiter       = NewRateLimiter(10)
	TutorLimiter      = NewRateLimiter(30)
	PlaygroundLimiter = NewRateLimiter(20)
	GeneralLimiter    = NewRateLimiter(100)
)
