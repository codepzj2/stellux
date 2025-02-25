package repo

import (
	"context"
	"server/internal/file/internal/domain"
	"server/internal/file/internal/repo/dao"

	"github.com/samber/lo"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type IFileRepo interface {
	CreateMany(ctx context.Context, files []*domain.File) error
	DeleteMany(ctx context.Context, files []*domain.File) error
}

type FileRepo struct {
	fileDao dao.IFileDao
}

var _ IFileRepo = (*FileRepo)(nil)

func NewFileRepo(fileDao dao.IFileDao) *FileRepo {
	return &FileRepo{fileDao}
}

func (p *FileRepo) CreateMany(ctx context.Context, files []*domain.File) error {
	return p.fileDao.Create(ctx, files)
}

func (p *FileRepo) DeleteMany(ctx context.Context, files []*domain.File) error {
	fileIds := lo.Map(files, func(file *domain.File, _ int) bson.ObjectID {
		return file.ID
	})
	return p.fileDao.Delete(ctx, fileIds)
}
