package dal

import (
	"go_fiber_project/database"

	"gorm.io/gorm"
)

type Todo struct {
	ID          int    `json:"id" gorm:"primaryKey"`
	Title       string `json:"title" gorm:"not null max:255"`
	Description string `json:"description" gorm:"default:null max:1024"`
	Completed   bool   `json:"completed" gorm:"default:false"`
}

func CreateTodo(todo *Todo) *gorm.DB {
	return database.DB.Create(todo)
}

func GetTodos(dest any) (any, error) {
	result := database.DB.Model(&Todo{}).Find(dest)
	return dest, result.Error
}
