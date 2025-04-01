package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type application struct {
	config config
	router *gin.Engine
}

type config struct {
	addr string
}

func (app *application) run() error {
	// 设置路由
	app.setupRoutes()

	// 创建 HTTP 服务器
	server := &http.Server{
		Addr:         app.config.addr,
		Handler:      app.router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Printf("Starting server on %s", app.config.addr)

	return server.ListenAndServe()
}

func (app *application) setupRoutes() {
	app.router = gin.New()

	app.setupMiddleware()

	app.routes()
}

func (app *application) setupMiddleware() {
	app.router.Use(gin.Logger())
	app.router.Use(gin.Recovery())
}
