package service

import "MyBlogs/internal/repository"

type UserServiceImpl struct {
	UserRepo repository.UserRepo
}

func NewUserService(repo repository.UserRepo) UserService {
	return &UserServiceImpl{UserRepo: repo}
}
