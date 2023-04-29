package mysql

import (
	"log"
	"todo/models"
	"todo/utils/errors"

	"gorm.io/gorm"
)

func (r *TodoRepository) Insert(todo *models.Todo, tx *gorm.DB) error {
	var db = r.db
	if tx != nil {
		db = tx
	}
	err := db.Create(todo).Error
	if err != nil {
		log.Println("error-insert-todo:", err)
		return errors.CustomWrap(err, errors.ErrUnprocessableEntity, err.Error())
	}
	return nil
}
