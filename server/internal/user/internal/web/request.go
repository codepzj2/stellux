package web

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Nickname string `json:"nickname"`
	RoleId   *int   `json:"role_id" binding:"required"`
	Avatar   string `json:"avatar"`
	Email    string `json:"email"`
}

type UpdateUserRequest struct {
	ID       string `json:"id" binding:"required"`
	Username string `json:"username" binding:"required"`
	Nickname string `json:"nickname"`
	RoleId   *int   `json:"role_id" binding:"required"`
	Avatar   string `json:"avatar"`
	Email    string `json:"email"`
}
