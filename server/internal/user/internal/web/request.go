package web

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Nickname string `json:"nickname,omitempty"`
	RoleId   *int   `json:"role_id" binding:"required"`
	Avatar   string `json:"avatar,omitempty"`
	Email    string `json:"email,omitempty"`
	Sex      string `json:"sex,omitempty"`
	Hobby    string `json:"hobby,omitempty"`
}

type UpdateUserRequest struct {
	ID       string `json:"id,omitempty" binding:"required"`
	Username string `json:"username,omitempty" binding:"required"`
	Nickname string `json:"nickname,omitempty"`
	RoleId   *int   `json:"role_id" binding:"required"`
	Avatar   string `json:"avatar,omitempty"`
	Email    string `json:"email,omitempty"`
	Sex      string `json:"sex,omitempty"`
	Hobby    string `json:"hobby,omitempty"`
}
