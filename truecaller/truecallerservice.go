package truecaller

import "sync"

type TrueCallerService struct {
	db              Database
	userCache       sync.Map
	contactCache    sync.Map
	spamThreshold   int
	rateLimiter     *RateLimiter
	carrierService  CarrierService
	locationService LocationService
}
