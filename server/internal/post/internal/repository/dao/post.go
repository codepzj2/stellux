package dao

import (
	"context"

	"github.com/chenmingyong0423/go-mongox/v2"
	"github.com/chenmingyong0423/go-mongox/v2/builder/query"
	"github.com/chenmingyong0423/go-mongox/v2/builder/update"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Post struct {
	mongox.Model `bson:",inline"`
	Title        string   `bson:"title,omitempty"`
	Content      string   `bson:"content,omitempty"`
	Description  string   `bson:"description,omitempty"`
	Author       string   `bson:"author,omitempty"`
	Category     string   `bson:"category,omitempty"`
	Tags         []string `bson:"tags,omitempty"`
	IsPublish    bool     `bson:"is_publish,omitempty"`
	IsTop        bool     `bson:"is_top,omitempty"`
	Thumbnail    string   `bson:"thumbnail,omitempty"`
	LikeCount    int      `bson:"like_count,omitempty"`
	ViewCount    int      `bson:"view_count,omitempty"`
	ShareCount   int      `bson:"share_count,omitempty"`
}

type IPostDao interface {
	Create(ctx context.Context, post *Post) error
	Update(ctx context.Context, id string, post *Post) error
	Delete(ctx context.Context, id string) error
	Get(ctx context.Context, id bson.ObjectID) (*Post, error)
	GetList(ctx context.Context, cond bson.D, findOptions *options.FindOptionsBuilder) ([]*Post, int64, error)
}

var _ IPostDao = (*PostDao)(nil)

func NewPostDao(db *mongox.Database) *PostDao {
	return &PostDao{coll: mongox.NewCollection[Post](db, "post")}
}

type PostDao struct {
	coll *mongox.Collection[Post]
}

// Create 创建文章
func (d *PostDao) Create(ctx context.Context, post *Post) error {
	_, err := d.coll.Creator().InsertOne(ctx, post)
	return err
}

// Update 更新文章(排除点赞数、浏览数、分享数)
func (d *PostDao) Update(ctx context.Context, id string, post *Post) error {
	_, err := d.coll.Updater().Filter(query.Id(id)).Updates(update.NewBuilder().Set("title", post.Title).Set("content", post.Content).Set("description", post.Description).Set("author", post.Author).Set("category", post.Category).Set("tags", post.Tags).Set("is_published", post.IsPublish).Set("thumbnail", post.Thumbnail).Build()).UpdateOne(ctx)
	return err
}

// Delete 删除文章
func (d *PostDao) Delete(ctx context.Context, id string) error {
	_, err := d.coll.Deleter().Filter(query.Id(id)).DeleteOne(ctx)
	return err
}

// Get 获取文章
func (d *PostDao) Get(ctx context.Context, id bson.ObjectID) (*Post, error) {
	post, err := d.coll.Finder().Filter(query.Id(id)).FindOne(ctx)
	return post, err
}

// GetList 获取文章列表
func (d *PostDao) GetList(ctx context.Context, cond bson.D, findOptions *options.FindOptionsBuilder) ([]*Post, int64, error) {
	count, err := d.coll.Finder().Filter(cond).Count(ctx)
	if err != nil {
		return nil, 0, err
	}
	posts, err := d.coll.Finder().Filter(cond).Find(ctx, findOptions)
	return posts, count, err
}
