package mysql

import (
	"log"
	"todo/models"
	"todo/utils/errors"
	"todo/utils/request_util"
)

func (r *TodoRepository) FindAll(config request_util.PaginationConfig) ([]models.Todo, error) {
	results := make([]models.Todo, 0)

	err := r.db.
		Scopes(config.MetaScopes()...).
		Scopes(config.Scopes()...).
		Find(&results).Error
	if err != nil {
		log.Println("error-find-todo-items:", err)
		return nil, errors.CustomWrap(err, errors.ErrUnprocessableEntity, err.Error())
	}

	return results, nil
}
