package activity

import (
	"todo/models"
	"todo/utils/request_util"

	"gorm.io/gorm"
)

type Repository interface {
	Insert(activity *models.Activity, tx *gorm.DB) error
	Update(activity *models.Activity, tx *gorm.DB) error
	Delete(activity *models.Activity, tx *gorm.DB) error
	FindByID(activityID uint64) (*models.Activity, error)
	FindAll(config request_util.PaginationConfig) ([]models.Activity, error)
	Count(config request_util.PaginationConfig) (int64, error)
}
