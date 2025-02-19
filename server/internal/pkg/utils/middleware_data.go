package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func getUserId(ctx *gin.Context) (string, error) {
	val, isExist := ctx.Get("userId")
	if !isExist {
		return "", errors.New("userId不存在")
	}
	return val.(string), nil
}
