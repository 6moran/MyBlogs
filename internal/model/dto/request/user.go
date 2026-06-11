package request

// CreateUserRequest 创建用户请求
type CreateUserRequest struct {
	Username string `json:"username" binding:"required,min=3,max=50"`
	Password string `json:"password" binding:"required,min=6,max=50"`
	Nickname string `json:"nickname" binding:"required,min=1,max=50"`
	Email    string `json:"email" binding:"omitempty,email"`
	Role     int    `json:"role" binding:"required,oneof=0 1"`
}

// UpdateUserRequest 更新用户请求
type UpdateUserRequest struct {
	Nickname string `json:"nickname" binding:"omitempty,min=1,max=50"`
	Email    string `json:"email" binding:"omitempty,email"`
	Status   *int   `json:"status" binding:"omitempty,oneof=0 1"`
	Role     *int   `json:"role" binding:"omitempty,oneof=0 1"`
	Password string `json:"password" binding:"omitempty,min=6,max=50"`
}

// QueryUserRequest 查询用户请求
type QueryUserRequest struct {
	Page    int    `json:"page" form:"page" binding:"omitempty,min=1"`
	Size    int    `json:"size" form:"size" binding:"omitempty,min=1,max=100"`
	KeyWord string `json:"key_word" form:"key_word"`
	Status  *int   `json:"status" form:"status" binding:"omitempty,oneof=0 1"`
	Role    *int   `json:"role" form:"role" binding:"omitempty,oneof=0 1"`
}
