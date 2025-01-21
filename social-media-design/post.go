package main

import (
	"sync"
	"time"
)

type Post struct {
	Id              int
	UserId          int
	Likes           int
	Comments        []*Comment
	Content         string
	IsPublished     bool
	URLs            []*string
	PublishedAt     time.Time
	HiddenFromusers map[int]bool
	CommentEnabled  bool
	mu              sync.RWMutex
}

func NewPost(data string, postId int, userid int, urls []*string) *Post {

	post := &Post{
		Id:              postId,
		UserId:          userid,
		Likes:           0,
		Comments:        make([]*Comment, 0),
		Content:         data,
		IsPublished:     true,
		URLs:            urls,
		PublishedAt:     time.Now(),
		HiddenFromusers: make(map[int]bool),
		CommentEnabled:  true,
		mu:              sync.RWMutex{},
	}

	return post
}

func (post *Post) AddComment(comment *Comment) {
	post.mu.Lock()
	defer post.mu.Unlock()
	post.Comments = append(post.Comments, comment)

}

func (post *Post) GetComments() []*Comment {
	post.mu.RLock()
	defer post.mu.RUnlock()
	return post.Comments

}

func (post *Post) DeleteComment(comment *Comment) {
	post.mu.Lock()
	defer post.mu.Unlock()
	for i, com := range post.Comments {
		if com.Id == comment.Id {
			post.Comments = append(post.Comments[:i], post.Comments[i+1:]...)
			break
		}
	}

}

func (post *Post) AddLike() {
	post.mu.Lock()
	defer post.mu.Unlock()
	post.Likes += 1
}

func (post *Post) GetLikes() int {
	post.mu.RLock()
	defer post.mu.RUnlock()
	return post.Likes
}

func (post *Post) IsHiddenFromUsers(userId int) bool {
	post.mu.RLock()
	defer post.mu.RUnlock()
	return post.HiddenFromusers[userId]
}

func (post *Post) HideFromUser(userId int) {
	post.mu.Lock()
	defer post.mu.Unlock()
	post.HiddenFromusers[userId] = true
}

func (post *Post) UnHiddenFromUser(userId int) {
	delete(post.HiddenFromusers, userId)
}

func (post *Post) EnableComments() {
	post.CommentEnabled = true
}

func (post *Post) DisableComments() {
	post.CommentEnabled = false
}

func (post *Post) isCommentEnabled() bool {
	return post.CommentEnabled
}

func (post *Post) PublishPost() {
	post.IsPublished = true
	post.PublishedAt = time.Now()
}

func (post *Post) UnPublishPost() {
	post.IsPublished = false
}

func (post *Post) UpdateContent(content string) {
	post.Content = content
}

func (post *Post) IsPostPublished() bool {
	return post.IsPublished
}
