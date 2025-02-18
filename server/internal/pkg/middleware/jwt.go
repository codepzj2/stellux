package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/internal/pkg/utils"
	"strings"
)

func JWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 排除接口
		excludeUri := []string{
			"/user/login",
		}
		for _, uri := range excludeUri {
			if ctx.Request.RequestURI == uri {
				ctx.Next()
				return
			}
		}
		token := ctx.Request.Header.Get("Authorization")
		if token == "" || !strings.HasPrefix(token, "Bearer ") {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		claims, err := utils.ParseJwt(strings.TrimPrefix(token, "Bearer "))
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		ctx.Set("userId", claims.ID)
		ctx.Next()
	}
}
