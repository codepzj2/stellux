package api

type LoginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type CreateUserReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	RoleId   int    `json:"role_id" binding:"required"`
}

