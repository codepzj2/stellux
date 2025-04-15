package web

import (
	"net/http"

	"github.com/codepzj/stellux/server/internal/pkg/apiwrap"
	"github.com/codepzj/stellux/server/internal/post/internal/domain"
	"github.com/codepzj/stellux/server/internal/post/internal/service"
	"github.com/gin-gonic/gin"
)

func NewPostHandler(serv service.IPostService) *PostHandler {
	return &PostHandler{
		serv: serv,
	}
}

type PostHandler struct {
	serv service.IPostService
}

func (h *PostHandler) RegisterGinRoutes(engine *gin.Engine) {
	adminGroup := engine.Group("/admin-api/post")
	{
		adminGroup.POST("create", apiwrap.WrapWithBody(h.AdminCreatePost))
	}
}

func (h *PostHandler) AdminCreatePost(c *gin.Context, postReq PostRequest) (*apiwrap.Response[any], error) {
	post := h.DTOToDomain(postReq)
	err := h.serv.AdminCreatePost(c, post)
	if err != nil {
		return apiwrap.FailWithMsg(http.StatusInternalServerError, err.Error()), err
	}
	return apiwrap.SuccessWithMsg[any](nil, "创建文章成功"), nil
}

func (h *PostHandler) DTOToDomain(postReq PostRequest) *domain.Post {
	return &domain.Post{
		Title:       postReq.Title,
		Content:     postReq.Content,
		Description: postReq.Description,
		Author:      postReq.Author,
		Category:    postReq.Category,
		Tags:        postReq.Tags,
		IsPublished: postReq.IsPublished,
		IsTop:       postReq.IsTop,
		Thumbnail:   postReq.Thumbnail,
	}
}
