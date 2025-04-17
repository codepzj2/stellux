package service

import (
	"context"
	"errors"

	"github.com/codepzj/stellux/server/internal/pkg/utils"
	"github.com/codepzj/stellux/server/internal/user/internal/domain"
	"github.com/codepzj/stellux/server/internal/user/internal/repository"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type IUserService interface {
	CheckUserExist(ctx context.Context, user *domain.User) (bool, string)
	AdminCreate(ctx context.Context, user *domain.User) error
	AdminUpdate(ctx context.Context, user *domain.User) error
	AdminDelete(ctx context.Context, id string) error
	GetUserList(ctx context.Context, page *domain.Page) ([]*domain.User, int64, error)
}

var _ IUserService = (*UserService)(nil)

func NewUserService(repo repository.IUserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

type UserService struct {
	repo repository.IUserRepository
}

func (s *UserService) CheckUserExist(ctx context.Context, user *domain.User) (bool, string) {
	u, err := s.repo.GetByUsername(ctx, user.Username)
	if err != nil && err != mongo.ErrNoDocuments {
		return false, ""
	}
	if u == nil {
		return false, ""
	}
	return utils.CompareHashAndPassword(u.Password, user.Password), u.ID
}

func (s *UserService) AdminCreate(ctx context.Context, user *domain.User) error {
	u, err := s.repo.GetByUsername(ctx, user.Username)
	if err != nil && err != mongo.ErrNoDocuments {
		return err
	}
	if u != nil {
		return errors.New("用户已存在")
	}
	// 使用bcrypt生成hash密码
	user.Password, err = utils.GenerateHashPassword(user.Password)
	if err != nil {
		return err
	}
	return s.repo.Create(ctx, user)
}

func (s *UserService) AdminUpdate(ctx context.Context, user *domain.User) error {
	return s.repo.Update(ctx, user)
}

func (s *UserService) AdminDelete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

func (s *UserService) GetUserList(ctx context.Context, page *domain.Page) ([]*domain.User, int64, error) {
	return s.repo.FindByPage(ctx, page)
}
