package main

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/Dunsin-cyber/bkeeper/cmd/api/handlers"
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

  err := godotenv.Load(".env")
   if err != nil {
    e.Logger.Fatal("Error loading .env file", err)
  }

  db, err := common.NewDatabase()

  if err != nil {
    e.Logger.Fatal("Could not connect to the database", err.Error())
  }
  
  port := os.Getenv("PORT")
//   DBUrl := os.Getenv("DATABASE_URL")

  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  e.GET("/", hello)

  hdlr := handlers.Handler{
    DB: db,
  }

  app := Application{
    logger: e.Logger,
    server: e,
    handler: hdlr,
  }

  fmt.Println(app)

  // Start server
  if err := e.Start(fmt.Sprintf("localhost:%s", port)); err != nil && !errors.Is(err, http.ErrServerClosed) {
    slog.Error("failed to start server", "error", err)
  }
}

// Handler
func hello(c echo.Context) error {
  return c.String(http.StatusOK, "Hello, World!")
}