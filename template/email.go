package template

import (
	"fmt"
	"strconv"
)

type Email struct {
	Otp
}

func (s *Email) GenRandomOTP(len int) string {
	randomOtp := "3447" + strconv.Itoa(len)
	fmt.Printf("\nEmail : generating rando OTP %s \n ", randomOtp)
	return randomOtp
}

func (s *Email) SaveOTPCache(otp string) {
	fmt.Printf("\nEmail otp %s saved in cache\n ", otp)
}

func (s *Email) GetMessage(otp string) string {
	return "\nEmail OTP for login is " + otp
}

func (s *Email) SendNotification(message string) error {
	fmt.Printf("\nEmail sending sms:%s \n", message)
	return nil
}
