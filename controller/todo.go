package controller

import (
	"go-todo/config"
	"go-todo/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateTodo is used for creating a new todo
func CreateTodo(c *gin.Context) {
	// strconv.Atoi change a string into integer.
	// strconv.Atoi return two result, which is value and error.
	// The underscore has represent error, we're not gonna used it. So we put _
	completed, _ := strconv.Atoi(c.PostForm("Completed"))

	// Assign request payload into struture of todo model.
	todo := model.Todo{
		Title:     c.PostForm("Title"),
		Completed: completed,
	}

	// This line below is equivalent to "INSERT INTO table VALUES (value1, value1)".
	// We use & sign to get the address memory of todo model.
	config.DB.Create(&todo)

	// Return to the client as a JSON object.
	c.JSON(http.StatusCreated, todo)
}

// GetAllTodo is used for getting all todos
func GetAllTodo(c *gin.Context) {
	// model.Todo is a struct.
	// if we use [] sign in front of them,
	// they will return as a Slices (array) of Struct.
	var todos []model.Todo
	var response []model.TodoSchema

	// This line below is equivalent to "SELECt * FROM table".
	config.DB.Find(&todos)

	// If data is not found, the program will stop and return an error message.
	if len(todos) == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"Message": "No todos found",
		})
		return
	}

	// We use TodoSchema structure to return a todo data in response payload.
	// The underscore is an index. We're not gonna used it.
	for _, item := range todos {
		// This operation is convert 1 or 0
		// into true or false
		completed := false
		if item.Completed == 1 {
			completed = true
		}
		// append() also know as push(),
		// it used to add a new value to a slices.
		response = append(response, model.TodoSchema{
			ID:        item.ID,
			Title:     item.Title,
			Completed: completed,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"Data": response,
	})
}

// GetTodoByID is used for getting a todo by id
func GetTodoByID(c *gin.Context) {
	var todo model.Todo

	// Get the paramater
	ID := c.Param("id")

	// This line below is equivalent to "SELECT * FROM table WHERE id = id".
	config.DB.First(&todo, ID)

	if todo.ID == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"Message": "Item not found",
		})
		return
	}

	completed := false
	if todo.Completed == 1 {
		completed = true
	}

	response := model.TodoSchema{
		ID:        todo.ID,
		Title:     todo.Title,
		Completed: completed,
	}

	c.JSON(http.StatusOK, response)
}

// UpdateTodo is used for updating a todo
func UpdateTodo(c *gin.Context) {
	var todo model.Todo
	ID := c.Param("id")

	config.DB.First(&todo, ID)

	if todo.ID == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"Message": "Item not found",
		})
		return
	}

	completed, _ := strconv.Atoi(c.PostForm("Completed"))

	// This line below is equivalent to "UPDATE table SET field = value"
	config.DB.Model(&todo).Update("title", c.PostForm("Title"))
	config.DB.Model(&todo).Update("completed", completed)

	c.JSON(http.StatusOK, todo)
}

// DeleteTodo is used for deleting a todo
func DeleteTodo(c *gin.Context) {
	var todo model.Todo
	ID := c.Param("id")

	config.DB.First(&todo, ID)

	if todo.ID == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"Message": "Item not found",
		})
		return
	}

	// This line below is equivalent to "DELETE * FROM table WHERE id = model.id"
	config.DB.Delete(&todo)

	c.JSON(http.StatusOK, todo)
}
