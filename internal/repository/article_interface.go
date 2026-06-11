package repository

import "MyBlogs/internal/model/entity"

type ArticleRepo interface {
	CreateArticle(article *entity.Article) error
	DeleteArticle(id int) error
	EditArticle(article *entity.Article) error
	FindArticleLimit(aq ArticleQuery) ([]entity.Article, int, error)
	FindByID(id int) (*entity.Article, error)
	FindPublished(page, size int, tagID int, keyword string) ([]entity.Article, int64, error)
	IncrementViewCount(id int) (int, error)
	GetLikeCount(id int) (int, error)
	SetLikeCount(id int, count int) error

	// 点赞相关（Redis）
	IncrLike(articleID int) (int, error)
	GetRedisLikeCount(articleID int) (int, error)
	ExistsLikeKey(articleID int) (bool, error)
	SetRedisLikeCount(articleID int, count int) error
	GetAllLikeCounts() (map[int]int, error)

	// 文章详情缓存（Redis）
	GetArticleDetailCache(articleID int) (*entity.Article, error)
	SetArticleDetailCache(articleID int, article *entity.Article) error
}
