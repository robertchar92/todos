package mysql

import (
	"todo/service/activity"

	"gorm.io/gorm"
)

type ActivityRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) activity.Repository {
	return &ActivityRepository{
		db: db,
	}
}
