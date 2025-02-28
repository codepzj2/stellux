//go:build wireinject
// +build wireinject

package user

import (
	"server/internal/user/internal/api"
	"server/internal/user/internal/repo"
	"server/internal/user/internal/repo/dao"
	"server/internal/user/internal/service"

	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var userProvider = wire.NewSet(api.NewUserHandler, service.NewUserService, repo.NewUserRepo, dao.NewUserDao,
	wire.Bind(new(service.IUserService), new(*service.UserService)),
	wire.Bind(new(repo.IUserRepo), new(*repo.UserRepo)),
	wire.Bind(new(dao.IUserDao), new(*dao.UserDao)))

func InitUserModule(database *mongo.Database) *Module {
	panic(wire.Build(
		userProvider,
		wire.Struct(new(Module), "Hdl", "Svc"),
	))
}
