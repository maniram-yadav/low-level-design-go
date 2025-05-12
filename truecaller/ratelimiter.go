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
