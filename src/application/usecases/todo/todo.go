package todo

import (
	"github.com/chaiyawatkit/clean-architecture-gorm-mysql/src/domain"
	todoDomain "github.com/chaiyawatkit/clean-architecture-gorm-mysql/src/domain/todo"
	todoRepository "github.com/chaiyawatkit/clean-architecture-gorm-mysql/src/infrastructure/repository/todo"
)

type Service struct {
	TodoRepository todoRepository.Repository
}

var _ todoDomain.Service = &Service{}

func (s *Service) GetData(page int64, limit int64, sortBy string, sortDirection string, filters map[string][]string, searchText string, dateRangeFilters []domain.DateRangeFilter) (*todoDomain.DataTodo, error) {
	return s.TodoRepository.GetData(page, limit, sortBy, sortDirection, filters, searchText, dateRangeFilters)
}

func (s *Service) GetByID(id int) (*todoDomain.Todolist, error) {
	return s.TodoRepository.GetByID(id)
}

func (s *Service) Create(todo *todoDomain.NewTodo) (*todoDomain.Todolist, error) {
	todoModel := todo.ToDomainMapper()
	return s.TodoRepository.Create(todoModel)
}

func (s *Service) GetByMap(todoMap map[string]any) (*todoDomain.Todolist, error) {
	return s.TodoRepository.GetOneByMap(todoMap)
}

func (s *Service) Delete(id int) error {
	return s.TodoRepository.Delete(id)
}

func (s *Service) Update(id int, todoMap map[string]any) (*todoDomain.Todolist, error) {
	return s.TodoRepository.Update(id, todoMap)
}

func (s *Service) GetAll() (*[]todoDomain.Todolist, error) {
	return s.TodoRepository.GetAll()
}
