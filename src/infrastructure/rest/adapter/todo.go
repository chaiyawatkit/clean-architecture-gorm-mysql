package adapter

import (
	todoService "github.com/chaiyawatkit/clean-architecture-gorm-mysql/src/application/usecases/todo"
	todoRepository "github.com/chaiyawatkit/clean-architecture-gorm-mysql/src/infrastructure/repository/todo"
	todoController "github.com/chaiyawatkit/clean-architecture-gorm-mysql/src/infrastructure/rest/controllers/todo"
	"gorm.io/gorm"
)

func TodoAdapter(db *gorm.DB) *todoController.Controller {
	mRepository := todoRepository.Repository{DB: db}
	service := todoService.Service{TodoRepository: mRepository}
	return &todoController.Controller{TodoService: service}
}
