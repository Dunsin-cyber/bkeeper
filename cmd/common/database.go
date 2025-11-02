package common

import (
	"fmt"
	"os"
	"log"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)



func NewDatabase() (*gorm.DB,  error){

	 err := godotenv.Load(".env")
   if err != nil {
    panic("Error loading .env file")
  }

  host :=  os.Getenv("DB_HOST")
  username := os.Getenv("DB_USERNAME")
  password := os.Getenv("DB_PASSWORD")
  database := os.Getenv("DB_DATABASE")
  port := os.Getenv("DB_PORT")

dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, username, password, database, port)
db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

if err != nil {
	return nil, err
}

log.Default().Println("Database connection established")

return db, nil

}