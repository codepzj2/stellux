package ioc

import (
	"context"
	"fmt"
	"github.com/chenmingyong0423/go-mongox/v2"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
	"server/global"
)

func NewMongoDB() *mongox.Database {
	mongoClient, err := mongo.Connect(options.Client().ApplyURI(global.Env.URL).SetAuth(options.Credential{
		Username:   global.Env.MongoUserName,
		Password:   global.Env.MongoRootPassword,
		AuthSource: global.Env.MongoDatabase,
	}))
	if err != nil {
		fmt.Println(global.Env)
		panic(errors.Wrap(err, "数据库连接失败"))
	}
	if err = mongoClient.Ping(context.Background(), readpref.Primary()); err != nil {
		panic(errors.Wrap(err, "数据库无法ping通"))
	}
	return mongox.NewClient(mongoClient, &mongox.Config{}).NewDatabase(global.Env.MongoDatabase)
}
