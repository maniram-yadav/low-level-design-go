package main

type Comment struct {
	Id      string
	PostId  int
	Content string
	User    *User
}

func NewComment(id string, postId int, content string, user *User) *Comment {
	return &Comment{Id: id, PostId: postId, Content: content, User: user}
}
