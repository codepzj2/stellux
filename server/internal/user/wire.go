//go:build wireinject
// +build wireinject

package user

import (
	"github.com/chenmingyong0423/go-mongox/v2"
	"github.com/google/wire"
	"server/internal/user/internal/api"
	"server/internal/user/internal/repo"
	"server/internal/user/internal/repo/dao"
	"server/internal/user/internal/service"
)

var userProvider = wire.NewSet(api.NewUserHandler, service.NewUserService, repo.NewUserRepo, dao.NewUserDao,
	wire.Struct(new(Module), "Hdl", "Svc"),
	wire.Bind(new(service.IUserService), new(*service.UserService)),
	wire.Bind(new(repo.IUserRepo), new(*repo.UserRepo)),
	wire.Bind(new(dao.IUserDao), new(*dao.UserDao)))

func InitUserModule(db *mongox.Database) *Module {
	wire.Build(
		userProvider,
	)
	return nil
}
