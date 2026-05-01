package repositories

import "MyBlogs/internal/models/entity"

type TagRepo interface {
	CreateTag(tag *entity.Tag) error
	DeleteTag(id int) error
	QueryTagLimit(limit QueryLimit) ([]entity.Tag, error)
	EditTag(tag *entity.Tag) error
}
