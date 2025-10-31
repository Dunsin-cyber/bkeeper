package main

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
  e := echo.New()

  err := godotenv.Load(".env")
   if err != nil {
    e.Logger.Fatal("Error loading .env file", err)
  }

  
  port := os.Getenv("PORT")
//   DBUrl := os.Getenv("DATABASE_URL")

  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  e.GET("/", hello)

  // Start server
  if err := e.Start(fmt.Sprintf("localhost:%s", port)); err != nil && !errors.Is(err, http.ErrServerClosed) {
    slog.Error("failed to start server", "error", err)
  }
}

// Handler
func hello(c echo.Context) error {
  return c.String(http.StatusOK, "Hello, World!")
}