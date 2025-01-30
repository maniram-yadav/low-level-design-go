package main

import "fmt"

type ActivityManager struct {
	UserManager         *UserManager
	PostManager         *PostManager
	NotificationManager *Notificationmanager
}

func ActivityManagerFacade() *ActivityManager {
	return &ActivityManager{
		UserManager:         GetUserManagerInstance(),
		PostManager:         NewPostmanagerInstance(),
		NotificationManager: GetNotificationmanagerInstance(),
	}
}

func (am *ActivityManager) LoginUser(email, password string) (*User, error) {
	user, err := am.UserManager.LoginUser(email, password)
	return user, err
}

func (am *ActivityManager) AddUser(user *User) {
	am.UserManager.AddUser(user)
}

func (am *ActivityManager) RemoveUser(userId int) {
	am.UserManager.RemoveUser(userId)
}

func (am *ActivityManager) SendFriendRequest(requesterId, receiverId int) error {
	_, err := am.UserManager.GetUserById(receiverId)
	if err != nil {
		return err
	}
	am.NotificationManager.AddNotification(receiverId, FriendRequestNotificationType, fmt.Sprintf("%d have sent friend request to you", requesterId))
	fmt.Printf("Friend requeest sent to User id %d", requesterId)
	return nil

}

func (am *ActivityManager) AcceptFriendRequest(requesterId, receiverId int) error {
	err := am.UserManager.AddFriend(requesterId, receiverId)
	if err != nil {
		return err
	}
	am.NotificationManager.AddNotification(requesterId, FriendRequestAcceptedNotificationType, fmt.Sprintf("%d have accepted friend request for you", receiverId))
	fmt.Printf("Friend requeest accepted by %d", receiverId)
	return nil
}

//post related methods
func (am *ActivityManager) AddPost(post *Post) error {

	user, err := am.UserManager.GetUserById(post.UserId)
	if err != nil {
		return err
	}
	am.PostManager.AddPost(post, user)
	fmt.Printf("\nPost Added with Id %d\n", post.Id)
	return nil
}

func (am *ActivityManager) GetUserFeed(userid int) ([]*Post, error) {
	user, err := am.UserManager.GetUserById(userid)
	if err != nil {
		return nil, err
	}
	return am.PostManager.GetuserFeed(user), nil
}

func (am *ActivityManager) LikePost(userid, postid int) error {
	err := am.PostManager.LikePost(postid)
	if err != nil {
		return err
	}

	am.NotificationManager.AddNotification(userid, LikeNotificationType, fmt.Sprintf("User %d have liked Post %d", userid, postid))
	fmt.Printf("Post Liked %d \n", postid)
	return nil
}

func (am *ActivityManager) CommentPost(userid, postid int, commentText string) error {
	user, err := am.UserManager.GetUserById(userid)
	if err != nil {
		return err
	}
	post, err := am.PostManager.CommentPost(postid, commentText, user)
	am.NotificationManager.AddNotification(post.UserId, CommentNotificationType, fmt.Sprintf("User %d have commented at Post %d", userid, postid))
	fmt.Printf("Comment added on post %d \n", postid)
	return nil
}

func (am *ActivityManager) MentionUserInPost(postid, mentionUserId int) error {
	post, err := am.PostManager.GetPost(postid)
	if err != nil {
		return err
	}
	am.NotificationManager.AddNotification(post.UserId, MentionNotificationType, fmt.Sprintf("User %d have been mentioned at Post %d", mentionUserId, postid))
	fmt.Printf("User %d Mentioned in  post %d \n", mentionUserId, postid)
	return nil
}

func (am *ActivityManager) PublishPost(postId int) error {
	return am.PostManager.PublishPost(postId)
}

func (am *ActivityManager) UnPublishPost(postId int) error {
	return am.PostManager.UnPublishPost(postId)
}

func (am *ActivityManager) UpdatePost(postId int, content string) error {
	return am.PostManager.UpdatePost(postId, content)
}

func (am *ActivityManager) HidePostFromUser(postId int, userid int) error {
	_, err := am.UserManager.GetUserById(userid)
	if err != nil {
		return err
	}
	return am.PostManager.HidePostFromUser(postId, userid)
}

func (am *ActivityManager) EnableComment(postId int) error {
	return am.PostManager.EnableComments(postId)
}

func (am *ActivityManager) DisableComment(postId int) error {
	return am.PostManager.DisableComments(postId)
}

func (am *ActivityManager) GetNotification(userId int) ([]*Notification, error) {
	return am.NotificationManager.GetNotifications(userId)
}
