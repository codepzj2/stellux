//go:build wireinject
// +build wireinject

package app

import (
	"github.com/codepzj/stellux/server/internal/ioc"

	"github.com/codepzj/stellux/server/internal/post"

	"github.com/google/wire"
)

func InitApp() *HttpServer {
	wire.Build(
		ioc.InitMiddleWare,
		ioc.NewGin,
		ioc.NewMongoDB,

		post.InitPostModule,
		wire.FieldsOf(new(*post.Module), "Hdl"),

		NewHttpServer,
	)
	return nil
}
