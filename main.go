package main

import (
	"fifth_test/database"
	_ "fifth_test/docs"
	"fifth_test/routes"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
)

func setUpRoutes(app *fiber.App) {
	api := app.Group("/api")
	routes.PostRoute(api.Group("/posts"))
}

// @title Godel Swagger Documentation
// @version 1.0
// @description Swagger doc
// @termsOfService http://swagger.io/terms/

// @contact.name Ivan
// @contact.url https://github.com/ivankonevv/godel_fifth
// @contact.email ivankonewv@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// @BasePath /api
// @schemes http
func main() {
	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())

	database.ConnectDB()

	app.Get("/swagger/*", swagger.Handler)

	setUpRoutes(app)

	if err := app.Listen(":3000"); err != nil {
		log.Fatal("Error server starting.")
	}
}
