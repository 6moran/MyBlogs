package response

import "time"

type ArticleResponse struct {
	ID         int       `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Title      string    `json:"title"`
	Summary    string    `json:"summary"`
	CoverImage string    `json:"cover_image"`

	LikeCount int `json:"like_count"`
	ViewCount int `json:"view_count"`
}

// ArticleAdminResponse 后台文章列表响应体
type ArticleAdminResponse struct {
	ArticleResponse
	Status       int           `json:"status"`
	CommentCount int           `json:"comment_count"`
	Tags         []TagResponse `json:"tags"`
}

type ArticleDetailResponse struct {
	ArticleResponse
	Content      string            `json:"content"`
	CommentCount int               `json:"comment_count"`
	Tags         []TagResponse     `json:"tags"`
	Comments     []CommentResponse `json:"comments"`
}

// ArticleAdminDetailResponse 后台文章详细页面
type ArticleAdminDetailResponse struct {
	ArticleResponse
	Content string `json:"content"`
}

// ArticleListResponse 文章列表响应
type ArticleListResponse struct {
	Articles []ArticleListItemResponse `json:"articles"`
	Total    int64                     `json:"total"`
	Page     int                       `json:"page"`
	Size     int                       `json:"size"`
}

// ArticleListItemResponse 文章列表项响应
type ArticleListItemResponse struct {
	ID           int           `json:"id"`
	Title        string        `json:"title"`
	Summary      string        `json:"summary"`
	CoverImage   string        `json:"cover_image"`
	ViewCount    int           `json:"view_count"`
	LikeCount    int           `json:"like_count"`
	CommentCount int           `json:"comment_count"`
	CreatedAt    time.Time     `json:"created_at"`
	Tags         []TagResponse `json:"tags"`
}
