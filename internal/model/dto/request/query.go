package request

type QueryRequest struct {
	Page    int    `json:"page" form:"page"`         //第几页
	Size    int    `json:"size" form:"size"`         //每页的规模
	KeyWord string `json:"key_word" form:"key_word"` //搜索条件
}
