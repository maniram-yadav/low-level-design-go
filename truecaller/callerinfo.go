package truecaller

import "time"

type CallerInfo struct {
	PhoneNumber *PhoneNumber
	Name        string
	SpamScore   int
	IsSpam      bool
	IsContact   bool
	IsBlocked   bool
	Carrier     string
	Location    string
	Time        time.Time
}
