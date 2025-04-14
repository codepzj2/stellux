package dao

import (
	"github.com/chenmingyong0423/go-mongox/v2"
)

type Post struct {
	mongox.Model
}

type IPostDao interface {
}

var _ IPostDao = (*PostDao)(nil)

func NewPostDao(db *mongox.Database) *PostDao {
	return &PostDao{coll: mongox.NewCollection[Post](db, "post")}
}

type PostDao struct {
	coll *mongox.Collection[Post]
}
