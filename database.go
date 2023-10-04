package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func ConnectionDatabase() (*gorm.DB, error) {

	db, err := gorm.Open("postgre", "host=localhost port=5432 user=postgre dbname=admin sslmode=disable")
	if err != nil {
		return nil, err
	}
	return db, nil
}
