package iterator

type UserList struct {
	Users []*User
}

func NewList() *UserList {
	return &UserList{Users: make([]*User, 0)}
}

func (ul *UserList) Add(user User) {
	ul.Users = append(ul.Users, &user)
}

func (ul *UserList) CreateIterator() *UserIterator {
	return newUserIterator(ul.Users)
}
