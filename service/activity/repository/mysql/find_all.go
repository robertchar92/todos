package mysql

import (
	"log"
	"todo/models"
	"todo/utils/errors"
	"todo/utils/request_util"
)

func (r *ActivityRepository) FindAll(config request_util.PaginationConfig) ([]models.Activity, error) {
	results := make([]models.Activity, 0)

	err := r.db.
		Scopes(config.MetaScopes()...).
		Scopes(config.Scopes()...).
		Find(&results).Error
	if err != nil {
		log.Println("error-find-activity:", err)
		return nil, errors.CustomWrap(err, errors.ErrUnprocessableEntity, err.Error())
	}

	return results, nil
}
