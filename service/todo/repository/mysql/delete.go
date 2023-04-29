package mysql

import (
	"log"
	"todo/models"
	"todo/utils/errors"

	"gorm.io/gorm"
)

func (r *TodoRepository) Delete(todo *models.Todo, tx *gorm.DB) error {
	var db = r.db
	if tx != nil {
		db = tx
	}
	err := db.Delete(todo).Error
	if err != nil {
		log.Println("error-delete-todo:", err)
		return errors.CustomWrap(err, errors.ErrUnprocessableEntity, err.Error())
	}
	return nil
}
