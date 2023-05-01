package usecase

import (
	"fmt"
	"todo/models"
	"todo/service/activity/delivery/http/request"
	"todo/utils/errors"

	"github.com/jinzhu/copier"
)

func (u *ActivityUsecase) Update(activityID uint64, request request.ActivityUpdateRequest) (*models.Activity, error) {
	activityM, err := u.activityRepo.FindByID(activityID)
	if err != nil {
		err := errors.ErrUnprocessableEntity
		err.Message = fmt.Sprintf("Activity with ID %d Not Found", activityID)

		return nil, err
	}

	if request.Email == nil {
		request.Email = &activityM.Email
	}

	_ = copier.Copy(activityM, &request)

	err = u.activityRepo.Update(activityM, nil)
	if err != nil {
		err := errors.ErrUnprocessableEntity
		err.Message = "Error Updating activity."

		return nil, err
	}

	return activityM, nil
}
