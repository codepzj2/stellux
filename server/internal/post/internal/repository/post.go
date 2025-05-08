package repository

import (
	"context"

	"github.com/chenmingyong0423/go-mongox/v2"
	"github.com/chenmingyong0423/go-mongox/v2/bsonx"
	"github.com/chenmingyong0423/go-mongox/v2/builder/aggregation"
	"github.com/chenmingyong0423/go-mongox/v2/builder/query"
	"github.com/codepzj/stellux/server/internal/pkg/apiwrap"
	"github.com/codepzj/stellux/server/internal/post/internal/domain"
	"github.com/codepzj/stellux/server/internal/post/internal/repository/dao"
	"github.com/samber/lo"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type IPostRepository interface {
	Create(ctx context.Context, post *domain.Post) error
	Update(ctx context.Context, post *domain.Post) error
	SoftDelete(ctx context.Context, id bson.ObjectID) error
	Delete(ctx context.Context, id bson.ObjectID) error
	GetByID(ctx context.Context, id bson.ObjectID) (*domain.Post, error)
	GetDetailByID(ctx context.Context, id bson.ObjectID) (*domain.PostDetail, error)
	GetDetailList(ctx context.Context, page *apiwrap.Page) ([]*domain.PostDetail, int64, error)
}

var _ IPostRepository = (*PostRepository)(nil)

func NewPostRepository(dao dao.IPostDao) *PostRepository {
	return &PostRepository{dao: dao}
}

type PostRepository struct {
	dao dao.IPostDao
}

func (r *PostRepository) Create(ctx context.Context, post *domain.Post) error {
	return r.dao.Create(ctx, r.PostDomainToPostDO(post))
}

func (r *PostRepository) Update(ctx context.Context, post *domain.Post) error {
	return r.dao.Update(ctx, post.ID, r.PostDomainToUpdatePostDO(post))
}

func (r *PostRepository) SoftDelete(ctx context.Context, id bson.ObjectID) error {
	return r.dao.SoftDelete(ctx, id)
}

// Delete 删除文章
func (r *PostRepository) Delete(ctx context.Context, id bson.ObjectID) error {
	return r.dao.Delete(ctx, id)
}

// GetByID 获取文章
func (r *PostRepository) GetByID(ctx context.Context, id bson.ObjectID) (*domain.Post, error) {
	post, err := r.dao.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return r.PostDOToPostDomain(post), nil
}

// GetDetailByID 获取文章
func (r *PostRepository) GetDetailByID(ctx context.Context, id bson.ObjectID) (*domain.PostDetail, error) {
	postCategoryTags, err := r.dao.GetDetailByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return r.PostCategoryTagsDOToPostDetail(postCategoryTags), nil
}

// GetDetailList 获取文章列表
func (r *PostRepository) GetDetailList(ctx context.Context, page *apiwrap.Page) ([]*domain.PostDetail, int64, error) {
	cond := query.NewBuilder().Or(query.Regex("title", page.Keyword), query.Regex("description", page.Keyword)).Build()
	skip := (page.PageNo - 1) * page.PageSize
	limit := page.PageSize
	sortBuilder := bsonx.NewD().Add("created_at", -1)
	if page.Field != "" {
		sortBuilder.Add(page.Field, r.OrderConvertToInt(page.Order))
	}
	pagePipeline := aggregation.NewStageBuilder().Lookup("label", "category", &aggregation.LookUpOptions{
		LocalField:   "category_id",
		ForeignField: "_id",
	}).Unwind("$category", nil).Lookup("label", "tags", &aggregation.LookUpOptions{
		LocalField:   "tags_id",
		ForeignField: "_id",
	}).Skip(skip).Limit(limit).Sort(sortBuilder.Build()).Match(cond).Build()

	posts, count, err := r.dao.GetDetailList(ctx, pagePipeline, cond)
	if err != nil {
		return nil, 0, err
	}
	return r.PostCategoryTagsDOToPostDetailList(posts), count, nil
}

// PostDomain2PostDO 将domain.Post转换为dao.Post
func (r *PostRepository) PostDomainToPostDO(post *domain.Post) *dao.Post {
	return &dao.Post{
		Model:       mongox.Model{CreatedAt: post.CreatedAt},
		Title:       post.Title,
		Content:     post.Content,
		Description: post.Description,
		Author:      post.Author,
		CategoryID:  post.CategoryID,
		TagsID:      post.TagsID,
		IsPublish:   post.IsPublish,
		IsTop:       post.IsTop,
		Thumbnail:   post.Thumbnail,
	}
}

func (r *PostRepository) PostDomainToUpdatePostDO(post *domain.Post) *dao.UpdatePost {
	return &dao.UpdatePost{
		CreatedAt:   post.CreatedAt,
		Title:       post.Title,
		Content:     post.Content,
		Description: post.Description,
		Author:      post.Author,
		CategoryID:  post.CategoryID,
		TagsID:      post.TagsID,
		IsPublish:   post.IsPublish,
		IsTop:       post.IsTop,
		Thumbnail:   post.Thumbnail,
	}
}

// PostDOToPostDomain 将dao.Post转换为domain.Post
func (r *PostRepository) PostDOToPostDomain(post *dao.Post) *domain.Post {
	return &domain.Post{
		ID:          post.ID,
		CreatedAt:   post.CreatedAt,
		Title:       post.Title,
		Content:     post.Content,
		Description: post.Description,
		Author:      post.Author,
		CategoryID:  post.CategoryID,
		TagsID:      post.TagsID,
		IsPublish:   post.IsPublish,
		IsTop:       post.IsTop,
		Thumbnail:   post.Thumbnail,
	}
}

// PostCategoryTagsDOToPostDetail 将dao.PostCategoryTags转换为domain.PostDetail
func (r *PostRepository) PostCategoryTagsDOToPostDetail(postCategoryTags *dao.PostCategoryTags) *domain.PostDetail {
	return &domain.PostDetail{
		ID:          postCategoryTags.ID,
		CreatedAt:   postCategoryTags.CreatedAt,
		UpdatedAt:   postCategoryTags.UpdatedAt,
		Title:       postCategoryTags.Title,
		Content:     postCategoryTags.Content,
		Description: postCategoryTags.Description,
		Author:      postCategoryTags.Author,
		Category:    postCategoryTags.Category,
		Tags:        postCategoryTags.Tags,
		Thumbnail:   postCategoryTags.Thumbnail,
		IsPublish:   postCategoryTags.IsPublish,
		IsTop:       postCategoryTags.IsTop,
	}
}

func (r *PostRepository) PostCategoryTagsDOToPostDetailList(postCategoryTags []*dao.PostCategoryTags) []*domain.PostDetail {
	return lo.Map(postCategoryTags, func(postCategoryTag *dao.PostCategoryTags, _ int) *domain.PostDetail {
		return r.PostCategoryTagsDOToPostDetail(postCategoryTag)
	})
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
