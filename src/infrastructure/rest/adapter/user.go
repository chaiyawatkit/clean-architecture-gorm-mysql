package adapter

import (
	userService "github.com/chaiyawatkit/clean-architecture-gorm-mysql/src/application/usecases/user"
	userRepository "github.com/chaiyawatkit/clean-architecture-gorm-mysql/src/infrastructure/repository/user"
	userController "github.com/chaiyawatkit/clean-architecture-gorm-mysql/src/infrastructure/rest/controllers/user"
	"gorm.io/gorm"
)

func UserAdapter(db *gorm.DB) *userController.Controller {
	uRepository := userRepository.Repository{DB: db}
	service := userService.Service{UserRepository: uRepository}
	return &userController.Controller{UserService: service}
}
