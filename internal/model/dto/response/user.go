package response

import "time"

type UserResponse struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Nickname  string    `json:"nickname"`
	Avatar    string    `json:"avatar"`
	Email     string    `json:"email"`
	Status    int       `json:"status"`
	Role      int       `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}
