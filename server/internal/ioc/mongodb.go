package ioc

import (
	"github.com/chenmingyong0423/go-mongox/v2"
	"server/global"
)

func NewMongoDB() *mongox.Database {
	return mongox.NewClient(global.MongoClient, &mongox.Config{}).NewDatabase(global.Env.MongoDatabase)
}
