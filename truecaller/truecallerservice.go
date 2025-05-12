package truecaller

import (
	"context"
	"sync"
)

type TrueCallerService struct {
	db              Database
	userCache       sync.Map
	contactCache    sync.Map
	spamThreshold   int
	rateLimiter     *RateLimiter
	carrierService  CarrierService
	locationService LocationService
}


func (t *TrueCallerService) AddContact(ctx context.Context,userId string,contact *Contact) error {
	user,err := t.getUser(ctx,userId)
	if err != nil {
		return err
	}
	user.Lock()
	defer user.Unlock()

	for _,phone := contact.PhoneNumbers {
		if _,exists := user.Contacts[phone.String()];exists {
			return ErrContactExists
		}
	}
	for _,phone := range contact.PhoneNumbers {
		 user.Contacts[phone.String()] = contact
	}
	return t.db.SaveUser(ctx,user)
}


func (t *TrueCallerService) getUser(ctx context.Context, userID string) (*User, error) {
	if cached, ok := t.userCache.Load(userID); ok {
		return cached.(*User), nil
	}
	user, err := t.db.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	t.userCache.Store(userID, user)
	return user, nil
}
