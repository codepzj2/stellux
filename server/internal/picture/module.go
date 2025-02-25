package picture

import (
	"server/internal/picture/internal/api"
)

type (
	Handler = api.PictureHandler
	Module  struct {
		Hdl *Handler
	}
)
