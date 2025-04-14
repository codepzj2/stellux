package global

import (
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
	Enforcer = NewEnforcer(viper.GetString("mongodb.URL"))
}
