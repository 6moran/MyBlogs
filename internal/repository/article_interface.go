package repository

import "MyBlogs/internal/model/entity"

type ArticleRepo interface {
	CreateArticle(article *entity.Article) error
	DeleteArticle(id int) error
	EditArticle(article *entity.Article) error
	FindArticleLimit(aq ArticleQuery) ([]entity.Article, int, error)
}
