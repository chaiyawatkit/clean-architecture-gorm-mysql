package controllers

type JSONSwagger struct {
}

type MessageResponse struct {
	Message string `json:"message"`
}

type SortByDataRequest struct {
	Field     string `json:"field" example:"name"`
	Direction string `json:"direction" example:"asc"`
}

type FieldDateRangeDataRequest struct {
	Field     string `json:"field" example:"createdAt"`
	StartDate string `json:"startDate" example:"2021-01-01"`
	EndDate   string `json:"endDate" example:"2021-01-01"`
}
