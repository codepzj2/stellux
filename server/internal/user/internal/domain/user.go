package domain

import "github.com/chenmingyong0423/go-mongox/v2"

// User 用户表 RoleId 为 0 时为管理员，为 1 时为普通用户
type User struct {
	mongox.Model `bson:",inline"`
	Username     string `bson:"username"`
	Password     string `bson:"password"`
	RoleId       int64  `bson:"role_id"`
}
