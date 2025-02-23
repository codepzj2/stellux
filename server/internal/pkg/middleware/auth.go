package middleware

import (
	"server/global"

	"github.com/casbin/casbin/v2"
	mongodbadapter "github.com/casbin/mongodb-adapter/v3"
	"github.com/gin-gonic/gin"
	mongooptions "go.mongodb.org/mongo-driver/mongo/options"
)

func InitCasbin() *casbin.Enforcer {
	mongoClientOption := mongooptions.Client().ApplyURI(global.Env.URL)
	databaseName := global.Env.MongoDatabase
	a, err := mongodbadapter.NewAdapterWithClientOption(mongoClientOption, databaseName)
	if err != nil {
		panic(err)
	}
	e, err := casbin.NewEnforcer("config/policy.conf", a)
	if err != nil {
		panic(err)
	}
	e.LoadPolicy()
	return e
}

func Auth() gin.HandlerFunc {
	e := InitCasbin()
	e.AddPolicy("admin", "/api/v1/user", "GET")

	e.AddRoleForUser("78", "admin")
	e.SavePolicy()
	return func(ctx *gin.Context) {
		ctx.Next()
	}
}
