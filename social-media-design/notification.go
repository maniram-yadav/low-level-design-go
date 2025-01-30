package main

type NotificationType string

const (
	MessageNotificationType               NotificationType = "Message"
	CommentNotificationType               NotificationType = "Comment"
	LikeNotificationType                  NotificationType = "Like"
	FriendRequestNotificationType         NotificationType = "FriendRequest"
	FriendRequestAcceptedNotificationType NotificationType = "FriendRequeestAccepted"
	MentionNotificationType               NotificationType = "Mention"
)

type Notification struct {
	Id      string
	Content string
	Type    NotificationType
	UserId  int
}

func NewNotification(userid int, content string, notType NotificationType, id string) *Notification {
	notification := &Notification{Id: id, Content: content, Type: notType, UserId: userid}
	return notification
}
