package web

type UserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdateUserRequest struct {
	ID       string `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	RoleId   int    `json:"role_id"`
	Avatar   string `json:"avatar,omitempty"`
	Email    string `json:"email,omitempty"`
	Sex      string `json:"sex,omitempty"`
	Company  string `json:"company,omitempty"`
	Hobby    string `json:"hobby,omitempty"`
}
