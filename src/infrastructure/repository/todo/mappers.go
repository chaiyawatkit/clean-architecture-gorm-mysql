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

func fromDomainMapper(todo *domainTodo.Todolist) *Todo {
	return &Todo{
		ID:          todo.ID,
		Name:        todo.Name,
		Description: todo.Description,
		CreatedAt:   todo.CreatedAt,
	}
}

func arrayToDomainMapper(todos *[]Todo) *[]domainTodo.Todolist {
	todosDomain := make([]domainTodo.Todolist, len(*todos))
	for i, todo := range *todos {
		todosDomain[i] = *todo.toDomainMapper()
	}

	return &todosDomain
}
