package middleware

import (
	"log"
	"server/global"
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

// 确保在使用中间件时enforcer被初始化至内存中
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
		log.Println("casbin获取的userId为:", userId)
		ctx.Next()
	}
}
