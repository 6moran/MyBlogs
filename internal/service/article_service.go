package service

import (
	"MyBlogs/internal/model/dto/request"
	"MyBlogs/internal/model/dto/response"
	"MyBlogs/internal/model/entity"
	"MyBlogs/internal/repository"
	"MyBlogs/pkg/minio"
	"fmt"
	"mime/multipart"
)

type ArticleServiceImpl struct {
	ArtRepo repository.ArticleRepo
}

func NewArticleService(ARepo repository.ArticleRepo) ArticleService {
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

// GetArticleLimit 后台得到全部文章列表
func (a *ArticleServiceImpl) GetArticleLimit(articleRequest request.QueryArticleRequest) ([]response.ArticleAdminResponse, int, error) {
	arts, total, err := a.ArtRepo.FindArticleLimit(repository.ArticleQuery{
		Status:     articleRequest.Status,
		CategoryID: articleRequest.CategoryID,
		TagsID:     articleRequest.TagsID,
		SortWay:    articleRequest.SortWay,
		QueryLimit: repository.QueryLimit{
			Page:    articleRequest.Page,
			Size:    articleRequest.Size,
			Keyword: articleRequest.KeyWord,
		},
	})
	if err != nil {
		return nil, 0, fmt.Errorf("FindArticleLimit() failed,err:%w", err)
	}

	res := make([]response.ArticleAdminResponse, len(arts))
	for k, v := range arts {
		tags := make([]response.TagResponse, len(v.Tags))
		for k2, v2 := range v.Tags {
			tags[k2] = response.TagResponse(v2)
		}
		res[k] = response.ArticleAdminResponse{
			ArticleResponse: response.ArticleResponse{
				ID:         int(v.ID),
				CreatedAt:  v.CreatedAt,
				UpdatedAt:  v.UpdatedAt,
				Title:      v.Title,
				Summary:    v.Summary,
				CoverImage: v.CoverImage,
				LikeCount:  v.LikeCount,
				ViewCount:  v.ViewCount,
			},
			Status:       v.Status,
			CommentCount: v.CommentCount,
			Category:     response.CategoryResponse(v.Category),
			Tags:         tags,
		}
	}

	return res, total, nil
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
	return minio.SaveImage(file, fileHeader)
}
