package repository

import "MyBlogs/internal/model/entity"

type CategoryRepo interface {
	CreateCategory(category *entity.Category) error
	DeleteCategory(id int) error
	QueryCategoryLimit(limit QueryLimit) ([]entity.Category, error)
	EditCategory(category *entity.Category) error
}
