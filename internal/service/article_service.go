package service

import (
	"MyBlogs/internal/model/dto/request"
	"MyBlogs/internal/model/dto/response"
	"MyBlogs/internal/model/entity"
	"MyBlogs/internal/repository"
	"MyBlogs/pkg/logger"
	"MyBlogs/pkg/minio"
	"fmt"
	"mime/multipart"
	"time"

	"go.uber.org/zap"
)

type articleService struct {
	ArtRepo repository.ArticleRepo
}

func NewArticleService(ARepo repository.ArticleRepo) ArticleService {
	return &articleService{
		ArtRepo: ARepo,
	}
}

// 创建新的文章
func (a *articleService) CreateNewArticle(request request.CreateArticleRequest) error {
	err := a.ArtRepo.CreateArticle(&entity.Article{
		UserID:     request.UserID,
		Title:      request.Title,
		Summary:    request.Summary,
		CoverImage: request.CoverImage,
		Content:    request.Content,
		Status:     request.Status,
	})
	if err != nil {
		return fmt.Errorf("创建文章失败: %w", err)
	}
	return nil
}

// GetArticleLimit 后台得到全部文章列表
func (a *articleService) GetArticleLimit(articleRequest request.QueryArticleRequest) ([]response.ArticleAdminResponse, int, error) {
	arts, total, err := a.ArtRepo.FindArticleLimit(repository.ArticleQuery{
		Status:  articleRequest.Status,
		TagsID:  articleRequest.TagsID,
		SortWay: articleRequest.SortWay,
		QueryLimit: repository.QueryLimit{
			Page:    articleRequest.Page,
			Size:    articleRequest.Size,
			Keyword: articleRequest.KeyWord,
		},
	})
	if err != nil {
		return nil, 0, fmt.Errorf("查询文章列表失败: %w", err)
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
			Tags:         tags,
		}
	}

	return res, total, nil
}

// 通过id删除文章
func (a *articleService) DeleteArticle(id int) error {
	err := a.ArtRepo.DeleteArticle(id)
	if err != nil {
		return fmt.Errorf("删除文章失败: %w", err)
	}
	return nil
}

// 上传文章图片
func (a *articleService) SaveImage(file multipart.File, fileHeader *multipart.FileHeader, articleID int) (string, error) {
	return minio.SaveImage(file, fileHeader, articleID)
}

// GetPublishedArticles 获取已发布的文章列表
func (a *articleService) GetPublishedArticles(page, size int, tagID int, keyword string) (*response.ArticleListResponse, error) {
	articles, total, err := a.ArtRepo.FindPublished(page, size, tagID, keyword)
	if err != nil {
		return nil, fmt.Errorf("查询已发布文章失败: %w", err)
	}

	items := make([]response.ArticleListItemResponse, 0, len(articles))
	for _, v := range articles {
		tags := make([]response.TagResponse, len(v.Tags))
		for i, t := range v.Tags {
			tags[i] = response.TagResponse(t)
		}

		likeCount := a.getLikeCount(int(v.ID), v.LikeCount)

		items = append(items, response.ArticleListItemResponse{
			ID:           int(v.ID),
			Title:        v.Title,
			Summary:      v.Summary,
			CoverImage:   v.CoverImage,
			ViewCount:    v.ViewCount,
			LikeCount:    likeCount,
			CommentCount: v.CommentCount,
			CreatedAt:    v.CreatedAt,
			Tags:         tags,
		})
	}

	return &response.ArticleListResponse{
		Articles: items,
		Total:    total,
		Page:     page,
		Size:     size,
	}, nil
}

// getLikeCount 获取文章点赞数（优先Redis，否则重建缓存）
func (a *articleService) getLikeCount(articleID int, defaultCount int) int {
	count, err := a.ArtRepo.GetRedisLikeCount(articleID)
	if err != nil {
		logger.Warn("Redis获取点赞数失败，重建缓存", zap.Int("articleID", articleID), zap.Error(err))
		if setErr := a.ArtRepo.SetRedisLikeCount(articleID, defaultCount); setErr != nil {
			logger.Error("重建点赞数缓存失败", zap.Int("articleID", articleID), zap.Error(setErr))
		}
		return defaultCount
	}
	return count
}

// GetArticleDetail 获取文章详情
func (a *articleService) GetArticleDetail(id int) (*response.ArticleDetailResponse, error) {
	// 优先从缓存获取
	article, err := a.ArtRepo.GetArticleDetailCache(id)
	if err != nil {
		logger.Warn("获取文章详情缓存失败，从数据库查询", zap.Int("articleID", id), zap.Error(err))
		article, err = a.ArtRepo.FindByID(id)
		if err != nil {
			return nil, fmt.Errorf("查询文章详情失败: %w", err)
		}
		// 缓存文章详情
		if cacheErr := a.ArtRepo.SetArticleDetailCache(id, article); cacheErr != nil {
			logger.Error("缓存文章详情失败", zap.Int("articleID", id), zap.Error(cacheErr))
		}
	}

	// 增加浏览量
	viewCount, err := a.ArtRepo.IncrementViewCount(id)
	if err != nil {
		logger.Error("增加浏览量失败", zap.Int("articleID", id), zap.Error(err))
		viewCount = article.ViewCount
	}

	likeCount := a.getLikeCount(id, article.LikeCount)

	tags := make([]response.TagResponse, len(article.Tags))
	for i, t := range article.Tags {
		tags[i] = response.TagResponse(t)
	}

	return &response.ArticleDetailResponse{
		ArticleResponse: response.ArticleResponse{
			ID:         int(article.ID),
			CreatedAt:  article.CreatedAt,
			UpdatedAt:  article.UpdatedAt,
			Title:      article.Title,
			Summary:    article.Summary,
			CoverImage: article.CoverImage,
			LikeCount:  likeCount,
			ViewCount:  viewCount,
		},
		Content:      article.Content,
		CommentCount: article.CommentCount,
		Tags:         tags,
	}, nil
}

// LikeArticle 点赞文章
func (a *articleService) LikeArticle(articleID int) (int, error) {
	exists, err := a.ArtRepo.ExistsLikeKey(articleID)
	if err != nil {
		return 0, fmt.Errorf("检查点赞记录失败: %w", err)
	}
	if !exists {
		likeCount, err := a.ArtRepo.GetLikeCount(articleID)
		if err != nil {
			return 0, fmt.Errorf("查询文章点赞数失败: %w", err)
		}
		if err := a.ArtRepo.SetRedisLikeCount(articleID, likeCount); err != nil {
			return 0, fmt.Errorf("初始化点赞缓存失败: %w", err)
		}
	}

	return a.ArtRepo.IncrLike(articleID)
}

// StartLikeSync 启动点赞数定时同步任务
func (a *articleService) StartLikeSync(interval time.Duration) {
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for range ticker.C {
			if err := a.syncLikeCounts(); err != nil {
				logger.Error("同步点赞数失败", zap.Error(err))
			}
		}
	}()
}

// syncLikeCounts 同步Redis中的点赞数到MySQL
func (a *articleService) syncLikeCounts() error {
	likeCounts, err := a.ArtRepo.GetAllLikeCounts()
	if err != nil {
		return err
	}

	for articleID, count := range likeCounts {
		if err := a.ArtRepo.SetLikeCount(articleID, count); err != nil {
			logger.Error("更新文章点赞数失败",
				zap.Int("articleID", articleID),
				zap.Int("count", count),
				zap.Error(err))
		}
	}

	return nil
}
