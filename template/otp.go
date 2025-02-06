package template

type Otp struct {
	Iotp IOtp
}

func (o *Otp) GenAndSendOTP(otpLength int) error {
	otp := o.Iotp.GenRandomOTP(otpLength)
	o.Iotp.SaveOTPCache(otp)
	message := o.Iotp.GetMessage(otp)
	err := o.Iotp.SendNotification(message)

	if err != nil {
		return err
	}
	return nil
}
