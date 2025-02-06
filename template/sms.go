package template

import (
	"fmt"
	"strconv"
)

type Sms struct {
	Otp
}

func (s *Sms) GenRandomOTP(len int) string {
	randomOtp := "6767" + strconv.Itoa(len)
	fmt.Printf("\nSMS : generating rando OTP %s \n ", randomOtp)
	return randomOtp
}

func (s *Sms) SaveOTPCache(otp string) {
	fmt.Printf("\nSms otp %s saved in cache\n ", otp)
}

func (s *Sms) GetMessage(otp string) string {
	return "SMS OTP for login is " + otp
}

func (s *Sms) SendNotification(message string) error {
	fmt.Printf("SMS sending sms:%s \n", message)
	return nil
}
