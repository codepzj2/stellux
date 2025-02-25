package api

import (
	"fmt"
	"net/http"
	"server/internal/file/internal/service"
	"server/internal/pkg/http/resp"

	"github.com/gin-gonic/gin"
)

type IFileHandler interface {
	UploadFilesLocal(ctx *gin.Context)
}

type FileHandler struct {
	serv service.IFileService
}

var _ IFileHandler = (*FileHandler)(nil)

func NewFileHandler(serv service.IFileService) *FileHandler {
	return &FileHandler{serv}
}

func (h *FileHandler) RegisterGinRoutes(router *gin.Engine) {
	group := router.Group("/picture")
	{
		group.POST("/local/upload", h.UploadFilesLocal)
	}
}

// 上传图片到本地
func (h *FileHandler) UploadFilesLocal(ctx *gin.Context) {
	form, err := ctx.MultipartForm()
	if err != nil {
		resp.FailWithMsg(ctx, http.StatusBadRequest, "文件上传失败，请检查文件字段")
		return
	}
	files := form.File["files"]
	if len(files) == 0 {
		resp.FailWithMsg(ctx, http.StatusBadRequest, "未找到上传的文件")
		return
	}
	pictures, err := h.serv.UploadFilesLocal(ctx, files)
	if err != nil {
		resp.FailWithMsg(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	count := len(pictures)
	resp.SuccessWithMsg(ctx, fmt.Sprintf("上传成功，共%d张图片", count))
}
