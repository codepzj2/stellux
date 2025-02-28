package domain

import (
	"github.com/chenmingyong0423/go-mongox/v2"
	"github.com/samber/lo"
)

// User 用户表 RoleId 为 0 时为管理员，为 1 时为普通用户
type User struct {
	mongox.Model `bson:",inline"`
	Username     string `bson:"username"`
	Password     string `bson:"password"`
	RoleId       int    `bson:"role_id"`
}

func ToUserPtrSlice(users []User) []*User {
	return lo.Map(users, func(item User, index int) *User {
		return &item
	})
}
