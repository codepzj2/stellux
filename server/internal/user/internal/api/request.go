package api

import "server/internal/user/internal/domain"

type LoginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type CreateUserReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	RoleId   int    `json:"role_id" binding:"required"`
}

func toUser(loginReq LoginReq) domain.User {
	return domain.User{
		Username: loginReq.Username,
		Password: loginReq.Password,
	}
}
