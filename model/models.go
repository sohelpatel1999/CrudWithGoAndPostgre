package model

import (
	"github.com/jinzhu/gorm"
)

type Todo struct {
	Id     int
	Title  string
	Status string
}

func Createtool(db *gorm.DB, todo *Todo) error {
	return db.Create(todo).Error
}

func GetTodos(db *gorm.DB) ([]Todo, error) {
	var todos []Todo
	err := db.Find(&todos).Error
	return todos, err
}

func GetTodosbyid(db *gorm.DB, id int) (Todo, error) {
	var todos Todo
	err := db.First(&todos, id).Error
	return todos, err
}

func UpdateTodo(db *gorm.DB, id int, todo *Todo) error {
	return db.Model(&Todo{}).Where("id = ?", id).Updates(todo).Error
}

func DeleteTodo(db *gorm.DB, id int) error {
	return db.Where("id = ?", id).Delete(&Todo{}).Error
}
