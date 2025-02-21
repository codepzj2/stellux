//go:build wireinject
// +build wireinject

package posts

import (
	"github.com/chenmingyong0423/go-mongox/v2"
	"github.com/google/wire"
	"server/internal/posts/internal/api"
	"server/internal/posts/internal/repo"
	"server/internal/posts/internal/repo/dao"
	"server/internal/posts/internal/service"
)

var postsProvider = wire.NewSet(
	api.NewPostHandler,
	service.NewPostsService,
	repo.NewPostsRepo,
	dao.NewPostsDao,
	wire.Bind(new(service.IPostsService), new(*service.PostsService)),
	wire.Bind(new(repo.IPostsRepo), new(*repo.PostsRepo)),
	wire.Bind(new(dao.IPostsDao), new(*dao.PostsDao)),
	wire.Struct(new(Module), "Hdl", "Svc"),
)

func InitPostsModule(database *mongox.Database) *Module {
	wire.Build(
		postsProvider,
	)
	return nil
}
