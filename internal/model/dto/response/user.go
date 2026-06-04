package response

import "time"

type UserResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	NickName  string    `json:"nick_name"`
	Avatar    string    `json:"avatar"`
	Status    int       `json:"status"`
	Role      int       `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}
