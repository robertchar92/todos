package mysql

import (
	"log"
	"todo/models"
	"todo/utils/errors"

	"gorm.io/gorm"
)

func (r *ActivityRepository) FindByID(activityID uint64) (*models.Activity, error) {
	model := models.Activity{}
	err := r.db.Where("id = ?", activityID).First(&model).Error

	if err == gorm.ErrRecordNotFound {
		return nil, errors.ErrNotFound
	}

	if err != nil {
		log.Println("error-find-activity-by-id:", err)
		return nil, errors.CustomWrap(err, errors.ErrUnprocessableEntity, err.Error())
	}

	return &model, nil
}
