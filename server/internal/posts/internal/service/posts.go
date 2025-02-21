package service

import (
	"context"
	"server/internal/posts/internal/domain"
	"server/internal/posts/internal/repo"
)

type IPostsService interface {
	CreatePosts(ctx context.Context, posts *domain.Posts) error
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
