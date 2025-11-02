package main

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/Dunsin-cyber/bkeeper/cmd/api/handlers"
	"github.com/Dunsin-cyber/bkeeper/cmd/api/middlewares"
	"github.com/Dunsin-cyber/bkeeper/cmd/common"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Application struct {
  logger echo.Logger
  server *echo.Echo
  handler handlers.Handler

}

func main() {
  e := echo.New()
  
  e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
  Format: "time=${time_rfc3339} remote_ip=${remote_ip} method=${method}, uri=${uri}, status=${status}\n"}))
  e.Use(middleware.Recover())
  e.Use(middlewares.CustomMiddleware)

  err := godotenv.Load(".env")
   if err != nil {
    e.Logger.Fatal("Error loading .env file", err)
  }

  db, err := common.NewDatabase()

  if err != nil {
    e.Logger.Fatal(err.Error())
  }
  
  
  h := handlers.Handler{
    DB: db,
  }
  
  app := Application{
      logger: e.Logger,
      server: e,
      handler: h,
  }
  
  app.routes(h)
  fmt.Println(app)
    
    // Start server
    port := os.Getenv("PORT")
  if err := e.Start(fmt.Sprintf("localhost:%s", port)); err != nil && !errors.Is(err, http.ErrServerClosed) {
    slog.Error("failed to start server", "error", err)
  }
}
