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

func GetTodoById(id int, dest any) (any, error) {
	result := database.DB.Model(&Todo{}).Where("id = ?", id).First(dest)
	return dest, result.Error
}

func UpdateTodo(id int, todo *Todo) *gorm.DB {
	return database.DB.Model(&Todo{}).Where("id = ?", id).Updates(todo)
}

func IsExists(id int) bool {
	var count int64
	database.DB.Model(&Todo{}).Where("id = ?", id).Count(&count)
	return count > 0
}

func DeleteTodo(id int) *gorm.DB {
	return database.DB.Where("id = ?", id).Delete(&Todo{})
}
