package user

import (
	"github.com/gin-gonic/gin"
	user "github.com/linzheng99/go-backend-demo/features/user/service"
	"gorm.io/gorm"
)

func Routes(api *gin.RouterGroup, db *gorm.DB) {

	api.GET("/users/list", func(c *gin.Context) {
		users, err := user.GetAllUsers(db)
		if err != nil {
			c.JSON(500, gin.H{
				"message": "Failed to get users",
			})
		}

		c.JSON(200, gin.H{
			"users": users,
		})
	})

}
