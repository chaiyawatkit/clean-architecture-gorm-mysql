package todo

import "github.com/chaiyawatkit/clean-architecture-gorm-mysql/src/infrastructure/rest/controllers"

type NewTodoRequest struct {
	Name        string `json:"name" example:"Something" gorm:"unique" binding:"required"`
	Description string `json:"description" example:"Something" binding:"required"`
}

type DataTodoRequest struct {
	Limit           int64                                   `json:"limit" example:"10"`
	Page            int64                                   `json:"page" example:"1"`
	GlobalSearch    string                                  `json:"globalSearch" example:"chaiyawatkit"`
	Filters         map[string][]string                     `json:"filters"`
	SorBy           controllers.SortByDataRequest           `json:"sortBy"`
	FieldsDateRange []controllers.FieldDateRangeDataRequest `json:"fieldsDateRange"`
}
