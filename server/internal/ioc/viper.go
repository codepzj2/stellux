package ioc

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

type EnvConfig struct {
	DSN                 string `json:"DSN"`
	MongoDatabase       string `json:"MONGO_INITDB_DATABASE"`
	MongoRootUsername   string `json:"MONGO_INITDB_ROOT_USERNAME"`
	MongoRootPassword   string `json:"MONGO_INITDB_ROOT_PASSWORD"`
	MongoUserName       string `json:"MONGO_USERNAME"`
	MongoPassword       string `json:"MONGO_PASSWORD"`
	Port                string `json:"PORT"`
	TokenExpireDuration string `json:"TOKEN_EXPIRE_DURATION"`
}

func InitEnv() *EnvConfig {
	viper.SetConfigFile("./config/stellux.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("找不到环境配置文件: %w", err))
	}

	// 设置默认值
	viper.SetDefault("TokenExpireDuration", "86400s")

	// 加载配置
	envConfig := &EnvConfig{
		DSN:               viper.GetString("mongodb.DSN"),
		MongoDatabase:     viper.GetString("mongodb.MONGO_INITDB_DATABASE"),
		MongoRootUsername: viper.GetString("mongodb.MONGO_INITDB_ROOT_USERNAME"),
		MongoRootPassword: viper.GetString("mongodb.MONGO_INITDB_ROOT_PASSWORD"),
		MongoUserName:     viper.GetString("mongodb.MONGO_USERNAME"),
		MongoPassword:     viper.GetString("mongodb.MONGO_PASSWORD"),
		Port:              viper.GetString("server.PORT"),
	}
	log.Println(envConfig)
	// 监听配置文件
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Println("配置文件发生改变:", e.Name)
	})
	return envConfig
}
