package service

import "MyBlogs/internal/model/dto/response"

type CommentService interface {
	GetCommentsByArticleID(articleID, page, size int) (*response.CommentListResponse, error)
	CreateComment(articleID, userID int, content string, parentID int) error
	DeleteComment(id int) error
}
