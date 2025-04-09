package main

import (
	"github.com/gin-gonic/gin"
	userContainer "github.com/linzheng99/go-backend-demo/features/user/container"
	userRoutes "github.com/linzheng99/go-backend-demo/features/user/routes"
)

func (app *application) routes() {
	app.router = gin.New()

	setupMiddleware(app)

	api := app.router.Group("/api")

	userContainer := userContainer.NewContainer(app.db)
	userRoutes.Routes(api, userContainer.UserHandler)

}

func setupMiddleware(app *application) {
	app.router.Use(gin.Logger())
	app.router.Use(gin.Recovery())
}
