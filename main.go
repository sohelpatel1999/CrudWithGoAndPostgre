package main

import (
	"fmt"
	"gowithpostgrecrud/database"
	"gowithpostgrecrud/model"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	fmt.Println("sohel bole")
	db, err := database.ConnectionDatabase()
	if err != nil {
		panic("Error While Building Connection")
	}
	fmt.Println("connection build succesfully", db)
	defer db.Close()

	router.POST("/sohel", func(c *gin.Context) {
		var todo model.Todo

		if err := c.ShouldBindJSON(&todo); err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
			return
		}
		if err := model.Createtool(db, &todo); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": "server Error"})
			return
		}

		c.JSON(http.StatusOK, todo)

	})

	router.GET("/todo", func(c *gin.Context) {
		todos, err := model.GetTodos(db)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"Error": "database not found"})
			return
		}
		c.JSON(http.StatusOK, todos)
	})

	router.GET("/todo/:id", func(c *gin.Context) {
		todoID := c.Param("id")

		id , err := strconv.Atoi(todoID)
		if err != nil {
			c.JSON(http.StatusNotFound,gin.H{"Error": "at time of conversion"})
			
		}
		todo, err := model.GetTodosbyid(db,id)
		if err!= nil {
			c.JSON(http.StatusNotFound,gin.H{"Error":"Gettitng the data from model"})
		}
		c.JSON(http.StatusOK,todo)

	})

	router.POST("/todo", func(c *gin.Context) {
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
	router.Run(":8080")
}
