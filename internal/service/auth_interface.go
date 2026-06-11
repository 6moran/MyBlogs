package service

import "MyBlogs/internal/model/dto/response"

type AuthService interface {
	Login(username, password string) (*response.LoginResponse, error)
	GitHubLogin(code string) (*response.LoginResponse, error)
}
