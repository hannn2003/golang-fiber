package main

import (
	"fiber-demo/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	_ "fiber-demo/docs" // Import the generated docs
)

// @title Example API
// @version 1.0
// @description Fiber Swagger Example
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @host localhost:3000
// @BasePath /

func main() {
		app := fiber.New()

		app.Use(cors.New(cors.Config{
			AllowOrigins: "*",
			AllowCredentials: false,
		}))

		// Initialize routes
		routes.SwaggerRoute(app)
		routes.PublicRoutes(app)

		app.Static("/", "./public")
		app.Listen(":3000")
}