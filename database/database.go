package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func ConnectionDatabase() (*gorm.DB, error) {

	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=sohel password=admin sslmode=disable")
	if err != nil {
		fmt.Println("error in db")
		return nil, err
	}
	return db, nil
}

func connection() (*gorm.DB, error){
	dbb, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=sohel password=admin sslmode=disable")
	if err != nil {
		fmt.Println("error in db")
		return nil, err
	}
	return dbb, err
}
