//go:build wireinject
// +build wireinject

package picture

import (
	"server/internal/picture/internal/api"

	"github.com/google/wire"
)

func InitPictureModule() *Module {
	wire.Build(
		api.NewPictureHandler,
		wire.Struct(new(Module), "Hdl"),
	)
	return nil
}
