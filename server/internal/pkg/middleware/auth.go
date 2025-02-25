package middleware

import (
	"log"
	"net/http"
	"server/global"
	"server/internal/pkg/http/resp"
	"sync"

	"github.com/casbin/casbin/v2"
	mongodbadapter "github.com/casbin/mongodb-adapter/v3"
	"github.com/gin-gonic/gin"
	mongooptions "go.mongodb.org/mongo-driver/mongo/options"
)

var (
	once     sync.Once
	enforcer *casbin.Enforcer
)

// InitCasbin 确保在使用中间件时enforcer被初始化至内存中
func InitCasbin() *casbin.Enforcer {
	once.Do(func() {
		mongoClientOption := mongooptions.Client().ApplyURI(global.Env.URL)
		databaseName := global.Env.MongoDatabase
		a, err := mongodbadapter.NewAdapterWithClientOption(mongoClientOption, databaseName)
		if err != nil {
			panic(err)
		}
		enforcer, err = casbin.NewEnforcer("config/policy.conf", a)
		if err != nil {
			panic(err)
		}
		enforcer.LoadPolicy()
	})

	return enforcer
}

func Auth() gin.HandlerFunc {
	enforcer = InitCasbin()
	return func(ctx *gin.Context) {
		userId := ctx.GetString("userId")
		requestURI := ctx.Request.RequestURI
		method := ctx.Request.Method
		ok, err := enforcer.Enforce(userId, requestURI, method)
		if err != nil {
			resp.FailWithMsg(ctx, http.StatusInternalServerError, "权限校验失败")
			ctx.Abort()
			return
		}
		if ok {
			log.Println("userId为:", userId, "允许访问")
			ctx.Next()
		} else {
			log.Println("userId为:", userId, "禁止访问")
			ctx.AbortWithStatus(http.StatusForbidden)
		}
	}
}
