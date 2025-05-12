package truecaller

import (
	"sync"
	"time"
)

type RateLimiter struct {
	requests  int
	interval  time.Duration
	counts    map[string]int
	lastReset time.Time
	mu        sync.Mutex
}

func NewRateLimiter(requests int, interval time.Duration) *RateLimiter {
	return &RateLimiter{
		requests:  requests,
		interval:  interval,
		counts:    make(map[string]int),
		lastReset: time.Now(),
	}
}

func (rl *RateLimiter) Allow(userID string) bool {

	rl.mu.Lock()
	defer rl.mu.Unlock()

	if time.Since(rl.lastReset) > rl.interval {
		rl.counts = make(map[string]int)
		rl.lastReset = time.Now()
	}

	if rl.counts[userID] >= rl.requests {
		return false
	}

	rl.counts[userID]++
	return true

}
