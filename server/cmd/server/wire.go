//go:build wireinject
// +build wireinject

package main

import (
	"server/internal/ioc"
	"server/internal/picture"
	"server/internal/posts"
	"server/internal/user"

	"github.com/google/wire"
)

func InitApp() *HttpServer {
	wire.Build(
		ioc.NewMongoDB,
		ioc.InitMiddleWare,
		ioc.NewGin,

		user.InitUserModule,
		wire.FieldsOf(new(*user.Module), "Hdl"),

		posts.InitPostsModule,
		wire.FieldsOf(new(*posts.Module), "Hdl"),

		picture.InitPictureModule,
		wire.FieldsOf(new(*picture.Module), "Hdl"),

		NewHttpServer,
	)
	return nil
}
