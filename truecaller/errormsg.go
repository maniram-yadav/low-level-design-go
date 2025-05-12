package truecaller

import "errors"

var (
	ErrNumberNotFound   = errors.New("phone number not found")
	ErrInvalidNumber    = errors.New("invalid phone number format")
	ErrContactExists    = errors.New("contact already exists")
	ErrSpamReportExists = errors.New("spam report already exists")
	ErrUserNotFound     = errors.New("user not found")
	ErrUnauthorized     = errors.New("unauthorized access")
)
