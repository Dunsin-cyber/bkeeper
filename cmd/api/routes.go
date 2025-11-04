package main

import (
	"github.com/Dunsin-cyber/bkeeper/cmd/api/handlers"
)

func (app *Application) routes(handler handlers.Handler) {
	app.server.GET("/", handler.HealthCheck)
	app.server.POST("/register", handler.RegisterHandler)

}
