package todo

import (
	domainTodo "github.com/chaiyawatkit/clean-architecture-gorm-mysql/src/domain/todo"
)

func (todo *Todo) toDomainMapper() *domainTodo.Todolist {
	return &domainTodo.Todolist{
		ID:          todo.ID,
		Name:        todo.Name,
		Description: todo.Description,
		CreatedAt:   todo.CreatedAt,
	}
}

func fromDomainMapper(medicine *domainTodo.Todolist) *Todo {
	return &Todo{
		ID:          medicine.ID,
		Name:        medicine.Name,
		Description: medicine.Description,
		CreatedAt:   medicine.CreatedAt,
	}
}

func arrayToDomainMapper(todos *[]Todo) *[]domainTodo.Todolist {
	todosDomain := make([]domainTodo.Todolist, len(*todos))
	for i, todo := range *todos {
		todosDomain[i] = *todo.toDomainMapper()
	}

	return &todosDomain
}
