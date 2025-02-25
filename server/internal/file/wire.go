//go:build wireinject
// +build wireinject

package file

import (
	"github.com/chenmingyong0423/go-mongox/v2"
	"server/internal/file/internal/api"
	"server/internal/file/internal/repo"
	"server/internal/file/internal/repo/dao"
	"server/internal/file/internal/service"

	"github.com/google/wire"
)

func InitFileModule(database *mongox.Database) *Module {
	wire.Build(
		api.NewFileHandler,
		service.NewFileService,
		repo.NewFileRepo,
		dao.NewFileDao,
		wire.Bind(new(service.IFileService), new(*service.FileService)),
		wire.Bind(new(repo.IFileRepo), new(*repo.FileRepo)),
		wire.Bind(new(dao.IFileDao), new(*dao.FileDao)),
		wire.Struct(new(Module), "Hdl", "Svc"),
	)
	return nil
}
