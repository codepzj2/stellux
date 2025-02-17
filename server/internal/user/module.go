package user

import (
	"server/internal/user/internal/api"
	"server/internal/user/internal/service"
)

type (
	Handler = api.UserHandler
	Service = service.IUserService
	Module  struct {
		Svc Service
		Hdl *Handler
	}
)
