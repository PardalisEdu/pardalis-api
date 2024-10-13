package middleware

import (
	"net/http"
	"sync"
	"time"
)

type RateLimiter struct {
	rate   time.Duration
	burst  int
	mutex  sync.Mutex
	tokens map[string][]time.Time
}

func NewRateLimiter(rate time.Duration, burst int) *RateLimiter {
	return &RateLimiter{
		rate:   rate,
		burst:  burst,
		tokens: make(map[string][]time.Time),
	}
}

func (rl *RateLimiter) Allow(key string) bool {
	rl.mutex.Lock()
	defer rl.mutex.Unlock()

	now := time.Now()
	if _, exists := rl.tokens[key]; !exists {
		rl.tokens[key] = []time.Time{}
	}

	// Remove expired tokens
	var valid []time.Time
	for _, t := range rl.tokens[key] {
		if now.Sub(t) <= rl.rate*time.Duration(rl.burst) {
			valid = append(valid, t)
		}
	}
	rl.tokens[key] = valid

	if len(rl.tokens[key]) < rl.burst {
		rl.tokens[key] = append(rl.tokens[key], now)
		return true
	}

	oldestToken := rl.tokens[key][0]
	if now.Sub(oldestToken) >= rl.rate {
		rl.tokens[key] = append(rl.tokens[key][1:], now)
		return true
	}

	return false
}

func (rl *RateLimiter) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Use IP address as the rate limiting key
		key := r.RemoteAddr

		if !rl.Allow(key) {
			http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}
