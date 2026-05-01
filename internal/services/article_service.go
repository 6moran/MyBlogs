package services

import (
	"MyBlogs/internal/models/dto/request"
	"MyBlogs/internal/models/entity"
	"mime/multipart"
)

type ArticleService interface {
	CreateNewArticle(request request.CreateArticleRequest) error
	GetArticleLimit(articleRequest request.QueryArticleRequest) ([]entity.Article, int, error)
	DeleteArticle(id int) error
	SaveImage(file multipart.File, fileHeader *multipart.FileHeader) (string, error)
}
