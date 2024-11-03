package main

import (
	"go_fiber_project/dal"
	"go_fiber_project/database"
	"go_fiber_project/services"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Connect to the database
	database.Connect()
	database.Migrate(&dal.Todo{})
	defer database.Close()

	app := fiber.New()

	app.Get("/", services.HelloWorld)

	app.Post("/todos", services.CreateTodo)

	app.Get("/todos", services.GetTodos)

	app.Listen(":8001")
}
