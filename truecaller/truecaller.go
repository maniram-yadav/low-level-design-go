package truecaller

import (
	"fmt"
	"regexp"
)

type PhoneNumber struct {
	countryCode string
	number      string
}

func NewPhoneNumber(countryCode, number string) (*PhoneNumber, error) {

	valid := regexp.MustCompile(`^\+?[0-9]{1,3}$`).MatchString(countryCode)
	if !valid {
		return nil, ErrInvalidNumber
	}
	valid = regexp.MustCompile(`^[0-9]{6,15}$`).MatchString(number)
	if !valid {
		return nil, ErrInvalidNumber
	}
	return &PhoneNumber{
		countryCode: countryCode,
		number:      number,
	}, nil

}

func (p *PhoneNumber) String() string {
	return fmt.Sprintf("%s%s", p.countryCode, p.number)
}
