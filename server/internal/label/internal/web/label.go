package web

import (
	"github.com/codepzj/stellux/server/internal/label/internal/domain"
	"github.com/codepzj/stellux/server/internal/label/internal/service"
	"github.com/codepzj/stellux/server/internal/pkg/apiwrap"
	"github.com/gin-gonic/gin"
)

func NewLabelHandler(serv service.ILabelService) *LabelHandler {
	return &LabelHandler{
		serv: serv,
	}
}

type LabelHandler struct {
	serv service.ILabelService
}

func (h *LabelHandler) RegisterGinRoutes(engine *gin.Engine) {
	adminGroup := engine.Group("/admin-api/label")
	{
		adminGroup.POST("/create", apiwrap.WrapWithBody(h.AdminCreate))
		adminGroup.PUT("/edit", apiwrap.WrapWithBody(h.AdminUpdate))
		adminGroup.DELETE("/delete/:id", apiwrap.Wrap(h.AdminDelete))
	}
	labelGroup := engine.Group("/label")
	{
		labelGroup.GET("/:id", apiwrap.Wrap(h.GetByID))
		labelGroup.GET("/list", apiwrap.WrapWithBody(h.GetList))
	}
}

func (h *LabelHandler) AdminCreate(c *gin.Context, label *LabelRequest) *apiwrap.Response[any] {
	err := h.serv.Create(c, &domain.Label{
		LabelType: label.LabelType,
		Name:      label.Name,
	})
	if err != nil {
		return apiwrap.FailWithMsg(apiwrap.RuquestInternalServerError, err.Error())
	}
	return apiwrap.SuccessWithMsg("标签创建成功")
}

func (h *LabelHandler) AdminUpdate(c *gin.Context, label *LabelRequest) *apiwrap.Response[any] {
	err := h.serv.Update(c, label.ID, h.DTOToDomain(label))
	if err != nil {
		return apiwrap.FailWithMsg(apiwrap.RuquestInternalServerError, err.Error())
	}
	return apiwrap.SuccessWithMsg("标签更新成功")
}

func (h *LabelHandler) AdminDelete(c *gin.Context) *apiwrap.Response[any] {
	id := c.Param("id")
	err := h.serv.Delete(c, id)
	if err != nil {
		return apiwrap.FailWithMsg(apiwrap.RuquestInternalServerError, err.Error())
	}
	return apiwrap.SuccessWithMsg("标签删除成功")
}

func (h *LabelHandler) GetByID(c *gin.Context) *apiwrap.Response[any] {
	id := c.Param("id")
	label, err := h.serv.Get(c, id)
	if err != nil {
		return apiwrap.FailWithMsg(apiwrap.RuquestInternalServerError, err.Error())
	}
	return apiwrap.SuccessWithDetail[any](label, "标签获取成功")
}

func (h *LabelHandler) GetList(c *gin.Context, page *Page) *apiwrap.Response[any] {
	labels, count, err := h.serv.GetList(c, page.LabelType, page.PageNo, page.PageSize)
	if err != nil {
		return apiwrap.FailWithMsg(apiwrap.RuquestInternalServerError, err.Error())
	}
	labelVOList := h.DomainToVOList(labels)
	pageVO := apiwrap.ToPageVO(page.PageNo, page.PageSize, count, labelVOList)
	return apiwrap.SuccessWithDetail[any](pageVO, "标签列表获取成功")
}

func (h *LabelHandler) DTOToDomain(label *LabelRequest) *domain.Label {
	return &domain.Label{
		ID:        label.ID,
		LabelType: label.LabelType,
		Name:      label.Name,
	}
}

func (h *LabelHandler) DomainToVO(label *domain.Label) *LabelVO {
	return &LabelVO{
		ID:        label.ID,
		LabelType: label.LabelType,
		Name:      label.Name,
	}
}

func (h *LabelHandler) DomainToVOList(labels []*domain.Label) []*LabelVO {
	voList := make([]*LabelVO, 0)
	for _, label := range labels {
		voList = append(voList, h.DomainToVO(label))
	}
	return voList
}
