package api

import "server/internal/user/internal/domain"

type UserVO struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	RoleId   int    `json:"role_id"`
}

type LoginVO struct {
	User  UserVO `json:"user"`
	Token string `json:"token"`
}

// domain.User转UserVO
func toUserVO(user *domain.User) UserVO {
	return UserVO{
		Id:       user.ID.Hex(),
		Username: user.Username,
		RoleId:   int(user.RoleId),
	}
}
