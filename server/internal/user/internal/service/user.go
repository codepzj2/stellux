package service

import (
	"context"
	"server/internal/user/internal/domain"
	"server/internal/user/internal/repo"
)

type IUserService interface {
	CreateUser(ctx context.Context, user *domain.User) error
	FindIsExist(ctx context.Context, user *domain.User) (*domain.User, bool)
}

var _ IUserService = (*UserService)(nil)

type UserService struct {
	repo repo.IUserRepo
}

func NewUserService(repo repo.IUserRepo) *UserService {
	return &UserService{repo}
}

func (s *UserService) CreateUser(ctx context.Context, user *domain.User) error {
	return s.repo.CreateUser(ctx, user)
}

func (s *UserService) FindIsExist(ctx context.Context, user *domain.User) (*domain.User, bool) {
	return s.repo.FindIsExist(ctx, user)
}
