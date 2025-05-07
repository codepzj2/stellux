package service

import (
	"context"

	"github.com/codepzj/stellux/server/internal/pkg/apiwrap"
	"github.com/codepzj/stellux/server/internal/post/internal/domain"
	"github.com/codepzj/stellux/server/internal/post/internal/repository"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type IPostService interface {
	AdminCreatePost(ctx context.Context, post *domain.Post) error
	AdminUpdatePost(ctx context.Context, post *domain.Post) error
	AdminDeletePost(ctx context.Context, id bson.ObjectID) error
	GetPostById(ctx context.Context, id bson.ObjectID) (*domain.Post, error)
	GetPostDetailById(ctx context.Context, id bson.ObjectID) (*domain.PostDetail, error)
	GetPostDetailList(ctx context.Context, page *apiwrap.Page) ([]*domain.PostDetail, int64, error)
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

func (s *PostService) AdminDeletePost(ctx context.Context, id bson.ObjectID) error {
	return s.repo.Delete(ctx, id)
}

func (s *PostService) GetPostById(ctx context.Context, id bson.ObjectID) (*domain.Post, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *PostService) GetPostDetailById(ctx context.Context, id bson.ObjectID) (*domain.PostDetail, error) {
	return s.repo.GetDetailByID(ctx, id)
}

func (s *PostService) GetPostDetailList(ctx context.Context, page *apiwrap.Page) ([]*domain.PostDetail, int64, error) {
	return s.repo.GetDetailList(ctx, page)
}
