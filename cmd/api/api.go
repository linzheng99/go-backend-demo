package main

import (
	"log"
	"net"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/linzheng99/go-backend-demo/internal/config"
	"gorm.io/gorm"
)

type application struct {
	config *config.Config
	router *gin.Engine
	db     *gorm.DB
}

func (app *application) run() error {
	// 设置路由
	setupRoutes(app)

	// 创建 HTTP 服务器
	server := &http.Server{
		Addr:         ":" + app.config.Server.Port,
		Handler:      app.router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// 获取本机 IP 地址
	localIP, err := getLocalIP()
	if err != nil {
		log.Printf("Failed to get local IP: %v", err)
	}
	log.Printf("Starting server on %s", localIP+":"+app.config.Server.Port)

	return server.ListenAndServe()
}

func setupRoutes(app *application) {
	app.router = gin.New()

	setupMiddleware(app)

	app.routes()
}

func setupMiddleware(app *application) {
	app.router.Use(gin.Logger())
	app.router.Use(gin.Recovery())
}

func getLocalIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}
	return "", nil
}
