package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	notificationmanagerInstance *Notificationmanager
	syncOnce                    sync.Once
)

type Notificationmanager struct {
	notification map[int][]*Notification
	mu           sync.RWMutex
}

func GetNotificationmanagerInstance() *Notificationmanager {
	syncOnce.Do(func() {
		notificationmanagerInstance = &Notificationmanager{notification: make(map[int][]*Notification)}
	})
	return notificationmanagerInstance
}

func (nm *Notificationmanager) GetNotifications(userid int) ([]*Notification, error) {
	nm.mu.RLock()
	defer nm.mu.Unlock()
	userNotifications, ok := nm.notification[userid]
	if !ok {
		return nil, fmt.Errorf("no notification found for the given user")
	}
	return userNotifications, nil
}

func (nm *Notificationmanager) AddNotification(userid int, notType NotificationType, message string) {
	notification := &Notification{UserId: userid, Content: message, Type: notType, Id: fmt.Sprintf("Notification.%d", time.Now().UnixMicro())}
	nm.notification[userid] = append(nm.notification[userid], notification)
}
