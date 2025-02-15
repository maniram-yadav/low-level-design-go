package test

import (
	"fmt"
	"lld/template"
	"testing"
)

func TestTemplate(t *testing.T) {
	sms := &template.Sms{}
	o := template.Otp{
		Iotp: sms,
	}
	o.GenAndSendOTP(4)
	fmt.Print("\n\n")
	emailOtp := &template.Email{}
	o = template.Otp{
		Iotp: emailOtp,
	}
	o.GenAndSendOTP(4)
	t.Log("Template method test completed")
}
