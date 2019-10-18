package main

import (
	"go-todo/config"
	"go-todo/controller"

	"github.com/gin-gonic/gin"
)

func init() {
	config.DatabaseConnect()
}

func main() {
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		todo := v1.Group("/todo")
		{
			todo.POST("/", controller.CreateTodo)
			todo.GET("/", controller.GetAllTodo)
			todo.GET("/:id", controller.GetTodoByID)
			todo.PUT("/:id", controller.UpdateTodo)
			todo.DELETE("/:id", controller.DeleteTodo)
		}
	}
	router.Run()
}
