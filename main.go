package main

import (
	"fmt"
	"gowithpostgrecrud/database"
	"gowithpostgrecrud/model"

	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	fmt.Println("sohel bole")
	db, err := database.ConnectionDatabase()
	if err != nil {
		panic("Error While Building Connection")
	}
	fmt.Println("connection build succesfully", db)
	defer db.Close()

	r.POST("/todo", func(c *gin.Context) {
		var todo model.Todo
		if err := c.ShouldBindJSON(&todo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid request"})
			return
		}
		if err := model.Createtool(db, &todo); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to Create todo"})
			return
		}
		c.JSON(http.StatusOK, todo)
	})
	r.Run(":8080")
}
