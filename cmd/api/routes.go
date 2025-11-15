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
		publicAuthRoutes.POST("/forgot/password", handler.ForgotPasswordHandler)
		publicAuthRoutes.POST("/reset/password", handler.ResetPasswordHandler)
	}

	profileRoutes := apiGroup.Group("/profile", app.appMiddleware.AuthenticationMiddleware)
	{
		profileRoutes.GET("/authenticated/user", handler.GetAuthenticatedUser)
		profileRoutes.PATCH("/change/password", handler.UpdatePasswordHandler)
	}

	categoryRoutes := apiGroup.Group("/categories", app.appMiddleware.AuthenticationMiddleware)
	{
		categoryRoutes.GET("/all", handler.ListCategories)
		categoryRoutes.POST("/store", handler.CreateCategory)
		categoryRoutes.DELETE("/delete/:id", handler.DeleteCategory)
	}

	app.server.GET("/", handler.HealthCheck)

}
