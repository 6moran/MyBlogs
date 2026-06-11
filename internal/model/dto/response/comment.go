package response

import "time"

type CommentResponse struct {
	ID        int       `json:"id"`
	ParentID  int       `json:"parent_id"`
	RootID    int       `json:"root_id"`
	UserID    int       `json:"user_id"`
	Username  string    `json:"username"`
	Nickname  string    `json:"nickname"`
	Avatar    string    `json:"avatar"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

type CommentListResponse struct {
	Comments []CommentResponse `json:"comments"`
	Total    int64             `json:"total"`
	Page     int               `json:"page"`
	Size     int               `json:"size"`
}
