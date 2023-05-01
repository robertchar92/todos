package usecase

import (
	"fmt"
	"todo/models"
	"todo/utils/errors"
)

func (u *ActivityUsecase) Show(activityID uint64) (*models.Activity, error) {
	activityM, err := u.activityRepo.FindByID(activityID)
	if err != nil {
		err := errors.ErrUnprocessableEntity
		err.Message = fmt.Sprintf("Activity with ID %d Not Found", activityID)

		return nil, err
	}

	return activityM, nil
}
