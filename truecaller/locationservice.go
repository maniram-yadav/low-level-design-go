package truecaller

type LocationService interface {
	GetLocation(phone *PhoneNumber) (string, error)
}
