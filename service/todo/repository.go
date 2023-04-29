package todo

import (
	"todo/models"
	"todo/utils/request_util"

	"gorm.io/gorm"
)

type Repository interface {
	Insert(todo *models.Todo, tx *gorm.DB) error
	Update(todo *models.Todo, tx *gorm.DB) error
	Delete(todo *models.Todo, tx *gorm.DB) error
	FindByID(todoID uint64) (*models.Todo, error)
	FindAll(config request_util.PaginationConfig) ([]models.Todo, error)
	Count(config request_util.PaginationConfig) (int64, error)
}
