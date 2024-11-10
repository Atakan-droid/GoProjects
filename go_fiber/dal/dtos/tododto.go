package dtos

import (
	"go_fiber_project/dal"
)

type TodoCreateDTO struct {
	Title       string `json:"title" validate:"required,min=3,max=100"`
	Description string `json:"description" validate:"required,min=3,max=100"`
}

type TodoDto struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

type TodoUpdateDto struct {
	Title       string `json:"title" validate:"required,min=3,max=100"`
	Description string `json:"description" validate:"required,min=3,max=100"`
	Completed   bool   `json:"completed"`
}

func (t *TodoCreateDTO) MapTodoCreateDTO() (*dal.Todo, error) {
	return &dal.Todo{
		Title:       t.Title,
		Description: t.Description,
	}, nil

}

func (t *TodoUpdateDto) MapTodo() (*dal.Todo, error) {
	return &dal.Todo{
		Title:       t.Title,
		Description: t.Description,
		Completed:   t.Completed,
	}, nil
}
