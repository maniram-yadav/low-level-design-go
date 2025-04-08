package services

import (
	"errors"
	"lld/stackoverflow/models"
	"lld/stackoverflow/repositories"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo *repositories.UserRepository
}

func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) Register(username, email, passwrod string) (*models.User, error) {
	existing, _ := s.userRepo.FindByEmail(email)
	if existing != nil {
		return nil, errors.New("email already registered")
	}
	hashed, err := bcrypt.GenerateFromPassword([]byte(passwrod), bcrypt.DefaultCost)

	if err != nil {
		return nil,err
	
	}
	user := &models.User{
		Username: username,
		Email: email,
		PasswordHash: string(hashed),
		Reputation: 1
	}

	if err := s.userRepo.Create(user);err!=nil {
		return nil,err
	}
	return user,err
}
