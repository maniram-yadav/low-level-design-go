package observer

import "fmt"

type NotifyObserver interface {
	Update(qty int)
	GetId() int
}

type EmailNotificationObserver struct {
	Id      int
	EmailId string
}

type SMSNotificationObserver struct {
	Id    int
	Phone string
}

func (eno *EmailNotificationObserver) Update(qty int) {
	fmt.Printf("\n> Email Sending email to %v : stocks available %v\n", eno.EmailId, qty)
}

func (eno *EmailNotificationObserver) GetId() int {
	return eno.Id
}

func (sno *SMSNotificationObserver) Update(qty int) {
	fmt.Printf("\n> SMS Sending sms to %v : stocks available %v\n", sno.Phone, qty)
}

func (sno *SMSNotificationObserver) GetId() int {
	return sno.Id
}
