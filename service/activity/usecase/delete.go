package usecase

import (
	"fmt"
	"todo/utils/errors"
)

func (u *ActivityUsecase) Delete(activityID uint64) error {
	activityM, err := u.activityRepo.FindByID(activityID)
	if err != nil {
		err := errors.ErrUnprocessableEntity
		err.Message = fmt.Sprintf("Activity with ID %d Not Found", activityID)

		return err
	}

	err = u.activityRepo.Delete(activityM, nil)
	if err != nil {
		err := errors.ErrUnprocessableEntity
		err.Message = fmt.Sprintf("Failed to delete Activity Group with ID %d", activityID)

		return err
	}

	return nil
}
