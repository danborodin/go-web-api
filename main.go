package main

import (
	"log"
	"todoapi/database"
	"todoapi/routes"

	"github.com/gofiber/fiber/v2"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	app := fiber.New()

	err := database.Connect()

	if err != nil {
		log.Fatal(err)
	}

	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":8080"))

	defer database.DB.Close()
}
