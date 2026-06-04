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
	Status       int              `json:"status"`
	CommentCount int              `json:"comment_count"`
	Category     CategoryResponse `json:"category"`
	Tags         []TagResponse    `json:"tags"`
}

type ArticleDetailResponse struct {
	ArticleResponse
	Content      string           `json:"content"`
	CommentCount int              `json:"comment_count"`
	Category     CategoryResponse `json:"category"`
	Tags         []TagResponse    `json:"tags"`
	//一次文章请求只返回前三条
	Comments []CommentResponse `json:"comments"`
}

// ArticleAdminDetailResponse 后台文章详细页面
type ArticleAdminDetailResponse struct {
	ArticleResponse
	Content string `json:"content"`
}
