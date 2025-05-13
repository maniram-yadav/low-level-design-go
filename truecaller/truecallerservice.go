package truecaller

import (
	"context"
	"errors"
	"sync"
	"testing"
	"time"
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

	if allow := t.rateLimiter.Allow(userID); !allow {
		return nil, errors.New("rate limit exceeded")
	}

	user, err := t.getUser(ctx, userID)
	if err != nil {
		return nil, err
	}
	user.RLock()
	_, isBlocked := user.Blocked[callerNumber.String()]
	user.RUnlock()
	if isBlocked {
		return &CallerInfo{
			PhoneNumber: callerNumber,
			IsBlocked:   true,
			Time:        time.Now(),
		}, nil
	}
	user.RLock()
	contact, isContact := user.Contacts[callerNumber.String()]
	user.RUnlock()
	if isContact {

		carrier, _ := t.carrierService.GetCarrier(callerNumber)
		location, _ := t.locationService.GetLocation(callerNumber)
		return &CallerInfo{
			PhoneNumber: callerNumber,
			Name:        contact.Name,
			IsContact:   true,
			IsSpam:      contact.IsSpam,
			SpamScore:   contact.SpamCount,
			Carrier:     carrier,
			Location:    location,
			Time:        time.Now(),
		}, nil
	}
	globalContact, err := t.getContact(ctx, callerNumber)
	if err != nil && !errors.Is(err, ErrNumberNotFound) {
		return nil, err
	}

	spamCount, err := t.db.GetSpamCount(ctx, callerNumber)
	if err != nil && !errors.Is(err, ErrNumberNotFound) {
		return nil, err
	}

	carrier, err := t.carrierService.GetCarrier(callerNumber)
	location, err := t.locationService.GetLocation(callerNumber)

	return &CallerInfo{
		PhoneNumber: callerNumber,
		Name: func() string {
			if globalContact != nil {
				return globalContact.Name
			}
			return ""
		}(),
		IsSpam:    spamCount >= t.spamThreshold,
		SpamScore: spamCount,
		Carrier:   carrier,
		Location:  location,
		Time:      time.Now(),
	}, nil

}

func (t *TrueCallerService) AddContact(ctx context.Context, userId string, contact *Contact) error {
	user, err := t.getUser(ctx, userId)
	if err != nil {
		return err
	}
	user.Lock()
	defer user.Unlock()

	for _, phone := range contact.PhoneNumbers {
		if _, exists := user.Contacts[phone.String()]; exists {
			return ErrContactExists
		}
	}
	for _, phone := range contact.PhoneNumbers {
		user.Contacts[phone.String()] = contact
	}
	return t.db.SaveUser(ctx, user)
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

	user, err := t.getUser(ctx, userID)

	if err != nil {
		return err
	}

	user.Lock()
	if _, exists := user.SpamReports[phone.String()]; exists {
		user.Unlock()
		return ErrSpamReportExists
	}

	user.SpamReports[phone.String()] = true
	user.Unlock()

	if err := t.db.IncrementSpamCount(ctx, phone); err != nil {
		return err
	}

	user.RLock()
	contact, exist := user.Contacts[phone.String()]
	user.RUnlock()

	if exist {
		contact.Lock()
		contact.IsSpam = true
		contact.SpamCount++
		contact.LastUpdated = time.Now()
		contact.Unlock()
		return t.db.SaveContact(ctx, contact)
	}

	return nil
}

func (t *TrueCallerService) BlockNumber(ctx context.Context, userID string, phone *PhoneNumber) error {

	user, err := t.getUser(ctx, userID)

	if err != nil {
		return errors.New("user not exists")
	}

	user.Lock()
	user.Blocked[phone.String()] = true
	user.Unlock()

	return t.db.SaveUser(ctx, user)

}

