package service

import (
	"MyBlogs/internal/model/dto/request"
	"MyBlogs/internal/model/dto/response"
)

type UserService interface {
	GetUserInfo(userID int) (*response.UserResponse, error)
	CreateUser(req request.CreateUserRequest) error
	UpdateUser(id int, req request.UpdateUserRequest) error
	DeleteUser(id int) error
	GetUserByID(id int) (*response.UserResponse, error)
	GetUserList(req request.QueryUserRequest) ([]response.UserResponse, int64, error)
}
