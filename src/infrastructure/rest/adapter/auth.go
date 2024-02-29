package adapter

import (
	authService "github.com/chaiyawatkit/clean-architecture-gorm-mysql/src/application/usecases/auth"
	userRepository "github.com/chaiyawatkit/clean-architecture-gorm-mysql/src/infrastructure/repository/user"
	authController "github.com/chaiyawatkit/clean-architecture-gorm-mysql/src/infrastructure/rest/controllers/auth"
	"gorm.io/gorm"
)

func AuthAdapter(db *gorm.DB) *authController.Controller {
	uRepository := userRepository.Repository{DB: db}
	service := authService.Service{UserRepository: uRepository}
	return &authController.Controller{AuthService: service}
}