func (t *TrueCallerService) getContact(ctx context.Context, phone *PhoneNumber) (*Contact, error) {

	phoneNumber := phone.String()
	if cached, ok := t.contactCache.Load(phoneNumber); ok {
		return cached.(*Contact), nil
	}

	contact, err := t.db.GetContact(ctx, phone)
	if err != nil {
		return nil, err
	}

	t.contactCache.Store(phoneNumber, contact)
	return contact, nil
}

func (t *TrueCallerService) SearchNumber(ctx context.Context, userID string, phone *PhoneNumber) (*CallerInfo, error) {

	if !t.rateLimiter.Allow(userID) {
		return nil, errors.New("rate limit exceeded")
	}

	return t.IdentifyCaller(ctx, userID, phone)

}

func (t *TrueCallerService) SearchName(ctx context.Context, userID, name string) ([]*CallerInfo, error) {

	if !t.rateLimiter.Allow(userID) {
		return nil, errors.New("rate limit exceeded")
	}

	user, err := t.getUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	var results []*CallerInfo
	user.Lock()
	defer user.Unlock()

	for _, contact := range user.Contacts {
		if contact.Name == name {
			for _, phone := range contact.PhoneNumbers {
				carrier, _ := t.carrierService.GetCarrier(phone)
				location, _ := t.locationService.GetLocation(phone)

				results = append(results, &CallerInfo{
					PhoneNumber: phone,
					Name:        name,
					IsContact:   true,
					IsSpam:      contact.IsSpam,
					SpamScore:   contact.SpamCount,
					Carrier:     carrier,
					Location:    location,
				})
			}
		}
	}

	return results, nil
}

func NewTrueCallerServiceWithInMemoryDB() *TrueCallerService {
	db := NewInMemoryDB()

	ctx := context.Background()

	user1 := &User{
		ID:          "user1",
		Name:        "Alice",
		PhoneNumber: &PhoneNumber{countryCode: "+1", number: "1234567890"},
		Contacts:    make(map[string]*Contact),
		Blocked:     make(map[string]bool),
		SpamReports: make(map[string]bool),
	}
	db.SaveUser(ctx, user1)

	user2 := &User{
		ID:          "user2",
		Name:        "Bob",
		PhoneNumber: &PhoneNumber{countryCode: "+44", number: "9876543210"},
		Contacts:    make(map[string]*Contact),
		Blocked:     make(map[string]bool),
		SpamReports: make(map[string]bool),
	}
	db.SaveUser(ctx, user2)

	contact1 := &Contact{
		Name: "Charlie",
		PhoneNumbers: []*PhoneNumber{
			{countryCode: "+91", number: "5551234567"},
		},
		Email: "charlie@example.com",
	}
	db.SaveContact(ctx, contact1)

	contact2 := &Contact{
		Name: "David",
		PhoneNumbers: []*PhoneNumber{
			{countryCode: "+1", number: "5559876543"},
		},
		Email: "david@example.com",
	}
	db.SaveContact(ctx, contact2)

	db.IncrementSpamCount(ctx, &PhoneNumber{countryCode: "+1", number: "9998887777"})
	db.IncrementSpamCount(ctx, &PhoneNumber{countryCode: "+1", number: "9998887777"})
	db.IncrementSpamCount(ctx, &PhoneNumber{countryCode: "+1", number: "9998887777"})

	return &TrueCallerService{
		db:              db,
		spamThreshold:   3,
		rateLimiter:     NewRateLimiter(100, time.Minute),
		carrierService:  &CustomCarrierService{},
		locationService: &CusomLocationService{},
	}
}

func TestSearchNumber(t *testing.T) {
	service := NewTrueCallerServiceWithInMemoryDB()
	ctx := context.Background()

	tests := []struct {
		name     string
		userID   string
		phone    *PhoneNumber
		wantName string
		wantErr  bool
	}{
		{
			name:     "Search known contact",
			userID:   "user1",
			phone:    &PhoneNumber{countryCode: "+91", number: "5551234567"},
			wantName: "Charlie",
			wantErr:  false,
		},
		{
			name:     "Search unknown number",
			userID:   "user1",
			phone:    &PhoneNumber{countryCode: "+1", number: "1231231234"},
			wantName: "",
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := service.SearchNumber(ctx, tt.userID, tt.phone)
			if (err != nil) != tt.wantErr {
				t.Errorf("SearchNumber() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got.Name != tt.wantName {
				t.Errorf("SearchNumber() name = %v, want %v", got.Name, tt.wantName)
			}
		})
	}
}

