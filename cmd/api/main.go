package main

import (
	"github.com/linzheng99/go-backend-demo/internal/bootstrap"
	"github.com/linzheng99/go-backend-demo/internal/global"
)

func main() {
	bootstrap.Init()

	cfg := config{
		addr: global.Config.Server.Port,
	}

	app := &application{
		config: cfg,
	}

	// 运行应用
	if err := app.run(); err != nil {
		panic(err)
	}

}
