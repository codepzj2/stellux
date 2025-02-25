package api

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"server/internal/pkg/http/resp"
	"time"

	"github.com/gin-gonic/gin"
)

type IPictureHandler interface {
	UploadPicturesLocal(ctx *gin.Context)
}

type PictureHandler struct {
}

var _ IPictureHandler = (*PictureHandler)(nil)

func NewPictureHandler() *PictureHandler {
	return &PictureHandler{}
}

func (h *PictureHandler) RegisterGinRoutes(router *gin.Engine) {
	group := router.Group("/picture")
	{
		group.POST("/upload", h.UploadPicturesLocal)
	}
}

func (h *PictureHandler) UploadPicturesLocal(ctx *gin.Context) {
	form, err := ctx.MultipartForm()
	if err != nil {
		resp.FailWithMsg(ctx, http.StatusBadRequest, "文件上传失败，请检查文件字段")
		return
	}
	log.Println("form:", form)
	files := form.File["files"]
	if len(files) == 0 {
		resp.FailWithMsg(ctx, http.StatusBadRequest, "未找到上传的文件")
		return
	}

	var uploadErr error

	for _, file := range files {
		log.Println("上传文件名:", file.Filename)

		newFileName := time.Now().Format("20060102150405") + filepath.Ext(file.Filename)
		filePath := filepath.Join("images", newFileName)

		if err := os.MkdirAll("images", os.ModePerm); err != nil {
			log.Printf("创建目录失败: %v\n", err)
			uploadErr = fmt.Errorf("服务器错误，创建目录失败")
			break
		}

		if err := ctx.SaveUploadedFile(file, filePath); err != nil {
			log.Printf("保存文件失败: %v\n", err)
			uploadErr = fmt.Errorf("保存文件失败")
			break
		}

	}

	if uploadErr != nil {
		resp.FailWithMsg(ctx, http.StatusInternalServerError, uploadErr.Error())
		return
	}

	resp.SuccessWithMsg(ctx, "图片上传成功")
}
