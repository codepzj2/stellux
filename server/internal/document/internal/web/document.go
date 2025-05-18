package web

import (
	"github.com/codepzj/stellux/server/internal/document/internal/domain"
	"github.com/codepzj/stellux/server/internal/document/internal/service"
	"github.com/codepzj/stellux/server/internal/pkg/apiwrap"
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
)

func NewDocumentHandler(serv service.IDocumentService) *DocumentHandler {
	return &DocumentHandler{
		serv: serv,
	}
}

type DocumentHandler struct {
	serv service.IDocumentService
}

func (h *DocumentHandler) RegisterGinRoutes(engine *gin.Engine) {
	documentGroup := engine.Group("/document")
	{
		documentGroup.GET("/list", apiwrap.Wrap(h.FindAllByDocumentID))
	}
	adminGroup := engine.Group("/admin-api/document")
	{
		adminGroup.POST("/create", apiwrap.WrapWithBody(h.CreateDocument))
	}
}

func (h *DocumentHandler) CreateDocument(c *gin.Context, documentReq DocumentRequest) *apiwrap.Response[any] {
	err := h.serv.Create(c.Request.Context(), h.DocumentRequestToDomain(documentReq))
	if err != nil {
		return apiwrap.FailWithMsg(apiwrap.RuquestInternalServerError, err.Error())
	}
	return apiwrap.Success()
}

func (h *DocumentHandler) FindAllByDocumentID(c *gin.Context) *apiwrap.Response[any] {
	documentID := c.Query("document_id")
	documentList, err := h.serv.FindAllByDocumentID(c.Request.Context(), apiwrap.ConvertBsonID(documentID))
	if err != nil {
		return apiwrap.FailWithMsg(apiwrap.RuquestInternalServerError, err.Error())
	}
	return apiwrap.SuccessWithDetail[any](h.DocumentDomainListToVOList(documentList), "获取文档列表成功")
}

func (h *DocumentHandler) DocumentRequestToDomain(req DocumentRequest) *domain.Document {
	return &domain.Document{
		Title:      req.Title,
		Content:    req.Content,
		ParentID:   apiwrap.ConvertBsonID(req.ParentID),
		DocumentID: apiwrap.ConvertBsonID(req.DocumentID),
	}
}

func (h *DocumentHandler) DocumentDomainToVO(doc *domain.Document) *DocumentVO {
	return &DocumentVO{
		ID:         doc.ID.Hex(),
		CreatedAt:  doc.CreatedAt.String(),
		UpdatedAt:  doc.UpdatedAt.String(),
		Title:      doc.Title,
		Content:    doc.Content,
		ParentID:   apiwrap.BsonID(doc.ParentID.Hex()),
		DocumentID: apiwrap.BsonID(doc.DocumentID.Hex()),
	}
}

func (h *DocumentHandler) DocumentDomainListToVOList(docList []*domain.Document) []*DocumentVO {
	return lo.Map(docList, func(doc *domain.Document, _ int) *DocumentVO {
		return h.DocumentDomainToVO(doc)
	})
}
