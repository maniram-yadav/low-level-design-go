package truecaller

type CarrierService interface {
	GetCarrier(phone *PhoneNumber) (string, error)
}
