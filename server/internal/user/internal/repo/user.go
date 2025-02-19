package repo

import (
	"context"
	"server/internal/user/internal/domain"
	"server/internal/user/internal/repo/dao"
	"slices"

	"go.mongodb.org/mongo-driver/v2/bson"

	"github.com/pkg/errors"
)

type IUserRepo interface {
	CreateUser(ctx context.Context, user *domain.User) error
	FindIsExist(ctx context.Context, user *domain.User) (*domain.User, bool)
	FindAllUsersByPermission(ctx context.Context, userId bson.ObjectID) ([]*domain.User, error)
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

func (u *UserRepo) FindAllUsersByPermission(ctx context.Context, userId bson.ObjectID) ([]*domain.User, error) {
	user, err := u.dao.FindById(ctx, userId)
	// 判断用户是否存在
	if err != nil || user == nil {
		return nil, errors.Wrap(err, "获取用户信息失败")
	}
	// 判断用户权限是否在[0,1,2]之间
	roleId := user.RoleId
	if !slices.Contains([]int{0, 1, 2}, roleId) {
		return nil, errors.Wrap(err, "用户权限超出范围")
	}
	return u.dao.FindAllByRoleID(ctx, roleId)
}
