package ioc

import (
	"context"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
	"server/global"
)

func NewMongoDB() *mongo.Database {
	MongoClient, err := mongo.Connect(options.Client().ApplyURI(global.Env.URL))
	if err != nil {
		panic(errors.Wrap(err, "数据库连接失败"))
	}
	if err = MongoClient.Ping(context.Background(), readpref.Primary()); err != nil {
		panic(errors.Wrap(err, "数据库无法ping通"))
	}

	return MongoClient.Database(global.Env.MongoDatabase)
}
