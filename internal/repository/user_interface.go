package repository

import (
	"MyBlogs/internal/model/dto/request"
	"MyBlogs/internal/model/entity"
)

type UserRepo interface {
	FindByUsername(username string) (*entity.User, error)
	FindByGitHubID(githubID int) (*entity.User, error)
	FindByID(id int) (*entity.User, error)
	Create(user *entity.User) error
	Update(user *entity.User) error
	Delete(id int) error
	FindWithCondition(req request.QueryUserRequest) ([]entity.User, int64, error)
}
