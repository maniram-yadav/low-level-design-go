package service

import (
	"fmt"
	"lld/blogging/db"
	"lld/blogging/model"
	"lld/blogging/util"
	"sync"

	"gorm.io/gorm"
)

var (
	serviceInstance *BlogService
	once            sync.Once
)

type BlogService struct {
	db *gorm.DB
}

func NewBlogService() *BlogService {
	once.Do(func() {
		serviceInstance = &BlogService{db: db.InitDB()}
	})
	return serviceInstance
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

	if !util.Checkpassword(user.Password, input.Password) {
		return "", fmt.Errorf("invalid credentials")
	}
	token, tokenError := util.GenerateToken(user)
	return token, tokenError
}

func (service *BlogService) CreateBlog(blog *model.Blog) (*model.Blog, error) {
	err := service.db.Create(blog).Error
	return blog, err
}

func (service *BlogService) GetBlog(id int) (model.Blog, error) {
	var blog model.Blog
	err := service.db.Where("id= ?", id).First(&blog).Error
	return blog, err
}

func (service *BlogService) GetBlogs() ([]model.Blog, error) {
	var blogs []model.Blog
	err := service.db.Find(&blogs).Error
	return blogs, err
}

func (service *BlogService) UpdateBlog(blog *model.Blog) (*model.Blog, error) {
	err := service.db.Model(blog).Updates(blog).Error
	return blog, err
}
