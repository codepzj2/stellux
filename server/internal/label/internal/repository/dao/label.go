package dao

import (
	"context"
	"errors"

	"github.com/chenmingyong0423/go-mongox/v2"
	"github.com/chenmingyong0423/go-mongox/v2/builder/query"
	"github.com/chenmingyong0423/go-mongox/v2/builder/update"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Label struct {
	ID        bson.ObjectID `bson:"_id"`
	LabelType string        `bson:"label_type"`
	Name      string        `bson:"name"`
}

type LabelUpdate struct {
	LabelType string `bson:"label_type"`
	Name      string `bson:"name"`
}

type ILabelDao interface {
	Create(ctx context.Context, label *Label) error
	Update(ctx context.Context, id bson.ObjectID, label *Label) error
	Delete(ctx context.Context, id bson.ObjectID) error
	Get(ctx context.Context, id bson.ObjectID) (*Label, error)
	GetList(ctx context.Context, labelType string, limit int64, skip int64) ([]*Label, int64, error)
	GetAllByType(ctx context.Context, labelType string) ([]*Label, error)
}

var _ ILabelDao = (*LabelDao)(nil)

func NewLabelDao(db *mongox.Database) *LabelDao {
	return &LabelDao{coll: mongox.NewCollection[Label](db, "label")}
}

type LabelDao struct {
	coll *mongox.Collection[Label]
}

func (d *LabelDao) Create(ctx context.Context, label *Label) error {
	label.ID = bson.NewObjectID()
	res, err := d.coll.Creator().InsertOne(ctx, label)
	if err != nil {
		return err
	}
	if res.InsertedID == nil {
		return errors.New("标签创建失败")
	}
	return nil
}

func (d *LabelDao) Update(ctx context.Context, id bson.ObjectID, label *Label) error {
	res, err := d.coll.Updater().Filter(query.Id(id)).Updates(update.SetFields(d.LabelToUpdate(label))).UpdateOne(ctx)
	if err != nil {
		return err
	}
	if res.ModifiedCount == 0 {
		return errors.New("标签更新失败")
	}
	return nil
}

func (d *LabelDao) Delete(ctx context.Context, id bson.ObjectID) error {
	res, err := d.coll.Deleter().Filter(query.Id(id)).DeleteOne(ctx)
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return errors.New("标签删除失败")
	}
	return nil
}

func (d *LabelDao) Get(ctx context.Context, id bson.ObjectID) (*Label, error) {
	return d.coll.Finder().Filter(query.Id(id)).FindOne(ctx)
}

func (d *LabelDao) GetList(ctx context.Context, labelType string, limit int64, skip int64) ([]*Label, int64, error) {
	count, err := d.coll.Finder().Filter(query.Eq("label_type", labelType)).Count(ctx)
	if err != nil {
		return nil, 0, err
	}
	labelList, err := d.coll.Finder().Filter(query.Eq("label_type", labelType)).Limit(limit).Skip(skip).Find(ctx)
	if err != nil {
		return nil, 0, err
	}
	return labelList, count, nil
}

func (d *LabelDao) GetAllByType(ctx context.Context, labelType string) ([]*Label, error) {
	return d.coll.Finder().Filter(query.Eq("label_type", labelType)).Find(ctx)
}

func (d *LabelDao) LabelToUpdate(label *Label) *LabelUpdate {
	return &LabelUpdate{
		LabelType: label.LabelType,
		Name:      label.Name,
	}
}
