package global

import (
	"log/slog"

	"github.com/spf13/viper"
)

func InitViper(mode string) {
	switch mode {
	case "development":
		viper.SetConfigFile("config/stellux.development.yaml")
		slog.Info("使用开发环境配置文件")
	case "production":
		viper.SetConfigFile("config/stellux.production.yaml")
		slog.Info("使用生产环境配置文件")
	default:
		panic("mode参数无效")
	}
	err := viper.ReadInConfig()
	if err != nil {
		panic("viper读取配置文件失败")
	}
}
