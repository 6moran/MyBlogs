package repository

import (
	"MyBlogs/internal/model/entity"
	"fmt"
	"gorm.io/gorm"
)

type tagRepo struct {
	db *gorm.DB
}

func NewTagRepo(db *gorm.DB) TagRepo {
	return &tagRepo{db: db}
}

func (tr *tagRepo) CreateTag(tag *entity.Tag) error {
	result := tr.db.Create(tag)
	if result.Error != nil {
		return fmt.Errorf("创建标签失败: %w", result.Error)
	}
	return nil
}

func (tr *tagRepo) DeleteTag(id int) error {
	result := tr.db.Delete(&entity.Tag{}, id)
	if result.Error != nil {
		return fmt.Errorf("删除标签失败: %w", result.Error)
	}
	return nil
}

func (tr *tagRepo) QueryTagLimit(limit QueryLimit) ([]entity.Tag, error) {
	var tags []entity.Tag
	result := tr.db.Offset((limit.Page - 1) * limit.Size).Limit(limit.Size).Find(&tags)
	if result.Error != nil {
		return nil, fmt.Errorf("查询标签列表失败: %w", result.Error)
	}
	return tags, nil
}

func (tr *tagRepo) EditTag(tag *entity.Tag) error {
	result := tr.db.Model(tag).Updates(tag)
	if result.Error != nil {
		return fmt.Errorf("更新标签失败: %w", result.Error)
	}
	return nil
}

func (tr *tagRepo) FindAll() ([]entity.Tag, error) {
	var tags []entity.Tag
	if err := tr.db.Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}

func (tr *tagRepo) FindByID(id int) (*entity.Tag, error) {
	var tag entity.Tag
	if err := tr.db.First(&tag, id).Error; err != nil {
		return nil, err
	}
	return &tag, nil
}
