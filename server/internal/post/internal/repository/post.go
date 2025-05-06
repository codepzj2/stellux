package repository

import (
	"context"
	"fmt"

	"github.com/chenmingyong0423/go-mongox/v2/bsonx"
	"github.com/chenmingyong0423/go-mongox/v2/builder/aggregation"
	"github.com/chenmingyong0423/go-mongox/v2/builder/query"
	"github.com/codepzj/stellux/server/internal/pkg/apiwrap"
	"github.com/codepzj/stellux/server/internal/post/internal/domain"
	"github.com/codepzj/stellux/server/internal/post/internal/repository/dao"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type IPostRepository interface {
	Create(ctx context.Context, post *domain.Post) error
	Update(ctx context.Context, post *domain.Post) error
	Delete(ctx context.Context, id bson.ObjectID) error
	Get(ctx context.Context, id bson.ObjectID) (*domain.PostDetail, error)
	GetList(ctx context.Context, page *apiwrap.Page) ([]*domain.PostDetail, int64, error)
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
	return r.dao.Update(ctx, post.ID, r.PostDomainToPostDO(post))
}

// Delete 删除文章
func (r *PostRepository) Delete(ctx context.Context, id bson.ObjectID) error {
	return r.dao.Delete(ctx, id)
}

// Get 获取文章
func (r *PostRepository) Get(ctx context.Context, id bson.ObjectID) (*domain.PostDetail, error) {
	postCategoryTags, err := r.dao.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return r.PostCategoryTagsDOToPostDetail(postCategoryTags), nil
}

// GetList 获取文章列表
func (r *PostRepository) GetList(ctx context.Context, page *apiwrap.Page) ([]*domain.PostDetail, int64, error) {
	cond := query.NewBuilder().Or(query.Regex("title", page.Keyword), query.Regex("content", page.Keyword), query.Regex("description", page.Keyword)).Build()
	skip := (page.PageNo - 1) * page.PageSize
	limit := page.PageSize
	sortBuilder := bsonx.NewD().Add("created_at", -1)
	if page.Field != "" {
		sortBuilder.Add(page.Field, r.OrderConvertToInt(page.Order))
	}
	fmt.Println(sortBuilder.Build())
	pagePipeline := aggregation.NewStageBuilder().Lookup("label", "category", &aggregation.LookUpOptions{
		LocalField:   "category_id",
		ForeignField: "_id",
	}).Unwind("$category", nil).Lookup("label", "tags", &aggregation.LookUpOptions{
		LocalField:   "tags_id",
		ForeignField: "_id",
	}).Skip(skip).Limit(limit).Sort(sortBuilder.Build()).Match(cond).Build()

	posts, count, err := r.dao.GetList(ctx, pagePipeline, cond)
	if err != nil {
		return nil, 0, err
	}
	return r.PostCategoryTagsDOToPostDetailList(posts), count, nil
}

// PostDomain2PostDO 将domain.Post转换为dao.Post
func (r *PostRepository) PostDomainToPostDO(post *domain.Post) *dao.Post {
	return &dao.Post{
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
	}
}

func (r *PostRepository) PostCategoryTagsDOToPostDetailList(postCategoryTags []*dao.PostCategoryTags) []*domain.PostDetail {
	postDetailList := make([]*domain.PostDetail, len(postCategoryTags))
	for i, postCategoryTag := range postCategoryTags {
		postDetailList[i] = r.PostCategoryTagsDOToPostDetail(postCategoryTag)
	}
	return postDetailList
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
