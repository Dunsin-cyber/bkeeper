package main

import (
	"log"

	"github.com/Dunsin-cyber/bkeeper/common"
	"github.com/Dunsin-cyber/bkeeper/internal/models"
)

func main() {
	db, err := common.NewDatabase()
	if err != nil {
		panic("Failed to connect to database")
	}
	err = db.AutoMigrate(&models.UserModel{})

	if err != nil {
		panic(err)
	}

	log.Println("Migration completed")

}
