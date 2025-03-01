package dao_test

import (
	"context"
	"fmt"
	"server/internal/file/internal/repo/dao"
	"server/internal/ioc"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FindByPage(t *testing.T) {
	db := ioc.NewMongoDB()
	fileDao := dao.NewFileDao(db)
	files, err := fileDao.FindListByPage(context.TODO(), 1, 2)
	assert.NoError(t, err)
	for _, file := range files {
		fmt.Println(*file)
	}
}
