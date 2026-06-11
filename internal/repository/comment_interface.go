package repository

import "MyBlogs/internal/model/entity"

type CommentRepo interface {
	FindByArticleID(articleID int, page, size int) ([]entity.Comment, int64, error)
	FindByID(id int) (*entity.Comment, error)
	Create(comment *entity.Comment) error
	Delete(id int) error
	DeleteByParentID(parentID int) error
	DeleteByRootID(rootID int) error
	CountByArticleID(articleID int) (int64, error)
}
