package todo

import (
	"errors"
	useCaseTodo "github.com/chaiyawatkit/clean-architecture-gorm-mysql/src/application/usecases/todo"
	"github.com/chaiyawatkit/clean-architecture-gorm-mysql/src/domain"
	domainError "github.com/chaiyawatkit/clean-architecture-gorm-mysql/src/domain/errors"
	domainTodo "github.com/chaiyawatkit/clean-architecture-gorm-mysql/src/domain/todo"
	"github.com/chaiyawatkit/clean-architecture-gorm-mysql/src/infrastructure/rest/controllers"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	TodoService useCaseTodo.Service
}

// NewTodo godoc
// @Tags todolist
// @Summary Create New Todolist
// @Security ApiKeyAuth
// @Description Create new todoList on the system
// @Accept  json
// @Produce  json
// @Param data body NewTodoRequest true "body data"
// @Success 200 {object} domainTodo.Todolist
// @Failure 400 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Router /todolist [post]
func (c *Controller) NewTodo(ctx *gin.Context) {
	var request NewTodoRequest

	if err := controllers.BindJSON(ctx, &request); err != nil {
		appError := domainError.NewAppError(err, domainError.ValidationError)
		_ = ctx.Error(appError)
		return
	}
	newTodo := domainTodo.NewTodo{
		Name:        request.Name,
		Description: request.Description,
	}

	domainTodo, err := c.TodoService.Create(&newTodo)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	responseTodo := domainToResponseMapper(domainTodo)

	ctx.JSON(http.StatusOK, responseTodo)
}

// GetAllTodo godoc
// @Tags todolist
// @Summary Get all TodoList
// @Security ApiKeyAuth
// @Description Get all TodoList
// @Success 200 {object} []ResponseTodo
// @Failure 400 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Router /todolist [get]
func (c *Controller) GetAllTodo(ctx *gin.Context) {
	todos, err := c.TodoService.GetAll()

	if err != nil {
		appError := domainError.NewAppErrorWithType(domainError.UnknownError)
		_ = ctx.Error(appError)
		return
	}

	ctx.JSON(http.StatusOK, arrayDomainToResponseMapper(todos))

}

// GetDataTodo godoc
// @Tags todolist
// @Summary Get all TodoList by query
// @Security ApiKeyAuth
// @Description Get all TodoList by query
// @Param data body DataTodoRequest true "body data"
// @Success 200 {object} []useCaseTodo.PaginationResultTodo
// @Failure 400 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Router /todolist/data [post]
func (c *Controller) GetDataTodos(ctx *gin.Context) {
	var request DataTodoRequest

	if err := controllers.BindJSON(ctx, &request); err != nil {
		appError := domainError.NewAppError(err, domainError.ValidationError)
		_ = ctx.Error(appError)
		return
	}

	var dateRangeFiltersDomain []domain.DateRangeFilter = make([]domain.DateRangeFilter, len(request.FieldsDateRange))
	for i, dateRangeFilter := range request.FieldsDateRange {
		dateRangeFiltersDomain[i] = domain.DateRangeFilter{Field: dateRangeFilter.Field, Start: dateRangeFilter.StartDate, End: dateRangeFilter.EndDate}
	}

	todos, err := c.TodoService.GetData(request.Page, request.Limit, request.SorBy.Field, request.SorBy.Direction, request.Filters, request.GlobalSearch, dateRangeFiltersDomain)
	if err != nil {
		appError := domainError.NewAppErrorWithType(domainError.UnknownError)
		_ = ctx.Error(appError)
		return
	}

	numPages, nextCursor, prevCursor := controllers.PaginationValues(request.Limit, request.Page, todos.Total)

	var response = PaginationResultTodo{
		Data:       arrayDomainToResponseMapper(todos.Data),
		Total:      todos.Total,
		Limit:      request.Limit,
		Current:    request.Page,
		NextCursor: nextCursor,
		PrevCursor: prevCursor,
		NumPages:   numPages,
	}

	ctx.JSON(http.StatusOK, response)
}

// GetTodoByID godoc
// @Tags todolist
// @Summary Get TodoList by ID
// @Security ApiKeyAuth
// @Description Get TodoList by ID on the system
// @Param todo_id path int true "id of todolist"
// @Success 200 {object} domainTodo.Todolist
// @Failure 400 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Router /todolist/{id} [get]
func (c *Controller) GetTodoByID(ctx *gin.Context) {
	todoID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		appError := domainError.NewAppError(errors.New("todo id is invalid"), domainError.ValidationError)
		_ = ctx.Error(appError)
		return
	}

	domainTodo, err := c.TodoService.GetByID(todoID)
	if err != nil {
		appError := domainError.NewAppError(err, domainError.ValidationError)
		_ = ctx.Error(appError)
		return
	}

	ctx.JSON(http.StatusOK, domainTodo)
}

// UpdateTodo godoc
// @Tags todolist
// @Summary Update  TodoList by id
// @Security ApiKeyAuth
// @Description Update  TodoList by id
// @Param data body NewTodoRequest true "body data"
// @Success 200 {object} []useCaseTodo.PaginationResultTodo
// @Failure 400 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Router /todolist/{id} [post]
func (c *Controller) UpdateTodo(ctx *gin.Context) {
	todoID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		appError := domainError.NewAppError(errors.New("param id is necessary in the url"), domainError.ValidationError)
		_ = ctx.Error(appError)
		return
	}
	var requestMap map[string]any

	err = controllers.BindJSONMap(ctx, &requestMap)
	if err != nil {
		appError := domainError.NewAppError(err, domainError.ValidationError)
		_ = ctx.Error(appError)
		return
	}

	err = updateValidation(requestMap)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	var todo *domainTodo.Todolist
	todo, err = c.TodoService.Update(todoID, requestMap)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, todo)

}

// DeleteTodo godoc
// @Tags todolist
// @Summary Get TodoList by ID
// @Security ApiKeyAuth
// @Description Delete TodoList by ID
// @Param id path int true "id of todolist"
// @Security ApiKeyAuth
// @Success 200 {object} MessageResponse
// @Failure 400 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Router /todolist/{id} [delete]
func (c *Controller) DeleteTodo(ctx *gin.Context) {
	todoID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		appError := domainError.NewAppError(errors.New("param id is necessary in the url"), domainError.ValidationError)
		_ = ctx.Error(appError)
		return
	}

	err = c.TodoService.Delete(todoID)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "resource deleted successfully"})
}
