package ioc

import (
	"github.com/gin-gonic/gin"
	"server/internal/pkg/middleware"
)

func InitMiddleWare() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		gin.Recovery(),
		middleware.GinLogger(),
	}
}
