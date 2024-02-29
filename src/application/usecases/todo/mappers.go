package todo

import (
	domainTodo "github.com/chaiyawatkit/clean-architecture-gorm-mysql/src/domain/todo"
)

func (n *NewTodo) toDomainMapper() *domainTodo.Todolist {
	return &domainTodo.Todolist{
		Name:        n.Name,
		Description: n.Description,
	}
}
