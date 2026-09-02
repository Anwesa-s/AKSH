package ratelimit

import (
	"testing"
	"time"
)

func TestLimiter(t *testing.T) {
	limiter := NewLimiter(1, 3)

	key := "127.0.0.1"

	// First 3 requests should be allowed.
	for i := 0; i < 3; i++ {
		if !limiter.Allow(key) {
			t.Fatalf("request %d should be allowed", i+1)
		}
	}

	// 4th request should be rejected.
	if limiter.Allow(key) {
		t.Fatal("4th request should be rejected")
	}

	// Wait for one token to refill.
	time.Sleep(1100 * time.Millisecond)

	// Now one request should be allowed again.
	if !limiter.Allow(key) {
		t.Fatal("request should be allowed after token refill")
	}
}