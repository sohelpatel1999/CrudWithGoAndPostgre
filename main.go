package main

import (
	"fmt"
	"gowithpostgrecrud/database"
)

func main() {
	fmt.Println("sohel bole")
	db, err := database.ConnectionDatabase()
	if err != nil {
		panic("Error While Building Connection")
	}
	fmt.Println("connection build succesfully", db)
	defer db.Close()
}
