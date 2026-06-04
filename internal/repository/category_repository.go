package repository

import (
	"MyBlogs/internal/model/entity"
	"fmt"
	"gorm.io/gorm"
)

type CategoryRepoImpl struct {
	db *gorm.DB
}

func NewCategoryRepo(db *gorm.DB) CategoryRepo {
	return &CategoryRepoImpl{db: db}
}

func (cr *CategoryRepoImpl) CreateCategory(category *entity.Category) error {
	result := cr.db.Create(category)
	if result.Error != nil {
		return fmt.Errorf("Create() failed,err%w", result.Error)
	}
	return nil
}

func (cr *CategoryRepoImpl) DeleteCategory(id int) error {
	result := cr.db.Delete(&entity.Category{}, id)
	if result.Error != nil {
		return fmt.Errorf("Delete() failed,err:%w", result.Error)
	}
	return nil
}
func (cr *CategoryRepoImpl) QueryCategoryLimit(limit QueryLimit) ([]entity.Category, error) {
	var categories []entity.Category
	result := cr.db.Offset((limit.Page - 1) * limit.Size).Limit(limit.Size).Find(categories)
	if result.Error != nil {
		return nil, fmt.Errorf("Find() failed,err:%w", result.Error)
	}
	return categories, nil
}

func (cr *CategoryRepoImpl) EditCategory(category *entity.Category) error {
	result := cr.db.Model(category).Updates(category)
	if result.Error != nil {
		return fmt.Errorf("Updates() failed,err%w", result.Error)
	}
	return nil
}
