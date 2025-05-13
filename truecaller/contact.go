package truecaller

import (
	"sync"
	"time"
)

type Contact struct {
	Name         string
	PhoneNumbers []*PhoneNumber
	Email        string
	IsSpam       bool
	SpamCount    int
	LastUpdated  time.Time
	sync.Mutex
}
