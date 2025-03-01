package dao_test

import (
	"context"
	"fmt"
	"server/internal/ioc"
	"server/internal/user/internal/repo/dao"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserDao_FindAllByRoleID(t *testing.T) {
	db := ioc.NewMongoDB()
	userDao := dao.NewUserDao(db)
	users, err := userDao.FindAllByRoleID(context.Background(), 1)
	for _, user := range users {
		fmt.Println(*user)
	}
	assert.NoError(t, err)
}
