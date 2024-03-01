package todo

import (
	"github.com/chaiyawatkit/clean-architecture-gorm-mysql/src/domain"
	"time"
)

type Todolist struct {
	ID          int
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type NewTodo struct {
	Name        string
	Description string
}

func (newTodo *NewTodo) ToDomainMapper() *Todolist {
	return &Todolist{
		Name:        newTodo.Name,
		Description: newTodo.Description,
	}
}

type DataTodo struct {
	Data  *[]Todolist
	Total int64
}

type Service interface {
	GetAll() (*[]Todolist, error)
	GetData(page int64, limit int64, sortBy string, sortDirection string, filters map[string][]string, searchText string, dateRangeFilters []domain.DateRangeFilter) (*DataTodo, error)
	GetByID(id int) (*Todolist, error)
	Create(todo *NewTodo) (*Todolist, error)
	GetByMap(todoMap map[string]any) (*Todolist, error)
	Delete(id int) error
	Update(id int, todoMap map[string]any) (*Todolist, error)
}
