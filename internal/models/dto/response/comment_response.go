package response

import "time"

type CommentResponse struct {
	ID        int       `json:"id"`
	ParentID  int       `json:"parent_id"`
	RootID    int       `json:"root_id"`
	UserID    int       `json:"user_id"`
	Content   string    `json:"content"`
	LikeCount int       `json:"like_count"`
	CreateAt  time.Time `json:"create_at"`
}
