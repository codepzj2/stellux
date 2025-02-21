package service

import (
	"context"
	"server/internal/posts/internal/domain"
	"server/internal/posts/internal/repo"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type IPostsService interface {
	CreatePosts(ctx context.Context, posts *domain.Posts) error
	FindPostById(ctx context.Context, id bson.ObjectID) (*domain.Posts, error)
	FindAllPosts(ctx context.Context) ([]*domain.Posts, error)
}
type PostsService struct {
	repo repo.IPostsRepo
}

func NewPostsService(repo repo.IPostsRepo) *PostsService {
	return &PostsService{repo}
}

func (p *PostsService) CreatePosts(ctx context.Context, posts *domain.Posts) error {
	return p.repo.CreatePost(ctx, posts)
}

func (p *PostsService) FindPostById(ctx context.Context, id bson.ObjectID) (*domain.Posts, error) {
	return p.repo.FindPostById(ctx, id)
}

func (p *PostsService) FindAllPosts(ctx context.Context) ([]*domain.Posts, error) {
	return p.repo.FindAllPosts(ctx)
}
