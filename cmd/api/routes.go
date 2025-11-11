package main

import (
	"github.com/Dunsin-cyber/bkeeper/cmd/api/handlers"
)

func (app *Application) routes(handler handlers.Handler) {
	apiGroup := app.server.Group("/api/v1")

	publicAuthRoutes := apiGroup.Group("/auth")
	{
		publicAuthRoutes.POST("/register", handler.RegisterHandler)
		publicAuthRoutes.POST("/login", handler.LoginHandler)
	}

	profileRoutes := apiGroup.Group("/profile", app.appMiddleware.AuthenticationMiddleware)
	{
		profileRoutes.GET("/authenticated/user", handler.GetAuthenticatedUser)
		profileRoutes.PATCH("/change/password", handler.UpdatePasswordHandler)
	}

	app.server.GET("/", handler.HealthCheck)

}
