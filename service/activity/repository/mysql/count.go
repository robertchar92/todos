package mysql

import (
	"log"
	"todo/models"
	"todo/utils/errors"
	"todo/utils/request_util"
)

func (r *ActivityRepository) Count(config request_util.PaginationConfig) (int64, error) {
	var count int64

	err := r.db.
		Model(&models.Activity{}).
		Scopes(config.Scopes()...).
		Count(&count).Error
	if err != nil {
		log.Println("error-count-authentication:", err)
		return 0, errors.CustomWrap(err, errors.ErrUnprocessableEntity, err.Error())
	}

	return count, nil
}
