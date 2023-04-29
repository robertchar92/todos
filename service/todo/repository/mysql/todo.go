package mysql

import (
	"todo/service/todo"

	"gorm.io/gorm"
)

type TodoRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) todo.Repository {
	return &TodoRepository{
		db: db,
	}
}
