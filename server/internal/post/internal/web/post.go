package web

import (
	"github.com/codepzj/stellux/server/internal/pkg/apiwrap"
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
		adminGroup.PUT("update", apiwrap.WrapWithBody(h.AdminUpdatePost))
	}
	postGroup := engine.Group("/post")
	{
		postGroup.GET("/detail/list", apiwrap.WrapWithBody(h.GetDetailPostList))
		postGroup.GET("/detail/:id", apiwrap.WrapWithUri(h.GetDetailPostById))
		postGroup.GET("/:id", apiwrap.WrapWithUri(h.GetPostById))
	}
}

func (h *PostHandler) AdminCreatePost(c *gin.Context, postReq PostRequest) *apiwrap.Response[any] {
	post := h.PostDTOToDomain(postReq)
	err := h.serv.AdminCreatePost(c, post)
	if err != nil {
		return apiwrap.FailWithMsg(apiwrap.RuquestInternalServerError, err.Error())
	}
	return apiwrap.SuccessWithMsg("创建文章成功")
}

func (h *PostHandler) AdminUpdatePost(c *gin.Context, postUpdateReq PostUpdateRequest) *apiwrap.Response[any] {
	postUpdate := h.PostUpdateDTOToDomain(postUpdateReq)
	err := h.serv.AdminUpdatePost(c, postUpdate)
	if err != nil {
		return apiwrap.FailWithMsg(apiwrap.RuquestInternalServerError, err.Error())
	}
	return apiwrap.SuccessWithMsg("更新文章成功")
}

func (h *PostHandler) GetDetailPostList(c *gin.Context, pageReq apiwrap.Page) *apiwrap.Response[any] {
	postDetailList, total, err := h.serv.GetPostDetailList(c, &pageReq)
	if err != nil {
		return apiwrap.FailWithMsg(apiwrap.RuquestInternalServerError, err.Error())
	}
	postVos := h.PostDetailListToVOList(postDetailList)

	pageVo := apiwrap.ToPageVO(pageReq.PageNo, pageReq.PageSize, total, postVos)
	return apiwrap.SuccessWithDetail[any](pageVo, "获取文章列表成功")
}

func (h *PostHandler) GetDetailPostById(c *gin.Context, postIDRequest PostIDRequest) *apiwrap.Response[any] {
	postDetail, err := h.serv.GetPostDetailById(c, apiwrap.ConvertBsonID(postIDRequest.ID))
	if err != nil {
		return apiwrap.FailWithMsg(apiwrap.RuquestInternalServerError, err.Error())
	}
	return apiwrap.SuccessWithDetail[any](h.PostDetailToVO(postDetail), "获取文章详情成功")
}

func (h *PostHandler) GetPostById(c *gin.Context, postIDRequest PostIDRequest) *apiwrap.Response[any] {
	post, err := h.serv.GetPostById(c, apiwrap.ConvertBsonID(postIDRequest.ID))
	if err != nil {
		return apiwrap.FailWithMsg(apiwrap.RuquestInternalServerError, err.Error())
	}
	return apiwrap.SuccessWithDetail[any](h.PostToVO(post), "获取文章成功")
}
