package routes

import (
	todoController "github.com/chaiyawatkit/clean-architecture-gorm-mysql/src/infrastructure/rest/controllers/todo"
	"github.com/chaiyawatkit/clean-architecture-gorm-mysql/src/infrastructure/rest/middlewares"
	"github.com/gin-gonic/gin"
)

func TodoRoutes(router *gin.RouterGroup, controller *todoController.Controller) {

	routerTodo := router.Group("/todolist")
	routerTodo.Use(middlewares.AuthJWTMiddleware())
	{
		routerTodo.GET("/", controller.GetAllTodo)
		routerTodo.POST("/", controller.NewTodo)
		routerTodo.GET("/:id", controller.GetTodoByID)
		routerTodo.PUT("/:id", controller.UpdateTodo)
		routerTodo.DELETE("/:id", controller.DeleteTodo)
		routerTodo.POST("/data", controller.GetDataTodos)
	}

}
