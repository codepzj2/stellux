package web

import "github.com/codepzj/stellux/server/internal/user/internal/domain"

type UserVO struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	RoleId   int    `json:"role_id"`
	Avatar   string `json:"avatar"`
	Email    string `json:"email"`
	Sex      string `json:"sex"`
	Hobby    string `json:"hobby"`
}

type LoginVO struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (h *UserHandler) DomainToVO(user *domain.User) *UserVO {
	return &UserVO{
		ID:       user.ID,
		Username: user.Username,
		Nickname: user.Nickname,
		RoleId:   user.RoleId,
		Avatar:   user.Avatar,
		Email:    user.Email,
		Sex:      user.Sex,
		Hobby:    user.Hobby,
	}
}

func (h *UserHandler) DomainToVOList(users []*domain.User) []*UserVO {
	voList := make([]*UserVO, len(users))
	for i, user := range users {
		voList[i] = h.DomainToVO(user)
	}
	return voList
}
