package bootstrap

import (
	"github.com/linzheng99/go-backend-demo/internal/config"
	"github.com/linzheng99/go-backend-demo/internal/global"
)

func Init() {
	global.Config = config.InitConfig()
}
