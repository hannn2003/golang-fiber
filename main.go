package main

import (
	"fiber-demo/configs"
	"fiber-demo/controllers"
	"fiber-demo/queries"
	"fiber-demo/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"gorm.io/gorm"

	_ "fiber-demo/docs" // Import the generated docs
)

var DB *gorm.DB

func Inject(db *gorm.DB) {
	DB = db
}
// @title Example API
// @version 1.0
// @description Fiber Swagger Example
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @host localhost:3000
// @BasePath /
func main() {
		err := godotenv.Load(".env")
		if err != nil {
			panic("Error loading .env file")
		}

		database := configs.InitDatabase()
		Inject(database)
		
		userRepository := queries.NewUserRepository(database)
		userController := controllers.NewUserController(userRepository)

		app := fiber.New()

		app.Use(cors.New(cors.Config{
			AllowOrigins: "*",
			AllowCredentials: false,
		}))

		// Initialize routes
		routes.SwaggerRoute(app)
		routes.PublicRoutes(app, userController)
		
		app.Static("/", "./public")
		app.Listen(":3000")
}