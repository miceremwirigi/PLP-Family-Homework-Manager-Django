package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/miceremwirigi/PLP-Family-Homework-Manager-Django/go-version/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	app := fiber.New()
	db, err := gorm.Open(postgres.Open("host=localhost dbname=homework_manager user=homework_manager password=homework_manager port=5432 sslmode=disable timezone=Africa/Nairobi"), &gorm.Config{})
	if err != nil {
		log.Println("\n\nInitializing Database Failed")
		log.Fatal(err)
	} else {
		log.Println("\n\nInitializing Database Success")
	}

	fmt.Println("\nMaking MIgrations...")
	models.MakeMigrations(db, models.My_models)

	if err := app.Listen(":8080"); err != nil {
		panic(err)
	}
}
