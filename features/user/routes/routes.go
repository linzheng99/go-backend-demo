package userRoutes

import (
	"github.com/gin-gonic/gin"
	userHandler "github.com/linzheng99/go-backend-demo/features/user/handler"
)

func Routes(router *gin.RouterGroup, handler *userHandler.UserHandler) {
	userGroup := router.Group("/users")
	{
		userGroup.POST("/", handler.CreateUser)
		userGroup.GET("/list", handler.GetAllUsers)
	}
}
