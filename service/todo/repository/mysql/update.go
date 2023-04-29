package mysql

import (
	"log"
	"todo/models"
	"todo/utils/errors"

	"gorm.io/gorm"
)

func (r *TodoRepository) Update(todo *models.Todo, tx *gorm.DB) error {
	var db = r.db
	if tx != nil {
		db = tx
	}
	err := db.Save(todo).Error
	if err != nil {
		log.Println("error-update-todo:", err)
		return errors.CustomWrap(err, errors.ErrUnprocessableEntity, err.Error())
	}
	return nil
}
