package truecaller

type LocationService interface {
	GetLocation(phone *PhoneNumber) (string, error)
}

type CusomLocationService struct{}

func (m *CusomLocationService) GetLocation(phone *PhoneNumber) (string, error) {
	prefix := phone.countryCode
	switch prefix {
	case "+1":
		return "United States", nil
	case "+44":
		return "United Kingdom", nil
	case "+91":
		return "India", nil
	default:
		return "Unknown Location", nil
	}
}
