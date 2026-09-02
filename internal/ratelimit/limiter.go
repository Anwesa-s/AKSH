package ratelimit

import (
	"sync"
	"time"
)

type Client struct {
	Tokens     float64
	LastRefill time.Time
}

type Limiter struct {
	mu         sync.Mutex
	clients    map[string]*Client
	rate       float64
	burst      float64
}

func NewLimiter(rate float64, burst float64) *Limiter {
	return &Limiter{
		clients: make(map[string]*Client),
		rate:    rate,
		burst:   burst,
	}
}

func (l *Limiter) Allow(key string) bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	now := time.Now()

	client, exists := l.clients[key]

	if !exists {
		l.clients[key] = &Client{
			Tokens:     l.burst - 1,
			LastRefill: now,
		}
		return true
	}

	elapsed := now.Sub(client.LastRefill).Seconds()

	client.Tokens += elapsed * l.rate

	if client.Tokens > l.burst {
		client.Tokens = l.burst
	}

	client.LastRefill = now

	if client.Tokens < 1 {
		return false
	}

	client.Tokens--

	return true
}