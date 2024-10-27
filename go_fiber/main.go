package main

import (
	"go_fiber_project/dal"
	"go_fiber_project/dal/dtos"
	"go_fiber_project/database"
	"go_fiber_project/validate"

	"github.com/gofiber/fiber/v2"
)

type ApiResult struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Errors  []string    `json:"errors"`
}

func CreateSuccess(data interface{}) *ApiResult {
	return &ApiResult{
		Success: true,
		Data:    data,
	}
}

func CreateError(errors []string) *ApiResult {
	return &ApiResult{
		Success: false,
		Errors:  errors,
	}
}

func main() {
	// Connect to the database
	database.Connect()
	database.Migrate(&dal.Todo{})
	defer database.Close()

	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(CreateSuccess("Hello, World!"))
	})

	app.Post("/todos", func(ctx *fiber.Ctx) error {
		t := new(dtos.TodoCreateDTO)
		if err := ctx.BodyParser(t); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(CreateError([]string{"Failed to parse body"}))
		}

		isValid, valErrors := validate.Validate(t)
		if !isValid {
			return ctx.Status(fiber.StatusBadRequest).JSON(CreateError(valErrors))
		}

		todo, err := t.MapTodoCreateDTO()
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(CreateError([]string{err.Error()}))
		}

		createResult := database.DB.Create(&todo)
		if createResult.Error != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(CreateError([]string{createResult.Error.Error()}))
		}

		return ctx.Status(fiber.StatusCreated).JSON(CreateSuccess(todo))
	})

	app.Listen(":8001")
}
