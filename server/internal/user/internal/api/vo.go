package api

import (
	"server/internal/user/internal/domain"
)

type UserVO struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	RoleId   int    `json:"role_id"`
}

type LoginVO struct {
	User  UserVO `json:"user"`
	Token string `json:"token"`
}
type UserListVO struct {
	Id        string `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	RoleId    int    `json:"role_id"`
}

// domain.User转UserVO
func toUserVO(user *domain.User) UserVO {
	return UserVO{
		Id:       user.ID.Hex(),
		Username: user.Username,
		RoleId:   int(user.RoleId),
	}
}

func toUserListVO(users []*domain.User) []UserListVO {
	userListVO := make([]UserListVO, len(users))
	for i, u := range users {
		userListVO[i] = UserListVO{
			Id:        u.ID.Hex(),
			CreatedAt: u.CreatedAt.String(),
			UpdatedAt: u.UpdatedAt.String(),
			Username:  u.Username,
			Password:  u.Password,
			RoleId:    u.RoleId,
		}
	}
	return userListVO
}
