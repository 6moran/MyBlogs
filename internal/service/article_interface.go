package service

import (
	"MyBlogs/internal/model/dto/request"
	"MyBlogs/internal/model/dto/response"
	"mime/multipart"
	"time"
)

type ArticleService interface {
	CreateNewArticle(request request.CreateArticleRequest) error
	GetArticleLimit(articleRequest request.QueryArticleRequest) ([]response.ArticleAdminResponse, int, error)
	DeleteArticle(id int) error
	SaveImage(file multipart.File, fileHeader *multipart.FileHeader, articleID int) (string, error)
	GetPublishedArticles(page, size int, tagID int, keyword string) (*response.ArticleListResponse, error)
	GetArticleDetail(id int) (*response.ArticleDetailResponse, error)
	LikeArticle(articleID int) (int, error)
	StartLikeSync(interval time.Duration)
}
