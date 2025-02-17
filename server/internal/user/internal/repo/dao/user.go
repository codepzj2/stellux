package dao

import (
	"context"
	"github.com/chenmingyong0423/go-mongox/v2"
	"github.com/chenmingyong0423/go-mongox/v2/builder/query"
	"github.com/pkg/errors"
	"server/internal/user/internal/domain"
)

type IUserDao interface {
	CreateOne(ctx context.Context, user *domain.User) error
	FindByName(ctx context.Context, name string) (*domain.User, error)
}

var _ IUserDao = (*UserDao)(nil)

type UserDao struct {
	userColl *mongox.Collection[domain.User]
}

func NewUserDao(database *mongox.Database) *UserDao {
	return &UserDao{userColl: mongox.NewCollection[domain.User](database, "user")}
}

func (u *UserDao) CreateOne(ctx context.Context, user *domain.User) error {
	_, err := u.userColl.Creator().InsertOne(ctx, user)
	if err != nil {
		return errors.Wrapf(err, "添加user失败,%+v", user)
	}
	return err
}

func (u *UserDao) FindByName(ctx context.Context, username string) (*domain.User, error) {
	return u.userColl.Finder().Filter(query.Eq("username", username)).FindOne(ctx)
}
