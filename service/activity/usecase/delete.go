package usecase

import (
	"fmt"
	"todo/models"
	"todo/utils/errors"
)

func (u *ActivityUsecase) Delete(activityID uint64) error {
	err := u.activityRepo.Delete(&models.Activity{ID: activityID}, nil)
	if err != nil {
		err := errors.ErrUnprocessableEntity
		err.Message = fmt.Sprintf("Activity Group with ID %d Not Found", activityID)

		return err
	}

	return nil
}
