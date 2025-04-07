package userHandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	userModel "github.com/linzheng99/go-backend-demo/features/user/model"
	userService "github.com/linzheng99/go-backend-demo/features/user/service"
)

type UserHandler struct {
	userService *userService.UserService
}

func NewUserHandler(userService *userService.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	user := userModel.User{}
	c.BindJSON(&user)
	err := h.userService.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "User created successfully"})
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := h.userService.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"users": users})
}
