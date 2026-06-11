package request

// 创建文章请求体
type CreateArticleRequest struct {
	UserID     int    `json:"user_id"`
	Title      string `json:"title"`
	Summary    string `json:"summary"`
	CoverImage string `json:"cover_image"`
	Content    string `json:"content"`
	Status     int    `json:"status"`
}

type QueryArticleRequest struct {
	QueryRequest
	Status  int   `json:"status" form:"status"`     //1-草稿 2-已发布 3-私密
	TagsID  []int `json:"tags_id" form:"tags_id"`   //标签ID
	SortWay int   `json:"sort_way" form:"sort_way"` //0-相关性 1-最新 2-最热
}

// PublishedArticleRequest 前台已发布文章查询请求
type PublishedArticleRequest struct {
	Page    int    `json:"page" form:"page"`
	Size    int    `json:"size" form:"size"`
	TagID   int    `json:"tag_id" form:"tag_id"`
	Keyword string `json:"keyword" form:"keyword"`
}
