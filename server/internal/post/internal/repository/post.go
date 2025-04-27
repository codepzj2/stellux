package repository

import (
	"context"

	"github.com/chenmingyong0423/go-mongox/v2/builder/query"
	"github.com/codepzj/stellux/server/internal/post/internal/domain"
	"github.com/codepzj/stellux/server/internal/post/internal/repository/dao"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type IPostRepository interface {
	Create(ctx context.Context, post *domain.Post) error
	Update(ctx context.Context, post *domain.Post) error
	Delete(ctx context.Context, id string) error
	Get(ctx context.Context, id string) (*domain.Post, error)
	GetList(ctx context.Context, page *domain.Page) ([]*domain.Post, int64, error)
}

var _ IPostRepository = (*PostRepository)(nil)

func NewPostRepository(dao dao.IPostDao) *PostRepository {
	return &PostRepository{dao: dao}
}

type PostRepository struct {
	dao dao.IPostDao
}

func (r *PostRepository) Create(ctx context.Context, post *domain.Post) error {
	return r.dao.Create(ctx, r.DomainToDao(post))
}

// Update 更新文章(排除点赞数、浏览数、分享数)
func (r *PostRepository) Update(ctx context.Context, post *domain.Post) error {
	return r.dao.Update(ctx, post.ID, r.DomainToDao(post))
}

// Delete 删除文章
func (r *PostRepository) Delete(ctx context.Context, id string) error {
	return r.dao.Delete(ctx, id)
}

// Get 获取文章
func (r *PostRepository) Get(ctx context.Context, id string) (*domain.Post, error) {
	bid, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	post, err := r.dao.Get(ctx, bid)
	if err != nil {
		return nil, err
	}
	return r.DaoToDomain(post), nil
}

// GetList 获取文章列表
func (r *PostRepository) GetList(ctx context.Context, page *domain.Page) ([]*domain.Post, int64, error) {
	cond := query.NewBuilder()
	findOptions := options.Find().SetSkip((page.PageNo - 1) * page.PageSize).SetLimit(page.PageSize)
	sortField := "createdAt"
	sortOrder := -1
	// 标题或描述包含关键词
	if page.Keyword != "" {
		cond.RegexOptions("title", page.Keyword, "i").RegexOptions("description", page.Keyword, "i")
	}

	if page.Field != "" && page.Order != "" {
		sortField = page.Field
		sortOrder = r.OrderConvertToInt(page.Order)
	}
	findOptions.SetSort(bson.M{sortField: sortOrder})

	posts, count, err := r.dao.GetList(ctx, cond.Build(), findOptions)
	if err != nil {
		return nil, 0, err
	}
	return r.DaoToDomainList(posts), count, nil
}

// DomainToDao 将domain.Post转换为dao.Post(排除点赞数、浏览数、分享数)
func (r *PostRepository) DomainToDao(post *domain.Post) *dao.Post {
	return &dao.Post{
		Title:       post.Title,
		Content:     post.Content,
		Description: post.Description,
		Author:      post.Author,
		Category:    post.Category,
		Tags:        post.Tags,
		IsPublish:   post.IsPublish,
		IsTop:       post.IsTop,
		Thumbnail:   post.Thumbnail,
	}
}

// DaoToDomain 将dao.Post转换为domain.Post
func (r *PostRepository) DaoToDomain(post *dao.Post) *domain.Post {
	return &domain.Post{
		ID:          post.ID.Hex(),
		CreatedAt:   post.CreatedAt,
		UpdatedAt:   post.UpdatedAt,
		Title:       post.Title,
		Content:     post.Content,
		Description: post.Description,
		Author:      post.Author,
		Category:    post.Category,
		Tags:        post.Tags,
		IsPublish:   post.IsPublish,
		IsTop:       post.IsTop,
		Thumbnail:   post.Thumbnail,
		LikeCount:   post.LikeCount,
		ViewCount:   post.ViewCount,
		ShareCount:  post.ShareCount,
	}
}

func (r *PostRepository) DaoToDomainList(posts []*dao.Post) []*domain.Post {
	domains := make([]*domain.Post, 0, len(posts))
	for _, post := range posts {
		domains = append(domains, r.DaoToDomain(post))
	}
	return domains
}

func (r *PostRepository) OrderConvertToInt(order string) int {
	switch order {
	case "DESC":
		return -1
	case "ASC":
		return 1
	default:
		return -1
	}
}
