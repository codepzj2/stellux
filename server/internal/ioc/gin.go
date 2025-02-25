package ioc

import (
	"server/internal/file"
	"server/internal/posts"
	"server/internal/user"

	"github.com/gin-gonic/gin"
)

// NewGin 初始化gin服务器
func NewGin(usrHdl *user.Handler, postsHdl *posts.Handler, fileHdl *file.Handler, middleware []gin.HandlerFunc) *gin.Engine {
	router := gin.New()

	// 中间件
	router.Use(middleware...)

	// 初始化路由
	{
		usrHdl.RegisterGinRoutes(router)
		postsHdl.RegisterGinRoutes(router)
		fileHdl.RegisterGinRoutes(router)
	}

	return router
}
