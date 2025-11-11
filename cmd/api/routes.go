package main

import (
	"github.com/Dunsin-cyber/bkeeper/cmd/api/handlers"
)

func (app *Application) routes(handler handlers.Handler) {
	app.server.GET("/", handler.HealthCheck)
	app.server.POST("/register", handler.RegisterHandler)
	app.server.POST("/login", handler.LoginHandler)
	app.server.GET("/authenticated/user", handler.GetAuthenticatedUser, app.appMiddleware.AuthenticationMiddleware)

}
