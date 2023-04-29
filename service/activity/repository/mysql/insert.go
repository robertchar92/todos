package mysql

import (
	"log"
	"todo/models"
	"todo/utils/errors"

	"gorm.io/gorm"
)

func (r *ActivityRepository) Insert(activity *models.Activity, tx *gorm.DB) error {
	var db = r.db
	if tx != nil {
		db = tx
	}
	err := db.Create(activity).Error
	if err != nil {
		log.Println("error-insert-activity:", err)
		return errors.CustomWrap(err, errors.ErrUnprocessableEntity, err.Error())
	}
	return nil
}
