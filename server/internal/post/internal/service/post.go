package service

import (
	"github.com/codepzj/stellux/server/internal/post/internal/repository"
)


type IPostService interface {
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
