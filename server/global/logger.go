package global

import "log/slog"

func InitLogger(mode string) {
	if mode == "development" {
		slog.SetLogLoggerLevel(slog.LevelDebug)
	} else if mode == "production" {
		slog.SetLogLoggerLevel(slog.LevelInfo)
	} else {
		panic("mode参数无效")
	}
}
