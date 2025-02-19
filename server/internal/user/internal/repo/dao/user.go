package dao

import (
	"context"
	"server/internal/user/internal/domain"

	"go.mongodb.org/mongo-driver/v2/bson"

	"github.com/chenmingyong0423/go-mongox/v2"
	"github.com/chenmingyong0423/go-mongox/v2/builder/query"
	"github.com/pkg/errors"
)

type IUserDao interface {
	CreateOne(ctx context.Context, user *domain.User) error
	FindByName(ctx context.Context, username string) (*domain.User, error)
	FindById(ctx context.Context, userId bson.ObjectID) (*domain.User, error)
	FindIsExist(ctx context.Context, user *domain.User) (*domain.User, error)
	FindAll(ctx context.Context) ([]*domain.User, error)
	FindAllByRoleID(ctx context.Context, roleId int) ([]*domain.User, error)
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

func (u *UserDao) FindById(ctx context.Context, userId bson.ObjectID) (*domain.User, error) {
	return u.userColl.Finder().Filter(query.Id(userId)).FindOne(ctx)
}

func (u *UserDao) FindIsExist(ctx context.Context, user *domain.User) (*domain.User, error) {
	return u.userColl.Finder().Filter(query.NewBuilder().Eq("username", user.Username).Eq("password", user.Password).Build()).FindOne(ctx)
}

func (u *UserDao) FindAll(ctx context.Context) ([]*domain.User, error) {
	return u.userColl.Finder().Find(ctx)
}

func (u *UserDao) FindAllByRoleID(ctx context.Context, roleId int) ([]*domain.User, error) {
	return u.userColl.Finder().Filter(query.NewBuilder().Gte("role_id", roleId).Build()).Find(ctx)
}
