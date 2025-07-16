package routes

import (
	"fiber-demo/controllers"

	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(a *fiber.App, userController *controllers.UserController) {
	route := a.Group("/api/v1")

	route.Get("/users", userController.GetUsers)
	route.Get("/users/:id", userController.GetUserById)
	route.Post("/create-user", userController.UserSignUp)
}
