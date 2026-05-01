package services

import (
	"MyBlogs/internal/models/dto/request"
	"MyBlogs/internal/models/entity"
	"MyBlogs/internal/repositories"
	"MyBlogs/internal/utils"
	"fmt"
	"mime/multipart"
)

type ArticleServiceImpl struct {
	ArtRepo repositories.ArticleRepo
}

func NewArticleService(ARepo repositories.ArticleRepo) ArticleService {
	return &ArticleServiceImpl{
		ArtRepo: ARepo,
	}
}

// 创建新的文章
func (a *ArticleServiceImpl) CreateNewArticle(request request.CreateArticleRequest) error {
	err := a.ArtRepo.CreateArticle(&entity.Article{
		UserID:     request.UserID,
		CategoryID: request.CategoryID,
		Title:      request.Title,
		Summary:    request.Summary,
		CoverImage: request.CoverImage,
		Content:    request.Content,
		Status:     request.Status,
	})
	if err != nil {
		return fmt.Errorf("CreateArticle() failed,err:%w", err)
	}
	return nil
}

// 得到全部文章
func (a *ArticleServiceImpl) GetArticleLimit(articleRequest request.QueryArticleRequest) ([]entity.Article, int, error) {
	arts, total, err := a.ArtRepo.FindArticleLimit(repositories.ArticleQuery{
		Status:     articleRequest.Status,
		CategoryID: articleRequest.CategoryID,
		TagsID:     articleRequest.TagsID,
		SortWay:    articleRequest.SortWay,
		QueryLimit: repositories.QueryLimit{
			Page:    articleRequest.Page,
			Size:    articleRequest.Size,
			Keyword: articleRequest.KeyWord,
		},
	})
	if err != nil {
		return nil, 0, fmt.Errorf("FindArticleLimit() failed,err:%w", err)
	}
	return arts, total, nil
}

// 通过id删除文章
func (a *ArticleServiceImpl) DeleteArticle(id int) error {
	err := a.ArtRepo.DeleteArticle(id)
	if err != nil {
		return fmt.Errorf("DeleteArticle() failed,err:%w", err)
	}
	return nil
}

// 上传文章图片
func (a *ArticleServiceImpl) SaveImage(file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	return utils.SaveImage(file, fileHeader)
}
