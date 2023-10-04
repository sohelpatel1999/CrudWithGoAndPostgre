package model

import (
	"github.com/jinzhu/gorm"
)

type Todo struct {
	Title  string
	Status string
}

func Createtool(db *gorm.DB, todo *Todo) error {
	return db.Create(todo).Error
}
