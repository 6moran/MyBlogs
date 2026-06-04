package request

// 创建文章请求体
type CreateArticleRequest struct {
	UserID     int    `json:"user_id"`
	CategoryID int    `json:"category_id"`
	Title      string `json:"title"`
	Summary    string `json:"summary"`
	CoverImage string `json:"cover_image"`
	Content    string `json:"content"`
	Status     int    `json:"status"`
}

type QueryArticleRequest struct {
	QueryRequest
	Status     int   `json:"status" form:"status"`           //1-草稿 2-已发布 3-私密
	CategoryID int   `json:"category_id" form:"category_id"` //分类ID
	TagsID     []int `json:"tags_id" form:"tags_id"`         //标签ID
	SortWay    int   `json:"sort_way" form:"sort_way"`       //0-相关性 1-最新 2-最热
}
