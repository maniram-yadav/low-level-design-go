package truecaller

type CarrierService interface {
	GetCarrier(phone *PhoneNumber) (string, error)
}
type CustomCarrierService struct{}

func (c *CustomCarrierService) GetCarrier(phone *PhoneNumber) (string, error) {

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
