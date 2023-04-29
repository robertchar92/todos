package mysql

import (
	"log"
	"todo/models"
	"todo/utils/errors"

	"gorm.io/gorm"
)

func (r *TodoRepository) FindByID(todoID uint64) (*models.Todo, error) {
	model := models.Todo{}
	err := r.db.Where("id = ?", todoID).First(&model).Error

	if err == gorm.ErrRecordNotFound {
		return nil, errors.ErrNotFound
	}

	if err != nil {
		log.Println("error-find-todo-by-id:", err)
		return nil, errors.CustomWrap(err, errors.ErrUnprocessableEntity, err.Error())
	}

	return &model, nil
}
