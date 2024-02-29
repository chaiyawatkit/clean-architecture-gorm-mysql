package user

import "github.com/chaiyawatkit/clean-architecture-gorm-mysql/src/infrastructure/rest/controllers"

type NewUserRequest struct {
	UserName  string `json:"user" example:"Kuk" gorm:"unique" binding:"required"`
	Email     string `json:"email" example:"test@gmail.com" gorm:"unique" binding:"required"`
	FirstName string `json:"firstName" example:"Kuk" binding:"required"`
	LastName  string `json:"lastName" example:"Mmama" binding:"required"`
	Password  string `json:"password" example:"Password123" binding:"required"`
	Role      string `json:"role" example:"admin" binding:"required"`
}

type DataUserRequest struct {
	Limit           int64                                   `json:"limit" example:"10"`
	Page            int64                                   `json:"page" example:"1"`
	GlobalSearch    string                                  `json:"globalSearch" example:"Kuk"`
	Filters         map[string][]string                     `json:"filters"`
	SorBy           controllers.SortByDataRequest           `json:"sortBy"`
	FieldsDateRange []controllers.FieldDateRangeDataRequest `json:"fieldsDateRange"`
}
