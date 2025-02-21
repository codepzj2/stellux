package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/internal/pkg/http/resp"
	"server/internal/posts/internal/service"
)

type IPostHandler interface {
	CreatePosts(ctx *gin.Context)
}
type PostsHandler struct {
	serv service.IPostsService
}

func NewPostHandler(serv service.IPostsService) *PostsHandler {
	return &PostsHandler{serv}
}

func (h *PostsHandler) RegisterGinRoutes(router *gin.Engine) {
	group := router.Group("/posts")
	{
		group.POST("/create", h.CreatePosts)
	}
}

func (h *PostsHandler) CreatePosts(ctx *gin.Context) {
	var postsReq PostsReq
	if ctx.ShouldBindJSON(&postsReq) != nil {
		resp.FailWithMsg(ctx, http.StatusBadRequest, "参数错误")
		return
	}
	err := h.serv.CreatePosts(ctx, toPosts(postsReq))
	if err != nil {
		resp.FailWithMsg(ctx, http.StatusBadRequest, err.Error())
		return
	}
	resp.SuccessWithMsg(ctx, "新增文章成功")
}
