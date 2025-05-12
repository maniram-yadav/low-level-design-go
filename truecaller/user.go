package truecaller

import (
	"sync"
	"time"
)

type User struct {
	ID          string
	Name        string
	PhoneNumber *PhoneNumber
	Contacts    map[string]*Contact
	Blocked     map[string]bool
	SpamReports map[string]bool
	AccessToken string
	TokenExpiry time.Time
	sync.RWMutex
}
