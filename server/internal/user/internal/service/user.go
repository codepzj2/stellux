package service

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/bson"
	"server/internal/user/internal/domain"
	"server/internal/user/internal/repo"
)

type IUserService interface {
	CreateUser(ctx context.Context, user *domain.User) error
	FindIsExist(ctx context.Context, user *domain.User) (*domain.User, bool)
	FindAllUsers(ctx context.Context, userId string) ([]*domain.User, error)
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

// FindAllUsers 根据权限返回用户列表
func (s *UserService) FindAllUsers(ctx context.Context, userId string) ([]*domain.User, error) {
	uid, _ := bson.ObjectIDFromHex(userId)
	return s.repo.FindAllUsersByPermission(ctx, uid)
}
