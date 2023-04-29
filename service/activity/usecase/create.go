package usecase

import (
	"todo/models"
	"todo/service/activity/delivery/http/request"
	"todo/utils/errors"
)

func (u *ActivityUsecase) Create(request request.ActivityCreateRequest) (*models.Activity, error) {
	activityM := &models.Activity{
		Title: request.Title,
		Email: request.Email,
	}

	err := u.activityRepo.Insert(activityM, nil)
	if err != nil {
		err := errors.ErrUnprocessableEntity
		err.Message = "Error inserting activity."

		return nil, err
	}

	return activityM, nil
}
