package truecaller

type CarrierService interface {
	GetCarrier(phone *PhoneNumber) (string, error)
}
type RandomCarrier struct{}

func (c *RandomCarrier) GetCarrier(phone *PhoneNumber) (string, error) {
	// Simple mock - in reality this would call an external API
	prefix := phone.number[:3]
	switch prefix {
	case "123":
		return "Carrier A", nil
	case "456":
		return "Carrier B", nil
	default:
		return "Default Carrier", nil
	}
}
