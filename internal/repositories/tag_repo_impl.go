package repositories

import (
	"MyBlogs/internal/models/entity"
	"fmt"
	"gorm.io/gorm"
)

type TagRepoImpl struct {
	db *gorm.DB
}

func NewTagRepo(db *gorm.DB) TagRepo {
	return &TagRepoImpl{db: db}
}

func (tr *TagRepoImpl) CreateTag(tag *entity.Tag) error {
	result := tr.db.Create(tag)
	if result.Error != nil {
		return fmt.Errorf("Create() failed,err%w", result.Error)
	}
	return nil
}

func (tr *TagRepoImpl) DeleteTag(id int) error {
	result := tr.db.Delete(&entity.Category{}, id)
	if result.Error != nil {
		return fmt.Errorf("Delete() failed,err:%w", result.Error)
	}
	return nil
}

func (tr *TagRepoImpl) QueryTagLimit(limit QueryLimit) ([]entity.Tag, error) {
	var tags []entity.Tag
	result := tr.db.Offset((limit.Page - 1) * limit.Size).Limit(limit.Size).Find(tags)
	if result.Error != nil {
		return nil, fmt.Errorf("Find() failed,err:%w", result.Error)
	}
	return tags, nil
}

func (tr *TagRepoImpl) EditTag(tag *entity.Tag) error {
	result := tr.db.Model(tag).Updates(tag)
	if result.Error != nil {
		return fmt.Errorf("Updates() failed,err%w", result.Error)
	}
	return nil
}
