package main

import (
	"fmt"

	"github.com/linzheng99/go-backend-demo/internal/config"
	"github.com/linzheng99/go-backend-demo/internal/database"
)

func main() {
	// 初始化配置
	cfg := config.InitConfig()

	// 连接数据库
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		cfg.Database.Host,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.DBName,
		cfg.Database.Port,
	)
	db, err := database.ConnectDatabase(dsn)
	if err != nil {
		panic(err)
	}
	err = database.AutoMigrate(db)
	if err != nil {
		panic(err)
	}

	app := &application{
		config: cfg,
		db:     db,
	}

	// 运行应用
	if err := app.run(); err != nil {
		panic(err)
	}

}
