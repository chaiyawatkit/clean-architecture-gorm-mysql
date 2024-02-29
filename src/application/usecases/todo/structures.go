package todo

import (
	domainTodo "github.com/chaiyawatkit/clean-architecture-gorm-mysql/src/domain/todo"
)

type NewTodo struct {
	Name        string `json:"name" example:"Paracetamol"`
	Description string `json:"description" example:"Some Description"`
	EANCode     string `json:"ean_code" example:"9900000124"`
	Laboratory  string `json:"laboratory" example:"Roche"`
}

type PaginationResultTodo struct {
	Data       *[]domainTodo.Todolist
	Total      int64
	Limit      int64
	Current    int64
	NextCursor uint
	PrevCursor uint
	NumPages   int64
}
