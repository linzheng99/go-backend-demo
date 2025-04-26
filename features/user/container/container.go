package userContainer

import (
	"gorm.io/gorm"

	userHandler "github.com/linzheng99/go-backend-demo/features/user/handler"
	userRepo "github.com/linzheng99/go-backend-demo/features/user/repository"
	userService "github.com/linzheng99/go-backend-demo/features/user/service"
)

type Container struct {
	db             *gorm.DB
	UserRepository *userRepo.UserRepo
	UserService    *userService.UserService
	UserHandler    *userHandler.UserHandler
}

// create a new container
func NewContainer(db *gorm.DB) *Container {
	c := &Container{
		db: db,
	}
	c.initRepositories()
	c.initServices()
	c.initHandlers()
	return c
}

func (c *Container) initRepositories() {
	c.UserRepository = userRepo.NewUserRepository(c.db)
}

func (c *Container) initServices() {
	c.UserService = userService.NewUserService(c.UserRepository)
}

func (c *Container) initHandlers() {
	c.UserHandler = userHandler.NewUserHandler(c.UserService)
}
