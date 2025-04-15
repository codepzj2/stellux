package service

import (
	"context"

	"github.com/codepzj/stellux/server/internal/post/internal/domain"
	"github.com/codepzj/stellux/server/internal/post/internal/repository"
)

type IPostService interface {
	AdminCreatePost(ctx context.Context, post *domain.Post) error
	AdminUpdatePost(ctx context.Context, post *domain.Post) error
	AdminDeletePost(ctx context.Context, id string) error
	AdminGetPost(ctx context.Context, id string) (*domain.Post, error)
	AdminGetPostList(ctx context.Context, page *domain.Page) ([]*domain.Post, int64, error)
}

var _ IPostService = (*PostService)(nil)

func NewPostService(repo repository.IPostRepository) *PostService {
	return &PostService{
		repo: repo,
	}
}

type PostService struct {
	repo repository.IPostRepository
}

func (s *PostService) AdminCreatePost(ctx context.Context, post *domain.Post) error {
	return s.repo.Create(ctx, post)
}

func (s *PostService) AdminUpdatePost(ctx context.Context, post *domain.Post) error {
	return s.repo.Update(ctx, post)
}

func (s *PostService) AdminDeletePost(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

func (s *PostService) AdminGetPost(ctx context.Context, id string) (*domain.Post, error) {
	return s.repo.Get(ctx, id)
}

func (s *PostService) AdminGetPostList(ctx context.Context, page *domain.Page) ([]*domain.Post, int64, error) {
	return s.repo.GetList(ctx, page)
}
