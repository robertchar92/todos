package usecase

import (
	"todo/models"
	"todo/utils/errors"
	"todo/utils/request_util"
	response_util "todo/utils/response_utils"
)

func (u *ActivityUsecase) Index(paginationConfig request_util.PaginationConfig) ([]models.Activity, response_util.PaginationMeta, error) {
	meta := response_util.PaginationMeta{
		Offset: paginationConfig.Offset(),
		Limit:  paginationConfig.Limit(),
		Total:  0,
	}

	activity, err := u.activityRepo.FindAll(paginationConfig)
	if err != nil || len(activity) == 0 {

		err := errors.ErrUnprocessableEntity
		err.Message = "Activity not found!"

		return nil, meta, err
	}

	total, err := u.activityRepo.Count(paginationConfig)
	if err != nil {
		return nil, meta, err
	}

	meta.Total = total

	return activity, meta, nil
}
