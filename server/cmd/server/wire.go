//go:build wireinject
// +build wireinject

package main

import (
	"server/internal/ioc"
	"server/internal/user"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func InitApp() *gin.Engine {
	wire.Build(
		ioc.InitEnv,
		ioc.NewMongoDB,
		ioc.InitMiddleWare,
		ioc.NewGin,

		user.InitUserModule,
		wire.FieldsOf(new(*user.Module), "Hdl"),
	)
	return nil
}
