package main

import (
	userContainer "github.com/linzheng99/go-backend-demo/features/user"
	userRoutes "github.com/linzheng99/go-backend-demo/features/user/routes"
)

func (app *application) routes() {
	api := app.router.Group("/api")

	userContainer := userContainer.NewContainer(app.db)
	userRoutes.Routes(api, userContainer.UserHandler)

}
