package response

type ArticleResponse struct {
	ID              int    `json:"id"`
	CreatedAt       string `json:"created_at"`
	UpdateAt        string `json:"update_at"`
	Title           string `json:"title"`
	Summary         string `json:"summary"`
	CoverImage      string `json:"cover_image"`
	Content         string `json:"content"`
	Status          int    `json:"status"`
	LikeCount       int    `json:"like_count"`
	ViewCount       int    `json:"view_count"`
	CollectionCount int    `json:"collection_count"`
	CommentCount    int    `json:"comment_count"`

	User     UserResponse     `json:"user"`
	Category CategoryResponse `json:"category"`
	Tags     []TagResponse    `json:"tags"`
}

// 评论单独响应
