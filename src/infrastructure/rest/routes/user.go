package routes

import (
	userController "github.com/chaiyawatkit/clean-architecture-gorm-mysql/src/infrastructure/rest/controllers/user"
	"github.com/chaiyawatkit/clean-architecture-gorm-mysql/src/infrastructure/rest/middlewares"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.RouterGroup, controller *userController.Controller) {
	routerUser := router.Group("/user")
	routerUser.Use(middlewares.AuthJWTMiddleware())
	{
		routerUser.POST("/", controller.NewUser)
		routerUser.GET("/:id", controller.GetUsersByID)
		routerUser.GET("/", controller.GetAllUsers)
		routerUser.PUT("/:id", controller.UpdateUser)
		routerUser.DELETE("/:id", controller.DeleteUser)
	}
}
