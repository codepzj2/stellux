package dao

import (
	"context"
	"github.com/chenmingyong0423/go-mongox/v2"
	"github.com/chenmingyong0423/go-mongox/v2/builder/query"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/v2/bson"
	"server/internal/posts/internal/domain"
)

type IPostsDao interface {
	Create(ctx context.Context, posts *domain.Posts) error
	FindAll(ctx context.Context) ([]*domain.Posts, error)
	DeleteById(ctx context.Context, Id bson.ObjectID) error
}

type PostsDao struct {
	postColl *mongox.Collection[domain.Posts]
}

var _ IPostsDao = (*PostsDao)(nil)

func NewPostsDao(database *mongox.Database) *PostsDao {
	return &PostsDao{postColl: mongox.NewCollection[domain.Posts](database, "posts")}
}

func (p *PostsDao) Create(ctx context.Context, posts *domain.Posts) error {
	_, err := p.postColl.Creator().InsertOne(ctx, posts)
	if err != nil {
		return errors.Wrap(err, "添加文章失败")
	}
	return nil
}

func (p *PostsDao) FindAll(ctx context.Context) ([]*domain.Posts, error) {
	return p.postColl.Finder().Find(ctx)
}

func (p *PostsDao) DeleteById(ctx context.Context, Id bson.ObjectID) error {
	result, err := p.postColl.Deleter().Filter(query.Id(Id)).DeleteOne(ctx)
	if err != nil {
		return errors.Wrap(err, "删除失败")
	}
	if result.DeletedCount == 0 {
		return errors.Wrap(err, "删除条数为0，删除失败")
	}
	return nil
}
