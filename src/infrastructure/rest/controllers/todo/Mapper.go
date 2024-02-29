package todo

import (
	todoDomain "github.com/chaiyawatkit/clean-architecture-gorm-mysql/src/domain/todo"
)

func domainToResponseMapper(clientDomain *todoDomain.Todolist) (createClientResponse *ResponseTodo) {
	createClientResponse = &ResponseTodo{
		ID:          clientDomain.ID,
		Name:        clientDomain.Name,
		Description: clientDomain.Description,
		CreatedAt:   clientDomain.CreatedAt,
		UpdatedAt:   clientDomain.UpdatedAt}

	return
}

func arrayDomainToResponseMapper(clientsDomain *[]todoDomain.Todolist) *[]ResponseTodo {
	clientsResponse := make([]ResponseTodo, len(*clientsDomain))
	for i, client := range *clientsDomain {
		clientsResponse[i] = *domainToResponseMapper(&client)
	}
	return &clientsResponse
}
