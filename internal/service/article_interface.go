package service

import (
	"MyBlogs/internal/model/dto/request"
	"MyBlogs/internal/model/dto/response"
	"mime/multipart"
)

type ArticleService interface {
	CreateNewArticle(request request.CreateArticleRequest) error
	GetArticleLimit(articleRequest request.QueryArticleRequest) ([]response.ArticleAdminResponse, int, error)
	DeleteArticle(id int) error
	SaveImage(file multipart.File, fileHeader *multipart.FileHeader) (string, error)
}
