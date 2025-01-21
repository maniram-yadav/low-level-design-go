package main

type User struct {
	Id       int
	Name     string
	Password string
	Email    string
	Friends  map[int]*User
	Posts    []*Post
}

func NewUser(id int, name string, pass string, email string) *User {
	user := &User{Id: id, Name: name, Password: pass, Email: email, Friends: make(map[int]*User), Posts: make([]*Post, 0)}
	return user
}

func (user *User) AddFriends(friend *User) {
	user.Friends[friend.Id] = friend
}

func (user *User) GetFriends() []*User {
	friends := make([]*User, 0, len(user.Friends))
	for _, friend := range user.Friends {
		friends = append(friends, friend)
	}
	return friends
}

func (user *User) RemoveFriend(friend *User) {
	delete(user.Friends, friend.Id)

}

func (user *User) AddPost(post *Post) {
	user.Posts = append(user.Posts, post)

}

func (user *User) GetPosts() []*Post {
	return user.Posts
}
