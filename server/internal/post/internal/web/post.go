package web

import (
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
	postGroup := engine.Group("/post")
	{
		postGroup.GET("/list", apiwrap.WrapWithBody(h.GetPostList))
		postGroup.GET("/:id", apiwrap.Wrap(h.GetPostById))
	}
}

func (h *PostHandler) AdminCreatePost(c *gin.Context, postReq PostRequest) *apiwrap.Response[any] {
	post := h.DTOToDomain(postReq)
	err := h.serv.AdminCreatePost(c, post)
	if err != nil {
		return apiwrap.FailWithMsg(apiwrap.RuquestInternalServerError, err.Error())
	}
	return apiwrap.SuccessWithMsg("创建文章成功")
}

func (h *PostHandler) GetPostList(c *gin.Context, pageReq apiwrap.Page) *apiwrap.Response[any] {
	page := domain.Page{
		PageNo:   pageReq.PageNo,
		PageSize: pageReq.PageSize,
		Keyword:  pageReq.Keyword,
		Field:    pageReq.Field,
		Order:    pageReq.Order,
	}
	posts, total, err := h.serv.GetPostList(c, &page)
	if err != nil {
		return apiwrap.FailWithMsg(apiwrap.RuquestInternalServerError, err.Error())
	}
	postVos := h.DomainToVOList(posts)

	pageVo := apiwrap.ToPageVO(page.PageNo, page.PageSize, total, postVos)
	return apiwrap.SuccessWithDetail[any](pageVo, "获取文章列表成功")
}

func (h *PostHandler) GetPostById(c *gin.Context) *apiwrap.Response[any] {
	id := c.Param("id")
	post, err := h.serv.GetPostById(c, id)
	if err != nil {
		return apiwrap.FailWithMsg(apiwrap.RuquestInternalServerError, err.Error())
	}
	return apiwrap.SuccessWithDetail[any](h.DomainToVO(post), "获取文章详情成功")
}

func (h *PostHandler) DTOToDomain(postReq PostRequest) *domain.Post {
	return &domain.Post{
		Title:       postReq.Title,
		Content:     postReq.Content,
		Description: postReq.Description,
		Author:      postReq.Author,
		Category:    postReq.Category,
		Tags:        postReq.Tags,
		IsPublish:   postReq.IsPublish,
		IsTop:       postReq.IsTop,
		Thumbnail:   postReq.Thumbnail,
	}
}

func (h *PostHandler) DomainToVO(post *domain.Post) *PostVO {
	return &PostVO{
		ID:          post.ID,
		CreatedAt:   post.CreatedAt,
		Title:       post.Title,
		Content:     post.Content,
		Description: post.Description,
		Author:      post.Author,
		Category:    post.Category,
		Tags:        post.Tags,
		IsPublish:   post.IsPublish,
		IsTop:       post.IsTop,
		Thumbnail:   post.Thumbnail,
	}
}

func (h *PostHandler) DomainToVOList(posts []*domain.Post) []*PostVO {
	voList := make([]*PostVO, len(posts))
	for i, post := range posts {
		voList[i] = h.DomainToVO(post)
	}
	return voList
}
