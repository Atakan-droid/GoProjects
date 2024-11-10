package main

import (
	"go_fiber_project/controllers"
	"go_fiber_project/dal"
	"go_fiber_project/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger" // swagger handler
)

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8001
// @BasePath /
func main() {
	// Connect to the database
	database.Connect()
	database.Migrate(&dal.Todo{})
	defer database.Close()

	app := fiber.New()

	controllers.ConfigureTodoController(app)

	app.Get("/swagger/*", swagger.HandlerDefault) // default

	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:         "http://example.com/doc.json",
		DeepLinking: false,
		// Expand ("list") or Collapse ("none") tag groups by default
		DocExpansion: "none",
	}))

	app.Listen(":8001")
}
