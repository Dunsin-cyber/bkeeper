package main

import (
	"github.com/Dunsin-cyber/bkeeper/cmd/api/requests"
	"github.com/Dunsin-cyber/bkeeper/cmd/api/services"
	"github.com/Dunsin-cyber/bkeeper/common"
)

func main() {
	db, err := common.NewDatabase()
	if err != nil {
		panic(err)

	}
	categoryService := services.NewCategoryService(db)
	categories := []string{"Food", "Gifts", "Health", "Eat Out",
		"Medical", "Home",
	}

	for _, category := range categories {
		_, err := categoryService.Create(requests.CreateCatagoryRequest{
			Name:     category,
			IsCustom: false,
		})

		if err != nil {
			panic(err)

		}

		println("category " + category + " created")
	}
}
