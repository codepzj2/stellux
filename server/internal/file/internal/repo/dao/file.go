package dao

import (
	"context"
	"server/internal/file/internal/domain"

	"github.com/chenmingyong0423/go-mongox/builder/query"
	"github.com/chenmingyong0423/go-mongox/v2"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type IFileDao interface {
	Create(ctx context.Context, files []*domain.File) error
	Delete(ctx context.Context, filesId []bson.ObjectID) error
}

type FileDao struct {
	fileColl *mongox.Collection[domain.File]
}

var _ IFileDao = (*FileDao)(nil)

func NewFileDao(database *mongox.Database) *FileDao {
	return &FileDao{
		fileColl: mongox.NewCollection[domain.File](database, "file"),
	}
}

func (p *FileDao) Create(ctx context.Context, files []*domain.File) error {
	result, err := p.fileColl.Creator().InsertMany(ctx, files)
	if err != nil {
		return err
	}
	if result.InsertedIDs == nil {
		return err
	}
	return nil
}

func (p *FileDao) Delete(ctx context.Context, filesId []bson.ObjectID) error {
	result, err := p.fileColl.Deleter().Filter(query.In("_id", filesId)).DeleteMany(ctx)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.Wrap(err, "删除失败")
	}
	return nil
}
