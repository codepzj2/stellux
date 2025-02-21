package repo

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/bson"
	"server/internal/posts/internal/domain"
	"server/internal/posts/internal/repo/dao"
)

type IPostsRepo interface {
	CreatePost(ctx context.Context, posts *domain.Posts) error
	FindAll(ctx context.Context) ([]*domain.Posts, error)
	DeletePostById(ctx context.Context, Id bson.ObjectID) error
}

type PostsRepo struct {
	dao dao.IPostsDao
}

var _ IPostsRepo = (*PostsRepo)(nil)

func NewPostsRepo(dao dao.IPostsDao) *PostsRepo {
	return &PostsRepo{dao}
}

func (p *PostsRepo) CreatePost(ctx context.Context, posts *domain.Posts) error {
	return p.dao.Create(ctx, posts)
}

func (p *PostsRepo) FindAll(ctx context.Context) ([]*domain.Posts, error) {
	return p.dao.FindAll(ctx)
}

func (p *PostsRepo) DeletePostById(ctx context.Context, Id bson.ObjectID) error {
	return p.dao.DeleteById(ctx, Id)
}
