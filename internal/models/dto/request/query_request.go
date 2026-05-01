package request

type QueryRequest struct {
	Page    int    `json:"page"`     //第几页
	Size    int    `json:"size"`     //每页的规模
	KeyWord string `json:"key_word"` //搜索条件
}
