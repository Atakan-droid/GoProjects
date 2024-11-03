package services

import (
	"go_fiber_project/dal"
	"go_fiber_project/dal/dtos"
	"go_fiber_project/models"
	"go_fiber_project/validate"

	"github.com/gofiber/fiber/v2"
)

func HelloWorld(ctx *fiber.Ctx) error {
	return ctx.JSON(models.CreateSuccess("Hello, World!"))
}

func CreateTodo(ctx *fiber.Ctx) error {
	t := new(dtos.TodoCreateDTO)
	if err := ctx.BodyParser(t); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(models.CreateError([]string{"Failed to parse body"}))
	}

	isValid, valErrors := validate.Validate(t)
	if !isValid {
		return ctx.Status(fiber.StatusBadRequest).JSON(models.CreateError(valErrors))
	}

	todo, err := t.MapTodoCreateDTO()
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(models.CreateError([]string{err.Error()}))
	}

	createResult := dal.CreateTodo(todo)
	if createResult.Error != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(models.CreateError([]string{createResult.Error.Error()}))
	}

	return ctx.Status(fiber.StatusCreated).JSON(models.CreateSuccess(todo))
}

func GetTodos(ctx *fiber.Ctx) error {
	todos := []dtos.TodoDto{}

	res, err := dal.GetTodos(&todos)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(models.CreateError([]string{err.Error()}))
	}

	return ctx.JSON(models.CreateSuccess(res))
}
