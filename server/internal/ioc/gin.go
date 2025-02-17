package ioc

import (
	"github.com/gin-gonic/gin"
	"server/internal/user"
)

// NewGin 初始化gin服务器
func NewGin(usrHdl *user.Handler, middleware []gin.HandlerFunc) *gin.Engine {
	router := gin.New()

	// 中间件
	router.Use(middleware...)

	// 初始化路由
	{
		usrHdl.RegisterGinRoutes(router)
	}

	return router
}
