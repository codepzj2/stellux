package ioc

import (
	"context"
	"github.com/chenmingyong0423/go-mongox/v2"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

func NewMongoDB(envConfig *EnvConfig) *mongox.Database {
	mongoClient, err := mongo.Connect(options.Client().ApplyURI(envConfig.DSN).SetAuth(options.Credential{
		Username:   envConfig.MongoUserName,
		Password:   envConfig.MongoRootPassword,
		AuthSource: envConfig.MongoDatabase,
	}))
	if err != nil {
		panic(errors.Wrap(err, "数据库连接失败"))
	}
	if err = mongoClient.Ping(context.Background(), readpref.Primary()); err != nil {
		panic(errors.Wrap(err, "数据库无法ping通"))
	}
	return mongox.NewClient(mongoClient, &mongox.Config{}).NewDatabase(envConfig.MongoDatabase)
}
