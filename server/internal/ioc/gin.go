package ioc

import (
	"github.com/codepzj/stellux/server/internal/label"
	"github.com/codepzj/stellux/server/internal/post"
	"github.com/codepzj/stellux/server/internal/user"

	"github.com/gin-gonic/gin"
)

// NewGin 初始化gin服务器
func NewGin(userHdl *user.Handler, postHdl *post.Handler, labelHdl *label.Handler, middleware []gin.HandlerFunc) *gin.Engine {
	router := gin.New()

	// 中间件
	router.Use(middleware...)

	// 初始化路由
	{
		userHdl.RegisterGinRoutes(router)
		postHdl.RegisterGinRoutes(router)
		labelHdl.RegisterGinRoutes(router)
	}

	return router
}
