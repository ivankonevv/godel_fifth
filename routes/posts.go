package routes

import (
	"fifth_test/controllers"
	"github.com/gofiber/fiber/v2"
)

func PostRoute(route fiber.Router) {
	route.Get("/", controllers.GetAllPosts)
	route.Get("/:postId", controllers.GetPost)
	route.Post("/", controllers.CreatePost)
}