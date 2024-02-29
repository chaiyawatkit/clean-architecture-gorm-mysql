package routes

import (
	_ "github.com/chaiyawatkit/clean-architecture-gorm-mysql/docs"
	"github.com/chaiyawatkit/clean-architecture-gorm-mysql/src/infrastructure/rest/adapter"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

// @SecurityDefinitions.jwt
type Security struct {
	Authorization string `header:"Authorization" json:"Authorization"`
}

// @title  Golang Clean-Architecture
// @version 1.0
// @description Documentation's  Golang Clean-Architecture
// @termsOfService http://swagger.io/terms/

// @contact.name Chaiyawatkit
// @contact.url https://github.com/chaiyawatkit
// @contact.email chaiyawatkit160340@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// ApplicationV1Router is a function that contains all routes of the application
// @host localhost:8080
// @BasePath /v1
func ApplicationV1Router(router *gin.Engine, db *gorm.DB) {
	routerV1 := router.Group("/v1")

	{
		// Documentation Swagger
		{
			routerV1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		}

		AuthRoutes(routerV1, adapter.AuthAdapter(db))
		UserRoutes(routerV1, adapter.UserAdapter(db))
		TodoRoutes(routerV1, adapter.TodoAdapter(db))
	}
}
