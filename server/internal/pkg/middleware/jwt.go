package middleware

import (
	"log"
	"net/http"
	"server/internal/pkg/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 放行GET请求
		if ctx.Request.Method == "GET" {
			ctx.Next()
			return
		}	

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
		log.Println("用户携带的token:", token)
		if token == "" || !strings.HasPrefix(token, "Bearer ") {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		claims, err := utils.ParseJwt(strings.TrimPrefix(token, "Bearer "))
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		log.Println("解析后的claims:", claims)
		ctx.Set("userId", claims.ID)
		ctx.Next()
	}
}
