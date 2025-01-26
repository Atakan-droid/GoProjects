package controllers

import (
	"fmt"
	"go_fiber_project/services"

	"github.com/gofiber/fiber/v2"
)

func ConfigureTodoController(app *fiber.App) {
	app.Get("/", services.HelloWorld)

	app.Post(services.TodoServiceKey, services.CreateTodo)

	app.Get(services.TodoServiceKey, services.GetTodos)

	var getTodo = fmt.Sprintf("%s/:todoId", services.TodoServiceKey)
	app.Get(getTodo, services.GetTodoById)

	var updateTodo = fmt.Sprintf("%s/:todoId", services.TodoServiceKey)
	app.Put(updateTodo, services.UpdateTodo)

	var deleteTodo = fmt.Sprintf("%s/:todoId", services.TodoServiceKey)
	app.Delete(deleteTodo, services.DeleteTodo)
}
