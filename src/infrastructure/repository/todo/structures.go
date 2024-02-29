package todo

import (
	domainTodo "github.com/chaiyawatkit/clean-architecture-gorm-mysql/src/domain/todo"
	"time"
)

type Todo struct {
	ID          int       `json:"id" example:"123" gorm:"primaryKey"`
	Name        string    `json:"name" example:"Paracetamol" gorm:"unique"`
	Description string    `json:"description" example:"Some Description"`
	CreatedAt   time.Time `json:"created_at,omitempty" example:"2021-02-24 20:19:39" gorm:"autoCreateTime:mili"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" example:"2021-02-24 20:19:39" gorm:"autoUpdateTime:mili"`
}

func (*Todo) TableName() string {
	return "todo_lists"
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
