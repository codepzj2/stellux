package api

import (
	"fmt"
	"net/http"
	"server/internal/file/internal/service"
	"server/internal/pkg/wrap"
	"strconv"

	"github.com/gin-gonic/gin"
)

type IFileHandler interface {
	UploadFilesLocal(ctx *gin.Context)
	GetPhotosLocalByPage(ctx *gin.Context)
}

type FileHandler struct {
	serv service.IFileService
}

var _ IFileHandler = (*FileHandler)(nil)

func NewFileHandler(serv service.IFileService) *FileHandler {
	return &FileHandler{serv}
}

func (h *FileHandler) RegisterGinRoutes(router *gin.Engine) {
	router.Static("/images", "./static/images")
	group := router.Group("/picture")
	{
		group.GET("/list", h.GetPhotosLocalByPage)
		group.POST("/local/upload", h.UploadFilesLocal)
	}
}

// UploadFilesLocal 上传图片到本地
func (h *FileHandler) UploadFilesLocal(ctx *gin.Context) {
	form, err := ctx.MultipartForm()
	if err != nil {
		wrap.FailWithMsg(ctx, http.StatusBadRequest, "文件上传失败，请检查文件字段")
		return
	}
	files := form.File["files"]
	if len(files) == 0 {
		wrap.FailWithMsg(ctx, http.StatusBadRequest, "未找到上传的文件")
		return
	}
	pictures, err := h.serv.UploadFilesLocal(ctx, files)
	if err != nil {
		wrap.FailWithMsg(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	count := len(pictures)
	wrap.SuccessWithMsg(ctx, fmt.Sprintf("上传成功，共%d张图片", count))
}

func (h *FileHandler) GetPhotosLocalByPage(ctx *gin.Context) {
	pageStr := ctx.DefaultQuery("page_no", "1")
	sizeStr := ctx.DefaultQuery("size", "10")
	pageNo, err := strconv.Atoi(pageStr)
	if err != nil {
		wrap.FailWithMsg(ctx, http.StatusBadRequest, err.Error())
	}
	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		wrap.FailWithMsg(ctx, http.StatusBadRequest, err.Error())
	}

	photos, totalCount, totalPage, err := h.serv.GetPhotosByPage(ctx, int64(pageNo), int64(size))
	if err != nil {
		wrap.FailWithMsg(ctx, http.StatusInternalServerError, err.Error())
	}

	filePageDTO := wrap.ToPageVO(int64(pageNo), int64(size), totalCount, totalPage, photos)
	wrap.SuccessWithData(ctx, filePageDTO)
}
