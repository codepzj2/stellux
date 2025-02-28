package service

import (
	"mime/multipart"
	"os"
	"path/filepath"
	"server/internal/file/internal/domain"
	"server/internal/file/internal/repo"
	"strings"

	"github.com/pkg/errors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type IFileService interface {
	UploadFilesLocal(ctx *gin.Context, files []*multipart.FileHeader) ([]*domain.File, error)
}

type FileService struct {
	repo repo.IFileRepo
}

var _ IFileService = (*FileService)(nil)

func NewFileService(fileRepo repo.IFileRepo) *FileService {
	return &FileService{fileRepo}
}

func (s *FileService) UploadFilesLocal(ctx *gin.Context, files []*multipart.FileHeader) ([]*domain.File, error) {
	var pictures []*domain.File
	os.MkdirAll("images", os.ModePerm)
	for _, file := range files {
		// 生成新的文件名
		newFileName := strings.ReplaceAll(uuid.New().String(), "-", "") + filepath.Ext(file.Filename)
		networkPath := "images/" + newFileName
		filePath := "static/images/" + newFileName
		picture := &domain.File{
			Type: "local",
			Url:  networkPath,
			Dst:  filePath,
		}
		pictures = append(pictures, picture)
		// 保存文件到本地
		if err := ctx.SaveUploadedFile(file, filePath); err != nil {
			// 回滚
			for _, picture := range pictures {
				os.Remove(picture.Dst)
			}
			s.repo.DeleteMany(ctx, pictures)
			return nil, errors.New("保存文件到本地失败")
		}
	}

	// 保存到数据库
	if err := s.repo.CreateMany(ctx, pictures); err != nil {
		// 回滚
		for _, picture := range pictures {
			os.Remove(picture.Dst)
		}
		s.repo.DeleteMany(ctx, pictures)
		return nil, errors.Wrap(err, "保存到数据库失败")
	}
	return pictures, nil
}
