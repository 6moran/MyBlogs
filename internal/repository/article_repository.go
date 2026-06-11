package repository

import (
	"MyBlogs/internal/model/entity"
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

const likeKeyPrefix = "article:like:"
const articleDetailKeyPrefix = "article:detail:"

type articleRepo struct {
	db  *gorm.DB
	rdb *redis.Client
}

func NewArticleRepo(db *gorm.DB, rdb *redis.Client) ArticleRepo {
	return &articleRepo{db: db, rdb: rdb}
}

// CreateArticle 新建文章
func (art *articleRepo) CreateArticle(article *entity.Article) error {
	result := art.db.Create(article)
	if result.Error != nil {
		return fmt.Errorf("创建文章失败: %w", result.Error)
	}
	return nil
}

// DeleteArticle 删除文章
func (art *articleRepo) DeleteArticle(id int) error {
	result := art.db.Delete(&entity.Article{}, id)
	if result.Error != nil {
		return fmt.Errorf("删除文章失败: %w", result.Error)
	}
	return nil
}

// FindArticleLimit 分页查找文章列表
func (art *articleRepo) FindArticleLimit(aq ArticleQuery) ([]entity.Article, int, error) {
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
		return nil, 0, fmt.Errorf("查询文章总数失败: %w", err)
	}
	query = query.Offset((aq.Page - 1) * aq.Size).Limit(aq.Size)

	if aq.SortWay == 0 && aq.Keyword != "" {
		query = query.Order("relevance desc")
	} else if aq.SortWay == 1 {
		query = query.Order("created_at desc")
	} else if aq.SortWay == 2 {
		query = query.Order("view_count desc")
	}

	query = query.Preload("Tags")

	result := query.Debug().Find(&articles)
	if result.Error != nil {
		return nil, 0, fmt.Errorf("查询文章列表失败: %w", result.Error)
	}

	return articles, int(total), nil
}

//精准查找某个文章正文

// EditArticle 编辑文章
func (art *articleRepo) EditArticle(article *entity.Article) error {
	result := art.db.Model(article).Updates(article)
	if result.Error != nil {
		return fmt.Errorf("更新文章失败: %w", result.Error)
	}
	return nil
}

// FindByID 根据ID查找文章
func (art *articleRepo) FindByID(id int) (*entity.Article, error) {
	var article entity.Article
	err := art.db.Preload("Tags").Preload("User").First(&article, id).Error
	if err != nil {
		return nil, err
	}
	return &article, nil
}

// FindPublished 查找已发布的文章列表
func (art *articleRepo) FindPublished(page, size int, tagID int, keyword string) ([]entity.Article, int64, error) {
	var articles []entity.Article
	var total int64

	query := art.db.Model(&entity.Article{}).Where("status = 1")

	if tagID > 0 {
		query = query.Joins("JOIN article_tags ON article_tags.article_id = articles.id").
			Where("article_tags.tag_id = ?", tagID)
	}

	if keyword != "" {
		query = query.Where("title LIKE ? OR content LIKE ?", "%"+keyword+"%", "%"+keyword+"%")
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * size
	err := query.Preload("Tags").
		Offset(offset).Limit(size).
		Order("created_at DESC").
		Find(&articles).Error
	if err != nil {
		return nil, 0, err
	}

	return articles, total, nil
}

// IncrementViewCount 增加文章浏览量
func (art *articleRepo) IncrementViewCount(id int) (int, error) {
	var viewCount int
	result := art.db.Model(&entity.Article{}).Where("id = ?", id).
		UpdateColumn("view_count", gorm.Expr("view_count + 1")).
		Select("view_count").Row()
	if result.Err() != nil {
		return 0, result.Err()
	}
	if err := result.Scan(&viewCount); err != nil {
		return 0, err
	}
	return viewCount, nil
}

// GetLikeCount 获取文章点赞数
func (art *articleRepo) GetLikeCount(id int) (int, error) {
	var likeCount int
	err := art.db.Model(&entity.Article{}).Where("id = ?", id).Select("like_count").Scan(&likeCount).Error
	if err != nil {
		return 0, err
	}
	return likeCount, nil
}

// SetLikeCount 设置文章点赞数
func (art *articleRepo) SetLikeCount(id int, count int) error {
	return art.db.Model(&entity.Article{}).Where("id = ?", id).
		UpdateColumn("like_count", count).Error
}

func (art *articleRepo) likeKey(articleID int) string {
	return fmt.Sprintf("%s%d", likeKeyPrefix, articleID)
}

func (art *articleRepo) IncrLike(articleID int) (int, error) {
	ctx := context.Background()
	count, err := art.rdb.Incr(ctx, art.likeKey(articleID)).Result()
	if err != nil {
		return 0, err
	}
	return int(count), nil
}

func (art *articleRepo) GetRedisLikeCount(articleID int) (int, error) {
	ctx := context.Background()
	count, err := art.rdb.Get(ctx, art.likeKey(articleID)).Int()
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (art *articleRepo) ExistsLikeKey(articleID int) (bool, error) {
	ctx := context.Background()
	exists, err := art.rdb.Exists(ctx, art.likeKey(articleID)).Result()
	if err != nil {
		return false, err
	}
	return exists > 0, nil
}

func (art *articleRepo) SetRedisLikeCount(articleID int, count int) error {
	ctx := context.Background()
	return art.rdb.Set(ctx, art.likeKey(articleID), count, 0).Err()
}

func (art *articleRepo) GetAllLikeCounts() (map[int]int, error) {
	ctx := context.Background()

	keys, err := art.rdb.Keys(ctx, likeKeyPrefix+"*").Result()
	if err != nil {
		return nil, fmt.Errorf("获取keys失败: %w", err)
	}

	result := make(map[int]int)
	for _, key := range keys {
		parts := strings.Split(key, ":")
		if len(parts) != 3 {
			continue
		}
		articleID, err := strconv.Atoi(parts[2])
		if err != nil {
			continue
		}

		count, err := art.rdb.Get(ctx, key).Int()
		if err != nil {
			if err == redis.Nil {
				continue
			}
			continue
		}

		result[articleID] = count
	}

	return result, nil
}

// articleDetailKey 生成文章详情缓存 key
func (art *articleRepo) articleDetailKey(articleID int) string {
	return fmt.Sprintf("%s%d", articleDetailKeyPrefix, articleID)
}

// GetArticleDetailCache 从 Redis 获取文章详情缓存
func (art *articleRepo) GetArticleDetailCache(articleID int) (*entity.Article, error) {
	ctx := context.Background()
	data, err := art.rdb.Get(ctx, art.articleDetailKey(articleID)).Bytes()
	if err != nil {
		return nil, err
	}

	var article entity.Article
	if err := json.Unmarshal(data, &article); err != nil {
		return nil, fmt.Errorf("反序列化文章详情失败: %w", err)
	}
	return &article, nil
}

// SetArticleDetailCache 将文章详情缓存到 Redis（过期时间 30 分钟）
func (art *articleRepo) SetArticleDetailCache(articleID int, article *entity.Article) error {
	ctx := context.Background()
	data, err := json.Marshal(article)
	if err != nil {
		return fmt.Errorf("序列化文章详情失败: %w", err)
	}
	return art.rdb.Set(ctx, art.articleDetailKey(articleID), data, 30*time.Minute).Err()
}
