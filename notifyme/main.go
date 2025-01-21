package main

import (
	"lld/notifyme/observable"
	"lld/notifyme/observer"
	"os"
)

func main() {
	user1 := &observer.EmailNotificationObserver{EmailId: "abc1@email.com", Id: 1}
	user2 := &observer.SMSNotificationObserver{Phone: "9867534789", Id: 2}
	user3 := &observer.EmailNotificationObserver{EmailId: "abc3@email.com", Id: 3}
	item := observable.NewObservable()
	item.Add(user1)
	item.Add(user2)
	item.Add(user3)

	item.SetQuantity(9)
	item.Remove(user1)
	item.SetQuantity(0)
	os.Exit(1)
}
