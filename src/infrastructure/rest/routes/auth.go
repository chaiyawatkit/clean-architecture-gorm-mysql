package routes

import (
	authController "github.com/chaiyawatkit/clean-architecture-gorm-mysql/src/infrastructure/rest/controllers/auth"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.RouterGroup, controller *authController.Controller) {

	routerAuth := router.Group("/auth")
	{
		routerAuth.POST("/login", controller.Login)
		routerAuth.POST("/access-token", controller.GetAccessTokenByRefreshToken)
	}

}
