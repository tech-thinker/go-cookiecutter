package repository

import (
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/tech-thinker/go-cookiecutter/constants"
	"github.com/tech-thinker/go-cookiecutter/db/models"
	"github.com/tech-thinker/go-cookiecutter/logger"
)

// TodoRepo is interface for todo repository
type TodoRepo interface {
	Save(doc *models.Todo) error
	FindOne(doc models.Todo) (models.Todo, error)
	Update(doc *models.Todo, fieldsToUpdate []string) error
	FindAll(query models.TodoQuery) ([]models.Todo, int64, error)
}

type todoRepo struct {
	db orm.Ormer
}

// Save method save the object into database
func (repo *todoRepo) Save(doc *models.Todo) error {
	groupError := "SAVE_TODO"
	createdAt := time.Now().Unix()
	updatedAt := time.Now().Unix()
	doc.CreatedAt=&createdAt
	doc.UpdatedAt=&updatedAt
	id, err := repo.db.Insert(doc)
	if err != nil {
		logger.Log.WithError(err).Error(groupError)
		return err
	}
	doc.ID = id
	return nil
}

// FindOne method returns first found result
func (repo *todoRepo) FindOne(doc models.Todo) (models.Todo, error) {
	groupError := "FIND_ONE_TODO"

	qs := repo.db.QueryTable(new(models.Todo))
	if doc.ID != 0 {
		qs = qs.Filter("id", doc.ID)
	}

	var todo models.Todo
	err := qs.One(&todo)
	if err != nil {
		logger.Log.WithError(err).Error(groupError)
		return todo, err
	}
	return todo, nil
}

// Update method update object into models
func (repo *todoRepo) Update(doc *models.Todo, fieldsToUpdate []string) error {
	groupError := "UPDATE_TODO"
	
	updatedAt := time.Now().Unix()
	doc.UpdatedAt=&updatedAt
	_, err := repo.db.Update(doc, fieldsToUpdate...)
	if err != nil {
		logger.Log.WithError(err).Error(groupError)
		return err
	}
	return nil
}

// FindAll method search records and pagination result will return
func (repo *todoRepo) FindAll(query models.TodoQuery) ([]models.Todo, int64, error) {
	groupError := "FIND_ALL_BY_QUERY_TODO"
	var todos []models.Todo

	sortingOrder := ""
	sortBy := "id"
	if query.Sort != nil {
		if query.Sort.SortOrder != nil && *query.Sort.SortOrder != constants.SortOrderDesc {
			sortingOrder = "-"
		}

		if query.Sort.OrderBy != nil {
			sortBy = *query.Sort.OrderBy
		}
	}

	qs := repo.db.QueryTable(new(models.Todo)).OrderBy(sortingOrder + sortBy)
	if query.Pagination != nil && query.Pagination.Page != nil && query.Pagination.Limit != nil {
		qs = qs.Offset((*query.Pagination.Page - 1) * *query.Pagination.Limit).Limit(*query.Pagination.Limit)
	}

	if query.ID != 0 {
		qs = qs.Filter("id", query.ID)
	}
	if query.Task != nil {
		qs = qs.Filter("task__icontains", *query.Task)
	}

	_, err := qs.All(&todos)
	if err != nil {
		logger.Log.WithError(err).Error(groupError)
		return todos, 0, err
	}

	count, err := qs.Count()
	if err != nil {
		logger.Log.WithError(err).Error(groupError)
		return todos, 0, err
	}
	return todos, count, nil
}

// NewTodoRepo initializes todoRepo
func NewTodoRepo(db orm.Ormer) TodoRepo {
	return &todoRepo{
		db: db,
	}
}
