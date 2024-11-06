package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRateLimiter_Allow(t *testing.T) {
	limiter := NewRateLimiter(time.Second, 2)

	tests := []struct {
		name     string
		key      string
		attempts int
		want     bool
	}{
		{
			name:     "Under limit",
			key:      "test1",
			attempts: 1,
			want:     true,
		},
		{
			name:     "At limit",
			key:      "test2",
			attempts: 2,
			want:     true,
		},
		{
			name:     "Over limit",
			key:      "test3",
			attempts: 3,
			want:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var allowed bool
			for i := 0; i < tt.attempts; i++ {
				allowed = limiter.Allow(tt.key)
			}
			if allowed != tt.want {
				t.Errorf("RateLimiter.Allow() = %v, want %v", allowed, tt.want)
			}
		})
	}
}

func TestRateLimiter_Middleware(t *testing.T) {
	limiter := NewRateLimiter(time.Second, 2)
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	middleware := limiter.Middleware(handler)

	tests := []struct {
		name           string
		attempts       int
		wantLastStatus int
	}{
		{
			name:           "Under limit",
			attempts:       1,
			wantLastStatus: http.StatusOK,
		},
		{
			name:           "At limit",
			attempts:       2,
			wantLastStatus: http.StatusOK,
		},
		{
			name:           "Over limit",
			attempts:       3,
			wantLastStatus: http.StatusTooManyRequests,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var lastStatus int
			for i := 0; i < tt.attempts; i++ {
				req := httptest.NewRequest("GET", "/", nil)
				rec := httptest.NewRecorder()
				middleware.ServeHTTP(rec, req)
				lastStatus = rec.Code
			}
			if lastStatus != tt.wantLastStatus {
				t.Errorf("Last status = %v, want %v", lastStatus, tt.wantLastStatus)
			}
		})
	}
}
