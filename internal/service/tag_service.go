package service

import (
	"MyBlogs/internal/model/dto/request"
	"MyBlogs/internal/model/entity"
	"MyBlogs/internal/repository"
	"fmt"
)

type TagServiceImpl struct {
	TagRepo repository.TagRepo
}

func NewTagService(repo repository.TagRepo) TagService {
	return &TagServiceImpl{TagRepo: repo}
}

func (ts *TagServiceImpl) CreateNewTag(request request.TagRequest) error {
	err := ts.TagRepo.CreateTag(&entity.Tag{
		Name: request.Name,
	})
	if err != nil {
		return fmt.Errorf("CreateTag() failed,err:%w", err)
	}
	return nil
}

func (ts *TagServiceImpl) DeleteTag() error {
	return nil
}

func (ts *TagServiceImpl) EditTag() error {
	return nil
}

func (ts *TagServiceImpl) QueryTags() error {
	return nil
}
