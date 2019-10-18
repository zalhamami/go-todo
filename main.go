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
		v1.POST("/todo", controller.CreateTodo)
		v1.GET("/todo", controller.GetAllTodo)
		v1.GET("/todo/:id", controller.GetTodoByID)
		v1.PUT("/todo/:id", controller.UpdateTodo)
		v1.DELETE("/todo/:id", controller.DeleteTodo)
	}
	router.Run()
}
