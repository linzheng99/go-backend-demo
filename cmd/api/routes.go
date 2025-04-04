package main

import (
	"github.com/gin-gonic/gin"
)

func (app *application) routes() {
	api := app.router.Group("/api")
	{
		// 示例路由
		api.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Welcome to the API",
			})
		})
	}
}
