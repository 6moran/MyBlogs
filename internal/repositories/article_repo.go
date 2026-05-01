package repositories

import "MyBlogs/internal/models/entity"

type ArticleRepo interface {
	CreateArticle(article *entity.Article) error
	DeleteArticle(id int) error
	EditArticle(article *entity.Article) error
	FindArticleLimit(aq ArticleQuery) ([]entity.Article, int, error)
}
