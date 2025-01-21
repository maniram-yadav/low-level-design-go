package main

import (
	"fmt"
	"sync"
)

type UserManager struct {
	mu    sync.RWMutex
	Users map[int]*User
}

var (
	UserManagerInstance *UserManager
	userOnce            sync.Once
)

func GetUserManagerInstance() *UserManager {
	userOnce.Do(func() {
		UserManagerInstance = &UserManager{Users: make(map[int]*User)}
	})
	return UserManagerInstance
}

func (um *UserManager) AddUser(user *User) {
	um.mu.Lock()
	defer um.mu.Unlock()
	um.Users[user.Id] = user
	fmt.Printf("User %d with name %s added", user.Id, user.Name)
}

func (um *UserManager) GetUserById(userId int) (*User, error) {
	um.mu.RLock()
	defer um.mu.RUnlock()
	user, err := um.Users[userId]
	if err {
		return nil, fmt.Errorf("user not found")
	}
	fmt.Printf("user with id %d received", userId)
	return user, nil
}

func (um *UserManager) RemoveUser(userId int) {
	um.mu.Lock()
	defer um.mu.Unlock()
	delete(um.Users, userId)
	fmt.Printf("user %d removed", userId)
}

func (um *UserManager) UpdateUser(user *User) (*User, error) {
	um.mu.Lock()
	defer um.mu.Unlock()
	user, err := um.Users[user.Id]
	if err {
		return nil, fmt.Errorf("user not found")
	}
	um.Users[user.Id] = user
	fmt.Printf("User %d update", user.Id)
	return user, nil
}

func (um *UserManager) AddFriend(requesterid int, receiverid int) error {
	um.mu.Lock()
	defer um.mu.Unlock()
	requester, err1 := um.GetUserById(requesterid)
	receiver, err2 := um.GetUserById(receiverid)

	if err1 != nil {
		fmt.Printf("Requeting User does not exists")
		return err1
	}

	if err2 != nil {
		fmt.Printf("Receiving User does not exists")
		return err2
	}
	receiver.AddFriends(requester)
	requester.AddFriends(receiver)
	fmt.Printf("Friendship added between %d an %d", receiverid, receiverid)
	return nil
}

func (um *UserManager) LoginUser(email string, passowrd string) (*User, error) {
	um.mu.RLock()
	defer um.mu.RUnlock()

	for _, u := range um.Users {
		if u.Email == email && u.Password == passowrd {
			fmt.Printf("Login successfull for user with email %s", email)
			return u, nil
		}
	}
	return nil, fmt.Errorf("user not found with given credentials")
}
