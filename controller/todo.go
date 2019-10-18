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
	completed, _ := strconv.Atoi(c.PostForm("Completed"))
	todo := model.Todo{
		Title:     c.PostForm("Title"),
		Completed: completed,
	}
	config.DB.Create(&todo)

	c.JSON(http.StatusCreated, todo)
}

// GetAllTodo is used for getting all todos
func GetAllTodo(c *gin.Context) {
	var todos []model.Todo
	var response []model.TodoSchema

	config.DB.Find(&todos)

	if len(todos) == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"Message": "No todos found",
		})
		return
	}

	for _, item := range todos {
		completed := false
		if item.Completed == 1 {
			completed = true
		}
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
	ID := c.Param("id")

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

	config.DB.Delete(&todo)

	c.JSON(http.StatusOK, todo)
}
