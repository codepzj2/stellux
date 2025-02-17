package ioc

import (
	"encoding/json"
	"github.com/joho/godotenv"
)

type EnvConfig struct {
	DSN               string `json:"DSN"`
	MongoDatabase     string `json:"MONGO_INITDB_DATABASE"`
	MongoRootUsername string `json:"MONGO_INITDB_ROOT_USERNAME"`
	MongoRootPassword string `json:"MONGO_INITDB_ROOT_PASSWORD"`
	MongoUserName     string `json:"MONGO_USERNAME"`
	MongoPassword     string `json:"MONGO_PASSWORD"`
}

// InitEnv 初始化配置文件
func InitEnv() *EnvConfig {
	var envConfig EnvConfig
	err := godotenv.Load(".env")
	if err != nil {
		panic("请检查config文件夹下是否含有.env")
	}
	myEnv, _ := godotenv.Read()
	jsonEnv, _ := json.Marshal(myEnv)
	json.Unmarshal(jsonEnv, &envConfig)
	return &envConfig
}
