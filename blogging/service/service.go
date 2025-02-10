package service

import (
	"lld/blogging/model"
	"lld/blogging/util"

	"gorm.io/gorm"
)

type BlogService struct {
	db *gorm.DB
}

func (service *BlogService) RegisterUser(user model.User) error {
	err := service.db.Create(&user).Error
	return err
}

func (service *BlogService) LoginUser(input model.Input) (string, error) {
	util := util.Util{}
	var user model.User
	err := service.db.Where("email= ?", input.Email).First(&user).Error
	if err != nil {
		return "", err
	}
	token, tokenError := util.GenerateToken(user)
	return token, tokenError
}
