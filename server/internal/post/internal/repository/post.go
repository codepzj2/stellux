package repository

import (
	"github.com/codepzj/stellux/server/internal/post/internal/repository/dao"
)

type IPostRepository interface {
}

var _ IPostRepository = (*PostRepository)(nil)

func NewPostRepository(dao dao.IPostDao) *PostRepository {
	return &PostRepository{dao: dao}
}

type PostRepository struct {
	dao dao.IPostDao
}
