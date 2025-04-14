package ioc

import (
	"github.com/codepzj/stellux/server/internal/post"

	"github.com/gin-gonic/gin"
)

// NewGin 初始化gin服务器
func NewGin(postHdl *post.Handler, middleware []gin.HandlerFunc) *gin.Engine {
	router := gin.New()

	// 中间件
	router.Use(middleware...)

	// 初始化路由
	{
		postHdl.RegisterGinRoutes(router)
	}

	return router
}
