package dao

import (
	"context"
	"server/internal/file/internal/domain"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type IFileDao interface {
	Create(ctx context.Context, files []*domain.File) error
	Delete(ctx context.Context, filesId []bson.ObjectID) error
}

type FileDao struct {
	fileColl *mongo.Collection
}

var _ IFileDao = (*FileDao)(nil)

func NewFileDao(database *mongo.Database) *FileDao {
	return &FileDao{
		fileColl: database.Collection("file"),
	}
}

func (p *FileDao) Create(ctx context.Context, files []*domain.File) error {
	result, err := p.fileColl.InsertMany(ctx, files)
	if err != nil {
		return err
	}
	if result.InsertedIDs == nil {
		return err
	}
	return nil
}

func (p *FileDao) Delete(ctx context.Context, filesId []bson.ObjectID) error {
	result, err := p.fileColl.DeleteMany(ctx, bson.M{"_id": bson.M{"$in": filesId}})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.Wrap(err, "删除失败")
	}
	return nil
}
