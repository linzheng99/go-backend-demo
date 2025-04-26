package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/linzheng99/go-backend-demo/internal/config"
	"github.com/linzheng99/go-backend-demo/internal/tool"
	"gorm.io/gorm"
)

type application struct {
	config *config.Config
	router *gin.Engine
	db     *gorm.DB
}

func (app *application) run() error {
	// 设置路由
	app.routes()

	// 创建 HTTP 服务器
	server := &http.Server{
		Addr:         ":" + app.config.Server.Port,
		Handler:      app.router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// 获取本机 IP 地址
	localIP, err := tool.GetLocalIP()
	if err != nil {
		log.Printf("Failed to get local IP: %v", err)
	}
	log.Printf("Starting server on %s", localIP+":"+app.config.Server.Port)

	return server.ListenAndServe()
}
