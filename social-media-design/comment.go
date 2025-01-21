package main

type Comment struct {
	Id      int
	PostId  int
	Content string
	User    *User
}

func NewComment(id int, postId int, content string, user *User) *Comment {
	return &Comment{Id: id, PostId: postId, Content: content, User: user}
}
