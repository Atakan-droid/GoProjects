package services

import (
	"go_fiber_project/dal"
	"go_fiber_project/dal/dtos"
	"go_fiber_project/models"
	"go_fiber_project/validate"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

const TodoServiceKey = "/todos"

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

func GetTodoById(ctx *fiber.Ctx) error {
	id := ctx.Params("todoId")
	if id == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(models.CreateError([]string{"id required"}))
	}

	intId, err := strconv.Atoi(id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(models.CreateError([]string{"Invalid id"}))
	}

	todo := new(dtos.TodoDto)
	res, err := dal.GetTodoById(intId, todo)
	if err != nil || res == nil {
		return ctx.Status(fiber.StatusNotFound).JSON(models.CreateError([]string{err.Error()}))
	}

	return ctx.JSON(models.CreateSuccess(res))
}

func UpdateTodo(ctx *fiber.Ctx) error {
	id := ctx.Params("todoId")
	if id == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(models.CreateError([]string{"id required"}))
	}

	//Atoi and ParseInt are similar, but ParseInt is more flexible
	intId, err := strconv.Atoi(id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(models.CreateError([]string{"Invalid id"}))
	}

	t := new(dtos.TodoUpdateDto)
	if err := ctx.BodyParser(t); err != nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(models.CreateError([]string{"Failed to parse body or body is empty"}))
	}

	isValid, valErrors := validate.Validate(t)
	if !isValid {
		return ctx.Status(fiber.StatusBadRequest).JSON(models.CreateError(valErrors))
	}

	todo, err := t.MapTodo()
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(models.CreateError([]string{err.Error()}))
	}

	updateResult := dal.UpdateTodo(intId, todo)
	if updateResult.Error != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(models.CreateError([]string{updateResult.Error.Error()}))
	}

	return ctx.JSON(models.CreateSuccess(todo))
}

func DeleteTodo(ctx *fiber.Ctx) error {
	id := ctx.Params("todoId")
	if id == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(models.CreateError([]string{"id required"}))
	}

	intId, err := strconv.Atoi(id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(models.CreateError([]string{"Invalid id"}))
	}

	isTodoExist := dal.IsExists(intId)
	if !isTodoExist {
		return ctx.Status(fiber.StatusNotFound).JSON(models.CreateError([]string{"Todo not found"}))
	}

	result := dal.DeleteTodo(intId)
	if result.Error != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(models.CreateError([]string{result.Error.Error()}))
	}

	return ctx.JSON(models.CreateSuccess(nil))
}
