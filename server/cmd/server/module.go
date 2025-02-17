package main

import (
	"fmt"
	"log"

	"server/internal/ioc"

	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	engine    *gin.Engine
	envConfig *ioc.EnvConfig
}

func NewHttpServer(engine *gin.Engine, envConfig *ioc.EnvConfig) *HttpServer {
	return &HttpServer{
		engine:    engine,
		envConfig: envConfig,
	}
}

func (s *HttpServer) Start() {
	port := s.envConfig.Port
	if port == "" {
		log.Fatal("未配置端口，请检查.env文件中的PORT设置")
	}
	log.Printf("Server is running on port %s", port)
	if err := s.engine.Run(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatalf("启动服务器失败: %v", err)
	}
}
