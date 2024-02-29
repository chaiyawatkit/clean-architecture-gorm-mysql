package todo

import "time"

type MessageResponse struct {
	Message string `json:"message"`
}

type ResponseTodo struct {
	ID          int       `json:"id" example:"1099"`
	Name        string    `json:"name" example:"Aspirina"`
	Description string    `json:"description" example:"Some Description"`
	CreatedAt   time.Time `json:"createdAt,omitempty" example:"2021-02-24 20:19:39" gorm:"autoCreateTime:mili"`
	UpdatedAt   time.Time `json:"updatedAt,omitempty" example:"2021-02-24 20:19:39" gorm:"autoUpdateTime:mili"`
}

type PaginationResultTodo struct {
	Data       *[]ResponseTodo `json:"data"`
	Total      int64           `json:"total"`
	Limit      int64           `json:"limit"`
	Current    int64           `json:"current"`
	NextCursor int64           `json:"nextCursor"`
	PrevCursor int64           `json:"prevCursor"`
	NumPages   int64           `json:"numPages"`
}
