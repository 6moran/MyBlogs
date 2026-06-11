package service

import (
	"MyBlogs/internal/model/dto/request"
	"MyBlogs/internal/model/dto/response"
)

type TagService interface {
	CreateTag(req request.TagRequest) error
	UpdateTag(id int, req request.TagRequest) error
	DeleteTag(id int) error
	GetTagByID(id int) (*response.TagResponse, error)
	GetTagList(page, size int) ([]response.TagResponse, int64, error)
	GetAllTags() ([]response.TagResponse, error)
}
