package repo

import (
	"context"
	"server/internal/user/internal/domain"
	"server/internal/user/internal/repo/dao"

	"github.com/pkg/errors"
)

type IUserRepo interface {
	CreateUser(ctx context.Context, user *domain.User) error
	FindIsExist(ctx context.Context, user *domain.User) (*domain.User, bool)
}

var _ IUserRepo = (*UserRepo)(nil)

type UserRepo struct {
	dao dao.IUserDao
}

func NewUserRepo(dao dao.IUserDao) *UserRepo {
	return &UserRepo{dao}
}

func (u *UserRepo) CreateUser(ctx context.Context, user *domain.User) error {
	result, err := u.dao.FindByName(ctx, user.Username)
	if err != nil {
		// 如果集合中文档为空
		if err.Error() == "mongo: no documents in result" {
			return u.dao.CreateOne(ctx, user)
		}
		return err
	}
	if result != nil {
		return errors.New("用户名已存在")
	}
	return u.dao.CreateOne(ctx, user)
}

func (u *UserRepo) FindIsExist(ctx context.Context, user *domain.User) (*domain.User, bool) {
	result, err := u.dao.FindIsExist(ctx, user)
	if err != nil {
		return nil, false
	}
	return result, result != nil
}
