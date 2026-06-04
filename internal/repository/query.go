package repository

// 基础查询
type QueryLimit struct {
	Page    int
	Size    int
	Keyword string
}

// 查询文章
type ArticleQuery struct {
	QueryLimit
	Status     int //草稿、已发布
	SortWay    int //0-相关性 1-最新 2-最热
	CategoryID int
	TagsID     []int
}
