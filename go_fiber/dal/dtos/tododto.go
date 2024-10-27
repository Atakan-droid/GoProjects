package dtos

import (
	"go_fiber_project/dal"
)

type TodoCreateDTO struct {
	Title       string `json:"title" validate:"required,min=3,max=100"`
	Description string `json:"description" validate:"required,min=3,max=100"`
}

func (t *TodoCreateDTO) MapTodoCreateDTO() (*dal.Todo, error) {
	return &dal.Todo{
		Title:       t.Title,
		Description: t.Description,
	}, nil

}
