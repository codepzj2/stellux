package dao

import (
	"context"
	"errors"

	"github.com/chenmingyong0423/go-mongox/v2"
	"github.com/chenmingyong0423/go-mongox/v2/builder/query"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Document struct {
	mongox.Model `bson:",inline"`
	Title        string
	Content      string
	ParentID     bson.ObjectID `bson:"parent_id"`
	DocumentID   bson.ObjectID `bson:"document_id"`
}

type IDocumentDao interface {
	Create(ctx context.Context, doc *Document) error
	FindAllByDocumentID(ctx context.Context, documentID bson.ObjectID) ([]*Document, error)
}

var _ IDocumentDao = (*DocumentDao)(nil)

func NewDocumentDao(db *mongox.Database) *DocumentDao {
	return &DocumentDao{coll: mongox.NewCollection[Document](db, "document")}
}

type DocumentDao struct {
	coll *mongox.Collection[Document]
}

func (d *DocumentDao) Create(ctx context.Context, doc *Document) error {
	insertResult, err := d.coll.Creator().InsertOne(ctx, doc)
	if err != nil {
		return err
	}
	if insertResult.InsertedID == nil {
		return errors.New("新增文档失败")
	}
	return nil
}

func (d *DocumentDao) FindAllByDocumentID(ctx context.Context, documentID bson.ObjectID) ([]*Document, error) {
	documentList, err := d.coll.Finder().Filter(query.NewBuilder().Or(query.Id(documentID), query.Eq("document_id", documentID)).Build()).Find(ctx)
	if err != nil {
		return nil, err
	}
	return documentList, nil
}
