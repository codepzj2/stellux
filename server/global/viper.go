package global

import (
	"context"
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

var (
	Env         *EnvConfig
	MongoClient *mongo.Client = nil
)

type EnvConfig struct {
	URL               string `json:"URL"`
	MongoDatabase     string `json:"MONGO_INITDB_DATABASE"`
	MongoRootUsername string `json:"MONGO_INITDB_ROOT_USERNAME"`
	MongoRootPassword string `json:"MONGO_INITDB_ROOT_PASSWORD"`
	MongoUserName     string `json:"MONGO_USERNAME"`
	MongoPassword     string `json:"MONGO_PASSWORD"`
	Port              string `json:"PORT"`
}

func init() {
	viper.SetConfigFile("./config/stellux.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("找不到环境配置文件: %w", err))
	}

	// 加载配置
	Env = &EnvConfig{
		URL:               viper.GetString("mongodb.URL"),
		MongoDatabase:     viper.GetString("mongodb.MONGO_INITDB_DATABASE"),
		MongoRootUsername: viper.GetString("mongodb.MONGO_INITDB_ROOT_USERNAME"),
		MongoRootPassword: viper.GetString("mongodb.MONGO_INITDB_ROOT_PASSWORD"),
		MongoUserName:     viper.GetString("mongodb.MONGO_USERNAME"),
		MongoPassword:     viper.GetString("mongodb.MONGO_PASSWORD"),
		Port:              viper.GetString("server.PORT"),
	}

	// 监听配置文件
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Println("配置文件发生改变:", e.Name)
	})

	// 使用 mongo 连接数据库
	client, err := mongo.Connect(options.Client().ApplyURI(Env.URL).SetAuth(options.Credential{
		Username:   Env.MongoUserName,
		Password:   Env.MongoRootPassword,
		AuthSource: Env.MongoDatabase,
	}))
	if err != nil {
		panic(errors.Wrap(err, "数据库连接失败"))
	}
	if err = client.Ping(context.Background(), readpref.Primary()); err != nil {
		panic(errors.Wrap(err, "数据库无法ping通"))
	}

	MongoClient = client
}
