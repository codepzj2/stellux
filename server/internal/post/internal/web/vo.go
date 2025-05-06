package web

import (
	"time"

	"github.com/codepzj/stellux/server/internal/label"
	"github.com/codepzj/stellux/server/internal/pkg/apiwrap"
	"github.com/codepzj/stellux/server/internal/post/internal/domain"
	"github.com/samber/lo"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type PostVO struct {
	ID          string    `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	Description string    `json:"description"`
	Author      string    `json:"author"`
	Category    string    `json:"category"`
	Tags        []string  `json:"tags"`
	IsTop       bool      `json:"is_top"`
	Thumbnail   string    `json:"thumbnail"`
}

func GetCategoryNameFromLabel(label label.Domain) string {
	return label.Name
}

func GetTagNamesFromLabels(labels []label.Domain) []string {
	return lo.Map(labels, func(label label.Domain, _ int) string {
		return label.Name
	})
}

func (h *PostHandler) PostDTOToDomain(postReq PostRequest) *domain.Post {
	return &domain.Post{
		Title:       postReq.Title,
		Content:     postReq.Content,
		Description: postReq.Description,
		Author:      postReq.Author,
		CategoryID:  apiwrap.ConvertBsonID(postReq.CategoryID),
		TagsID: lo.Map(postReq.TagsID, func(id string, _ int) bson.ObjectID {
			return apiwrap.ConvertBsonID(id)
		}),
		IsPublish: postReq.IsPublish,
		IsTop:     postReq.IsTop,
		Thumbnail: postReq.Thumbnail,
	}
}

func (h *PostHandler) PostDetailToVO(post *domain.PostDetail) *PostVO {
	return &PostVO{
		ID:          post.ID.Hex(),
		CreatedAt:   post.CreatedAt,
		UpdatedAt:   post.UpdatedAt,
		Title:       post.Title,
		Content:     post.Content,
		Description: post.Description,
		Author:      post.Author,
		Category:    GetCategoryNameFromLabel(post.Category),
		Tags:        GetTagNamesFromLabels(post.Tags),
		IsTop:       post.IsTop,
		Thumbnail:   post.Thumbnail,
	}
}

func (h *PostHandler) PostDetailListToVOList(posts []*domain.PostDetail) []*PostVO {
	return lo.Map(posts, func(post *domain.PostDetail, _ int) *PostVO {
		return h.PostDetailToVO(post)
	})
}
