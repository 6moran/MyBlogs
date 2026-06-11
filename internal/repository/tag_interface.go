package repository

import "MyBlogs/internal/model/entity"

type TagRepo interface {
	CreateTag(tag *entity.Tag) error
	DeleteTag(id int) error
	QueryTagLimit(limit QueryLimit) ([]entity.Tag, error)
	EditTag(tag *entity.Tag) error
	FindAll() ([]entity.Tag, error)
	FindByID(id int) (*entity.Tag, error)
}
