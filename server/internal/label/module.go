package label

import (
	"github.com/codepzj/stellux/server/internal/label/internal/service"
	"github.com/codepzj/stellux/server/internal/label/internal/web"
)

type (
	Handler = web.LabelHandler
	Service = service.ILabelService
	Module  struct {
		Svc Service
		Hdl *Handler
	}
)
