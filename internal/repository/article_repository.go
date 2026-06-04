package repository

import (
	"MyBlogs/internal/model/entity"
	"fmt"
	"gorm.io/gorm"
)

type ArticleRepoImpl struct {
	db *gorm.DB
}

func NewArticleRepo(db *gorm.DB) ArticleRepo {
	return &ArticleRepoImpl{db: db}
}

// CreateArticle 新建文章
func (art *ArticleRepoImpl) CreateArticle(article *entity.Article) error {
	result := art.db.Create(article)
	if result.Error != nil {
		return fmt.Errorf("Create() failed,err%w", result.Error)
	}
	return nil
}

// DeleteArticle 删除文章
func (art *ArticleRepoImpl) DeleteArticle(id int) error {
	result := art.db.Delete(&entity.Article{}, id)
	if result.Error != nil {
		return fmt.Errorf("Delete() failed,err%w", result.Error)
	}
	return nil
}

// FindArticleLimit 分页查找文章列表
func (art *ArticleRepoImpl) FindArticleLimit(aq ArticleQuery) ([]entity.Article, int, error) {
	var articles []entity.Article
	var total int64
	query := art.db.Model(&entity.Article{})
	baseQuery := query
	if aq.Keyword != "" {
		query = query.
			Select("articles.*", "match(title,content) against (? in natural language mode) as relevance", aq.Keyword).
			Omit("content", "user_id").
			Where("match(title,content) against (? in natural language mode)", aq.Keyword)
		baseQuery = baseQuery.Where("match(title,content) against (? in natural language mode)", aq.Keyword)
	}
	if aq.Status != 0 {
		query = query.Where("status = ?", aq.Status)
		baseQuery = baseQuery.Where("status = ?", aq.Status)
	}
	if aq.CategoryID != 0 {
		query = query.Where("category_id = ?", aq.CategoryID)
		baseQuery = baseQuery.Where("category_id = ?", aq.CategoryID)
	}
	if len(aq.TagsID) != 0 {
		query = query.Joins("join article_tags on article_tags.article_id = articles.id").
			Where("article_tags.tag_id in (?)", aq.TagsID).Group("articles.id").
			Having("count(distinct tag_id) = ?", len(aq.TagsID))
		baseQuery = baseQuery.Joins("join article_tags on article_tags.article_id = articles.id").
			Where("article_tags.tag_id in (?)", aq.TagsID).Group("articles.id").
			Having("count(distinct tag_id) = ?", len(aq.TagsID))
	}
	//查总数
	if err := baseQuery.Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("baseQuery.Count() failed,err:%w", err)
	}
	query = query.Offset((aq.Page - 1) * aq.Size).Limit(aq.Size)

	if aq.SortWay == 0 && aq.Keyword != "" {
		query = query.Order("relevance desc")
	} else if aq.SortWay == 1 {
		query = query.Order("created_at desc")
	} else if aq.SortWay == 2 {
		query = query.Order("view_count desc")
	}

	query = query.Preload("Tags").Preload("Category")

	result := query.Debug().Find(&articles)
	if result.Error != nil {
		return nil, 0, fmt.Errorf("query.Find() failed,err:%w", result.Error)
	}

	return articles, int(total), nil
}

//精准查找某个文章正文

// EditArticle 编辑文章
func (art *ArticleRepoImpl) EditArticle(article *entity.Article) error {
	result := art.db.Model(article).Updates(article)
	if result.Error != nil {
		return fmt.Errorf("Updates() failed,err%w", result.Error)
	}
	return nil
}
