package main

import (
	"fmt"
	"sync"
	"time"
)

type PostManager struct {
	Posts map[int]*Post
	mu    sync.RWMutex
}

var (
	PostManagerInstance *PostManager
	postOnce            sync.Once
)

func NewPostmanagerInstance() *PostManager {
	postOnce.Do(func() {
		PostManagerInstance = &PostManager{Posts: make(map[int](*Post))}
	})
	return PostManagerInstance
}

func (pm *PostManager) DisableComments(postid int) error {
	pm.mu.Lock()
	defer pm.mu.Unlock()
	post, ok := pm.Posts[postid]
	if !ok {
		return fmt.Errorf("Post not found with id %d", postid)
	}
	post.DisableComments()
	fmt.Printf("Post Comments Disabled for id %d", postid)
	return nil
}

func (pm *PostManager) EnableComments(postid int) error {
	pm.mu.Lock()
	defer pm.mu.Unlock()
	post, ok := pm.Posts[postid]

	if !ok {
		return fmt.Errorf("Post not found with id %d", postid)
	}

	post.EnableComments()
	fmt.Printf("Post Comments Enabled for post id %d", postid)
	return nil
}

func (pm *PostManager) LikePost(postId int) error {
	pm.mu.Lock()
	defer pm.mu.Unlock()
	post, ok := pm.Posts[postId]

	if !ok {
		return fmt.Errorf("Post not found with id %d", postId)
	}

	post.AddLike()
	fmt.Printf("Post liked with postid %d", postId)
	return nil
}

func (pm *PostManager) AddComment(postId int, comment *Comment) error {
	pm.mu.Lock()
	defer pm.mu.Unlock()
	post, ok := pm.Posts[postId]

	if !ok {
		return fmt.Errorf("Post not found with id %d", postId)
	}

	post.AddComment(comment)
	fmt.Printf("Post Comment added for post id %d", postId)
	return nil
}

func (pm *PostManager) PublishPost(postId int) error {
	pm.mu.Lock()
	defer pm.mu.Unlock()
	post, ok := pm.Posts[postId]

	if !ok {
		return fmt.Errorf("Post not found with id %d", postId)
	}

	post.PublishPost()
	fmt.Printf("Post Published with id %d", postId)
	return nil
}

func (pm *PostManager) UnPublishPost(postId int) error {
	pm.mu.Lock()
	defer pm.mu.Unlock()
	post, ok := pm.Posts[postId]

	if !ok {
		return fmt.Errorf("Post not found with id %d", postId)
	}

	post.UnPublishPost()
	fmt.Printf("Post UnPublished with id %d", postId)
	return nil
}

func (pm *PostManager) HidePostFromUser(postId int, userId int) error {
	pm.mu.Lock()
	defer pm.mu.Unlock()
	post, ok := pm.Posts[postId]

	if !ok {
		return fmt.Errorf("Post not found with id %d", postId)
	}
	post.HideFromUser(userId)

	fmt.Printf("Post Hidden for user %d with post id %d", userId, postId)
	return nil
}

func (pm *PostManager) UnHidePostFromUser(postId int, userId int) error {
	pm.mu.Lock()
	defer pm.mu.Unlock()
	post, ok := pm.Posts[postId]

	if !ok {
		return fmt.Errorf("Post not found with id %d", postId)
	}
	post.UnHiddenFromUser(userId)

	fmt.Printf("Post UnHidden for user %d with post id %d", userId, postId)
	return nil
}

func (pm *PostManager) GetuserFeed(user *User) []*Post {

	pm.mu.Lock()
	defer pm.mu.Unlock()

	var feed []*Post
	feed = append(feed, user.GetPosts()...)
	friends := user.GetFriends()

	for _, friend := range friends {
		userPosts := friend.GetPosts()
		for _, post := range userPosts {
			if !post.IsHiddenFromUsers(friend.Id) && post.IsPostPublished() {
				feed = append(feed, post)
			}
		}
	}
	return feed
}

func (pm *PostManager) UpdatePost(postId int, contnent string) error {
	pm.mu.Lock()
	defer pm.mu.Unlock()
	post, ok := pm.Posts[postId]

	if !ok {
		return fmt.Errorf("Post not found with id %d", postId)
	}
	post.UpdateContent(contnent)

	fmt.Printf("Post content update post id %d", postId)
	return nil
}

func (pm *PostManager) CommentPost(postId int, commenttext string, user *User) (*Post, error) {
	pm.mu.Lock()
	defer pm.mu.Unlock()
	post, ok := pm.Posts[postId]

	if !ok {
		return nil, fmt.Errorf("Post not found with id %d", postId)
	}

	comment := NewComment(fmt.Sprintf("Comment.%d", time.Now().UnixMicro()), postId, commenttext, user)
	post.AddComment(comment)
	fmt.Printf("Post content update post id %d", postId)
	return post, nil
}

func (pm *PostManager) AddPost(post *Post, user *User) (*Post, error) {
	pm.mu.Lock()
	defer pm.mu.Unlock()
	pm.Posts[post.Id] = post
	user.AddPost(post)
	return post, nil
}

func (pm *PostManager) GetPost(postId int) (*Post, error) {
	pm.mu.Lock()
	defer pm.mu.Unlock()
	post, ok := pm.Posts[postId]

	if !ok {
		return nil, fmt.Errorf("Post not found with id %d", postId)
	}

	fmt.Printf("Post content received for post id %d", postId)
	return post, nil
}
