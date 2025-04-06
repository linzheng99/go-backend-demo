package main

import (
	user "github.com/linzheng99/go-backend-demo/features/user/routes"
)

func (app *application) routes() {
	api := app.router.Group("/api")

	user.Routes(api, app.db)

}
