package file

import (
	"server/internal/file/internal/api"
	"server/internal/file/internal/service"
)

type (
	Handler = api.FileHandler
	Service = service.IFileService
	Module  struct {
		Hdl *Handler
		Svc Service
	}
)