func TestSearchName(t *testing.T) {
	service := NewTrueCallerServiceWithInMemoryDB()
	ctx := context.Background()

	// Add some more contacts for testing
	service.AddContact(ctx, "user1", &Contact{
		Name: "Alice",
		PhoneNumbers: []*PhoneNumber{
			{countryCode: "+1", number: "1112223333"},
		},
	})
	service.AddContact(ctx, "user1", &Contact{
		Name: "Alice",
		PhoneNumbers: []*PhoneNumber{
			{countryCode: "+1", number: "4445556666"},
		},
	})

	tests := []struct {
		name       string
		userID     string
		searchName string
		wantCount  int
		wantErr    bool
	}{
		{
			name:       "Search single match",
			userID:     "user1",
			searchName: "Charlie",
			wantCount:  1,
			wantErr:    false,
		},
		{
			name:       "Search multiple matches",
			userID:     "user1",
			searchName: "Alice",
			wantCount:  2,
			wantErr:    false,
		},
		{
			name:       "Search no matches",
			userID:     "user1",
			searchName: "Nonexistent",
			wantCount:  0,
			wantErr:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			results, err := service.SearchName(ctx, tt.userID, tt.searchName)
			if (err != nil) != tt.wantErr {
				t.Errorf("SearchName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(results) != tt.wantCount {
				t.Errorf("SearchName() count = %d, want %d", len(results), tt.wantCount)
			}
		})
	}
}

func TestRateLimiting(t *testing.T) {
	service := NewTrueCallerServiceWithInMemoryDB()
	ctx := context.Background()
	phone := &PhoneNumber{countryCode: "+1", number: "1234567890"}

	for i := 0; i < 100; i++ {
		_, err := service.IdentifyCaller(ctx, "user1", phone)
		if err != nil {
			t.Errorf("RateLimiter failed too early on request %d: %v", i+1, err)
		}
	}

	_, err := service.IdentifyCaller(ctx, "user1", phone)
	if err == nil {
		t.Error("RateLimiter should have blocked the 101st request")
	}

	_, err = service.IdentifyCaller(ctx, "user2", phone)
	if err != nil {
		t.Error("RateLimiter blocked the wrong user")
	}
}

func TestConcurrentAccess(t *testing.T) {
	service := NewTrueCallerServiceWithInMemoryDB()
	ctx := context.Background()
	phone := &PhoneNumber{countryCode: "+1", number: "1231231234"}
	userID := "user1"

	const numWorkers = 50

	errChan := make(chan error, numWorkers)

	for i := 0; i < numWorkers; i++ {
		go func(id int) {
			var err error

			switch id % 4 {
			case 0:
				_, err = service.IdentifyCaller(ctx, userID, phone)
			case 1:
				err = service.ReportSpam(ctx, userID, phone)
			case 2:
				err = service.BlockNumber(ctx, userID, phone)
			case 3:
				_, err = service.SearchNumber(ctx, userID, phone)
			}

			errChan <- err
		}(i)
	}

	for i := 0; i < numWorkers; i++ {
		err := <-errChan
		if err != nil && err != ErrSpamReportExists {
			t.Errorf("Unexpected error in concurrent operation: %v", err)
		}
	}

	// Verify final state
	user, _ := service.db.GetUserByID(ctx, userID)
	if _, blocked := user.Blocked[phone.String()]; !blocked {
		t.Error("Expected number to be blocked after concurrent access")
	}

	count, _ := service.db.GetSpamCount(ctx, phone)
	if count < 1 {
		t.Error("Expected spam count to be incremented after concurrent access")
	}
}
