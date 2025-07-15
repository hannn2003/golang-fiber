package routes

import (
	"fiber-demo/controllers"

	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(a *fiber.App) {
	route := a.Group("/api/v1")
	route.Get("/", controllers.HealthApi)
	route.Get("/products", controllers.GetProducts)
}