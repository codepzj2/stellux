package global

import (
	"log/slog"

	"github.com/spf13/viper"
)

func InitViper(mode string) {
	if mode == "development" {
		viper.SetConfigFile("config/stellux.development.yaml")
		slog.Info("使用开发环境配置文件")
	} else if mode == "production" {
		viper.SetConfigFile("config/stellux.production.yaml")
		slog.Info("使用生产环境配置文件")
	} else {
		panic("mode参数无效")
	}
	err := viper.ReadInConfig()
	if err != nil {
		panic("viper读取配置文件失败")
	}
}
