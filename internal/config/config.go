package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/linzheng99/go-backend-demo/internal/tool"
	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}

// 初始化配置
func InitConfig() (c *Config) {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	rootDir, err := tool.GetRootDir()
	if err != nil {
		fmt.Println("Get root dir error: ", err)
		return nil
	}

	v := viper.New()
	// 设置配置文件名称和类型
	v.SetConfigName("config." + env)
	v.SetConfigType("yaml")
	v.AddConfigPath(filepath.Join(rootDir, "internal", "config"))

	if err := v.ReadInConfig(); err != nil {
		fmt.Println("Read config file error: ", err)
		return nil
	}

	var config Config

	if err := v.Unmarshal(&config); err != nil {
		fmt.Println("Unmarshal config file error: ", err)
		return nil
	}

	return &config
}
