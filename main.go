package main

import (
	"fifth_test/database"
	"fifth_test/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
)

func setUpRoutes(app *fiber.App) {
	api := app.Group("/api")
	routes.PostRoute(api.Group("/posts"))
}

func main() {
	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())

	database.ConnectDB()

	setUpRoutes(app)

	if err := app.Listen(":3000"); err != nil {
		log.Fatal("Error server starting.")
	}
}
