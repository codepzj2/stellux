package repository

import (
	"context"

	"github.com/codepzj/stellux/server/internal/label/internal/domain"
	"github.com/codepzj/stellux/server/internal/label/internal/repository/dao"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type ILabelRepository interface {
	Create(ctx context.Context, label *domain.Label) error
	Update(ctx context.Context, id string, label *domain.Label) error
	Delete(ctx context.Context, id string) error
	Get(ctx context.Context, id string) (*domain.Label, error)
	GetList(ctx context.Context, labelType string, pageNo int64, pageSize int64) ([]*domain.Label, int64, error)
}

var _ ILabelRepository = (*LabelRepository)(nil)

func NewLabelRepository(dao dao.ILabelDao) *LabelRepository {
	return &LabelRepository{dao: dao}
}

type LabelRepository struct {
	dao dao.ILabelDao
}

func (r *LabelRepository) Create(ctx context.Context, label *domain.Label) error {
	return r.dao.Create(ctx, r.DomainToDao(label))
}

func (r *LabelRepository) Update(ctx context.Context, id string, label *domain.Label) error {
	bid, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	return r.dao.Update(ctx, bid, r.DomainToDao(label))
}

func (r *LabelRepository) Delete(ctx context.Context, id string) error {
	bid, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	return r.dao.Delete(ctx, bid)
}

func (r *LabelRepository) Get(ctx context.Context, id string) (*domain.Label, error) {
	bid, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	label, err := r.dao.Get(ctx, bid)
	if err != nil {
		return nil, err
	}
	return r.DaoToDomain(label), nil
}

func (r *LabelRepository) GetList(ctx context.Context, labelType string, pageNo int64, pageSize int64) ([]*domain.Label, int64, error) {
	labels, count, err := r.dao.GetList(ctx, labelType, pageSize, (pageNo-1)*pageSize)
	if err != nil {
		return nil, 0, err
	}
	domainLabels := make([]*domain.Label, 0, len(labels))
	for _, label := range labels {
		domainLabels = append(domainLabels, r.DaoToDomain(label))
	}
	return domainLabels, count, nil
}

func (r *LabelRepository) DomainToDao(label *domain.Label) *dao.Label {
	return &dao.Label{
		LabelType: string(label.LabelType),
		Name:      label.Name,
	}
}

func (r *LabelRepository) DaoToDomain(label *dao.Label) *domain.Label {
	return &domain.Label{
		ID:        label.ID.Hex(),
		LabelType: label.LabelType,
		Name:      label.Name,
	}
}
