package truecaller

import (
	"context"
	"sync"
)

type Database interface {
	GetUserByID(ctx context.Context, id string) (*User, error)
	GetUserByPhone(ctx context.Context, phone *PhoneNumber) (*User, error)
	GetContact(ctx context.Context, phone *PhoneNumber) (*Contact, error)
	SaveUser(ctx context.Context, user *User) error
	SaveContact(ctx context.Context, contact *Contact) error
	IncrementSpamCount(ctx context.Context, phone *PhoneNumber) error
	GetSpamCount(ctx context.Context, phone *PhoneNumber) (int, error)
}

type InMemoryDB struct {
	users        map[string]*User    // userID -> User
	usersByPhone map[string]string   // phone string -> userID
	contacts     map[string]*Contact // phone string -> Contact
	spamCounts   map[string]int      // phone string -> spam count
	mu           sync.RWMutex        // protects all maps
}

func NewInMemoryDB() *InMemoryDB {
	return &InMemoryDB{
		users:        make(map[string]*User),
		usersByPhone: make(map[string]string),
		contacts:     make(map[string]*Contact),
		spamCounts:   make(map[string]int),
	}
}

func (db *InMemoryDB) GetUserByID(ctx context.Context, id string) (*User, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	user, exists := db.users[id]
	if !exists {
		return nil, ErrUserNotFound
	}

	// Return a copy to prevent external modifications
	return copyUser(user), nil
}

func (db *InMemoryDB) GetUserByPhone(ctx context.Context, phone *PhoneNumber) (*User, error) {
	return nil, nil
}
func (db *InMemoryDB) GetContact(ctx context.Context, phone *PhoneNumber) (*Contact, error) {
	return nil, nil
}
func (db *InMemoryDB) SaveUser(ctx context.Context, user *User) error {
	return nil
}
func (db *InMemoryDB) SaveContact(ctx context.Context, contact *Contact) error {
	return nil
}
func (db *InMemoryDB) IncrementSpamCount(ctx context.Context, phone *PhoneNumber) error {
	return nil
}
func (db *InMemoryDB) GetSpamCount(ctx context.Context, phone *PhoneNumber) (int, error) {
	return 0, nil
}

func copyUser(user *User) *User {

	contactCopy := make(map[string]*Contact)
	for k, v := range user.Contacts {
		contactCopy[k] = copyContact(v)
	}
	blockedCopy := make(map[string]bool)
	for k, v := range user.Blocked {
		blockedCopy[k] = v
	}

	spamReportsCopy := make(map[string]bool)
	for k, v := range user.SpamReports {
		spamReportsCopy[k] = v
	}

	return &User{
		ID:          user.ID,
		Name:        user.Name,
		PhoneNumber: copyPhoneNumber(user.PhoneNumber),
		Contacts:    contactCopy,
		Blocked:     blockedCopy,
		SpamReports: spamReportsCopy,
		AccessToken: user.AccessToken,
		TokenExpiry: user.TokenExpiry,
	}

}

func copyContact(contact *Contact) *Contact {
	if contact == nil {
		return nil
	}
	phonesCopy := make([]*PhoneNumber, len(contact.PhoneNumbers))

	for i, p := range contact.PhoneNumbers {
		phonesCopy[i] = copyPhoneNumber(p)
	}
	contactCopy := &Contact{
		Name:         contact.Name,
		PhoneNumbers: phonesCopy,
		Email:        contact.Email,
		IsSpam:       contact.IsSpam,
		SpamCount:    contact.SpamCount,
		LastUpdated:  contact.LastUpdated,
	}

	return contactCopy
}

func copyPhoneNumber(phone *PhoneNumber) *PhoneNumber {
	if phone == nil {
		return nil
	}
	phoneNumber := &PhoneNumber{
		countryCode: phone.countryCode,
		number:      phone.number,
	}
	return phoneNumber
}
