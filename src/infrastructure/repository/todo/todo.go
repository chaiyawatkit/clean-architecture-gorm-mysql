package todo

import (
	"encoding/json"
	"fmt"
	"github.com/chaiyawatkit/clean-architecture-gorm-mysql/src/domain"
	domainErrors "github.com/chaiyawatkit/clean-architecture-gorm-mysql/src/domain/errors"
	domainTodo "github.com/chaiyawatkit/clean-architecture-gorm-mysql/src/domain/todo"
	"github.com/chaiyawatkit/clean-architecture-gorm-mysql/src/infrastructure/repository"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

var ColumnsTodoMapping = map[string]string{
	"id":          "id",
	"name":        "name",
	"description": "description",
	"createdAt":   "created_at",
	"updatedAt":   "updated_at",
}

var ColumnsTodoStructure = map[string]string{
	"id":          "ID",
	"name":        "Name",
	"description": "Description",

	"createdAt": "CreatedAt",
	"updatedAt": "UpdatedAt",
}

// GetData Fetch all medicine data
func (r *Repository) GetData(page int64, limit int64, sortBy string, sortDirection string, filters map[string][]string, searchText string, dateRangeFilters []domain.DateRangeFilter) (*domainTodo.DataTodo, error) {
	var users []Todo
	var total int64
	offset := (page - 1) * limit

	var searchColumns = []string{"name", "description"}

	countResult := make(chan error)
	go func() {
		err := r.DB.Model(&Todo{}).Debug().Scopes(repository.ApplyFilters(ColumnsTodoMapping, filters, dateRangeFilters, searchText, searchColumns)).Count(&total).Error
		countResult <- err
	}()

	queryResult := make(chan error)
	go func() {
		query, err := repository.ComplementSearch((*repository.Repository)(r), sortBy, sortDirection, limit, offset, filters, dateRangeFilters, searchText, searchColumns, ColumnsTodoMapping)
		if err != nil {
			queryResult <- err
			return
		}
		err = query.Find(&users).Error
		queryResult <- err
	}()

	var countErr, queryErr error
	for i := 0; i < 2; i++ {
		select {
		case err := <-countResult:
			countErr = err
		case err := <-queryResult:
			queryErr = err
		}
	}

	if countErr != nil {
		return &domainTodo.DataTodo{}, countErr
	}
	if queryErr != nil {
		return &domainTodo.DataTodo{}, queryErr
	}

	return &domainTodo.DataTodo{
		Data:  arrayToDomainMapper(&users),
		Total: total,
	}, nil
}

// Create ... Insert New data
func (r *Repository) Create(newTodo *domainTodo.Todolist) (createdTodo *domainTodo.Todolist, err error) {
	todo := fromDomainMapper(newTodo)
	tx := r.DB.Debug().Create(todo)

	if tx.Error != nil {
		byteErr, _ := json.Marshal(tx.Error)
		var newError domainErrors.GormErr
		err = json.Unmarshal(byteErr, &newError)

		if err != nil {
			return
		}
		switch newError.Number {
		case 1062:
			err = domainErrors.NewAppErrorWithType(domainErrors.ResourceAlreadyExists)
		default:
			err = domainErrors.NewAppErrorWithType(domainErrors.UnknownError)
		}
		return
	}

	createdTodo = todo.toDomainMapper()
	return
}

func (r *Repository) GetByID(id int) (*domainTodo.Todolist, error) {
	var todo Todo
	err := r.DB.Where("id = ?", id).First(&todo).Error

	if err != nil {
		switch err.Error() {
		case gorm.ErrRecordNotFound.Error():
			err = domainErrors.NewAppErrorWithType(domainErrors.NotFound)
		default:
			err = domainErrors.NewAppErrorWithType(domainErrors.UnknownError)
		}
		return &domainTodo.Todolist{}, err
	}

	return todo.toDomainMapper(), nil
}

// GetOneByMap ... Fetch only one medicine by Map
func (r *Repository) GetOneByMap(todoMap map[string]any) (*domainTodo.Todolist, error) {
	var todo Todo

	tx := r.DB.Limit(1)
	for key, value := range todoMap {
		if !repository.IsZeroValue(value) {
			tx = tx.Where(fmt.Sprintf("%s = ?", key), value)
		}
	}

	err := tx.Find(&todo).Error
	if err != nil {
		err = domainErrors.NewAppErrorWithType(domainErrors.UnknownError)
		return nil, err
	}
	return todo.toDomainMapper(), err
}

// Update ... Update medicine
func (r *Repository) Update(id int, todoMap map[string]any) (*domainTodo.Todolist, error) {
	var todo Todo

	todo.ID = id
	err := r.DB.Model(&todo).
		Select("name", "description").
		Updates(todoMap).Error

	// err = config.DB.Save(medicine).Error
	if err != nil {
		byteErr, _ := json.Marshal(err)
		var newError domainErrors.GormErr
		err = json.Unmarshal(byteErr, &newError)
		if err != nil {
			return &domainTodo.Todolist{}, err
		}
		switch newError.Number {
		case 1062:
			err = domainErrors.NewAppErrorWithType(domainErrors.ResourceAlreadyExists)
			return &domainTodo.Todolist{}, err

		default:
			err = domainErrors.NewAppErrorWithType(domainErrors.UnknownError)
			return &domainTodo.Todolist{}, err
		}
	}

	err = r.DB.Where("id = ?", id).First(&todo).Error

	return todo.toDomainMapper(), err
}

func (r *Repository) Delete(id int) (err error) {
	tx := r.DB.Debug().Delete(&domainTodo.Todolist{}, id)
	if tx.Error != nil {
		err = domainErrors.NewAppErrorWithType(domainErrors.UnknownError)
		return
	}

	if tx.RowsAffected == 0 {
		err = domainErrors.NewAppErrorWithType(domainErrors.NotFound)
	}

	return
}

func (r *Repository) GetAll() (*[]domainTodo.Todolist, error) {
	var Todos []Todo
	err := r.DB.Find(&Todos).Error
	if err != nil {
		err = domainErrors.NewAppErrorWithType(domainErrors.UnknownError)
		return nil, err
	}

	return arrayToDomainMapper(&Todos), nil
}
