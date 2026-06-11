package repository

import (
	"MyBlogs/internal/model/entity"

	"gorm.io/gorm"
)

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepo {
	return &commentRepository{db: db}
}

func (r *commentRepository) FindByArticleID(articleID int, page, size int) ([]entity.Comment, int64, error) {
	var comments []entity.Comment
	var total int64

	query := r.db.Model(&entity.Comment{}).Where("article_id = ?", articleID)

	// 只查根评论
	query = query.Where("parent_id = 0")

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * size
	if err := query.Offset(offset).Limit(size).Order("created_at DESC").Find(&comments).Error; err != nil {
		return nil, 0, err
	}

	return comments, total, nil
}

func (r *commentRepository) FindByID(id int) (*entity.Comment, error) {
	var comment entity.Comment
	if err := r.db.First(&comment, id).Error; err != nil {
		return nil, err
	}
	return &comment, nil
}

func (r *commentRepository) Create(comment *entity.Comment) error {
	return r.db.Create(comment).Error
}

func (r *commentRepository) Delete(id int) error {
	return r.db.Delete(&entity.Comment{}, id).Error
}

func (r *commentRepository) DeleteByParentID(parentID int) error {
	return r.db.Where("parent_id = ?", parentID).Delete(&entity.Comment{}).Error
}

func (r *commentRepository) DeleteByRootID(rootID int) error {
	return r.db.Where("root_id = ?", rootID).Delete(&entity.Comment{}).Error
}

func (r *commentRepository) CountByArticleID(articleID int) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Comment{}).Where("article_id = ?", articleID).Count(&count).Error
	return count, err
}
