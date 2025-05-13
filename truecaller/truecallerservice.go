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

func (t *TrueCallerService) IdentifyCaller(ctx context.Context, userID string, callerNumber *PhoneNumber) (*CallerInfo, error) {
	return nil,nil
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
func (t *TrueCallerService) ReportSpam(ctx context.Context, userID string, phone *PhoneNumber) error {
	return nil
}

func (t *TrueCallerService) BlockNumber(ctx context.Context, userID string, phone *PhoneNumber) error {
	return nil
}


func (t *TrueCallerService) getContact(ctx context.Context, phone *PhoneNumber) (*Contact, error) {
	
	phoneNumber := phone.String()
	if cached,ok := t.contactCache.Load(phoneNumber); ok {
		return cached.(*Contact),nil
	}  

	contact , err := t.db.GetContact(ctx,phone)
	if err != nil {
		return nil,err
	}

	t.contactCache.Store(phoneNumber,contact)
	return contact,nil	
}


func (t *TrueCallerService) SearchNumber(ctx context.Context, userID string, phone *PhoneNumber) (*CallerInfo, error) { 
	
	if !t.rateLimiter.Allow(userID) {
		return nil,error.new("rate limit exceeded")
	}
	
	return t.IdentifyCaller(ctx,userID,phone)

}

func (t *TrueCallerService) SearchName(ctx context.Context, userID, name string) ([]*CallerInfo, error) {
	return nil,nil
}

