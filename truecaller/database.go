package truecaller

import "context"

type Database interface {
	GetUserByID(ctx context.Context, id string) (*User, error)
	GetUserByPhone(ctx context.Context, phone *PhoneNumber) (*User, error)
	GetContact(ctx context.Context, phone *PhoneNumber) (*Contact, error)
	SaveUser(ctx context.Context, user *User) error
	SaveContact(ctx context.Context, contact *Contact) error
	IncrementSpamCount(ctx context.Context, phone *PhoneNumber) error
	GetSpamCount(ctx context.Context, phone *PhoneNumber) (int, error)
}
