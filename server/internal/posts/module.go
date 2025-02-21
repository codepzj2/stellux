package posts

import (
	"server/internal/posts/internal/api"
	"server/internal/posts/internal/service"
)

type (
	Handler = api.PostsHandler
	Service = service.IPostsService
	Module  struct {
		Hdl *Handler
		Svc Service
	}
)
