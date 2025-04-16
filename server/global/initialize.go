package global

import (
	"fmt"

	"github.com/casbin/casbin/v2"
	"github.com/spf13/viper"
)

var (
	Enforcer *casbin.Enforcer
)

func init() {
	// 解析命令行参数
	ParseFlag()
	// 初始化logger
	InitLogger(*Mode)
	// 初始化viper
	InitViper(*Mode)
	// 初始化enforcer
	Enforcer = NewEnforcer(fmt.Sprintf("mongodb://%s:%s@%s:%d/%s", viper.GetString("mongodb.MONGO_USERNAME"), viper.GetString("mongodb.MONGO_PASSWORD"), viper.GetString("mongodb.HOST"), viper.GetInt("mongodb.PORT"), viper.GetString("mongodb.MONGO_INITDB_DATABASE")))
}
