package usecase

import (
	"fmt"
	"todo/utils/errors"
)

func (u *TodoUsecase) Delete(todoID uint64) error {
	todoM, err := u.todoRepo.FindByID(todoID)
	if err != nil {
		err := errors.ErrUnprocessableEntity
		err.Message = fmt.Sprintf("Todo with ID %d not found", todoID)

		return err
	}

	err = u.todoRepo.Delete(todoM, nil)
	if err != nil {
		err := errors.ErrUnprocessableEntity
		err.Message = fmt.Sprintf("Failed to delete Todo with ID %d", todoID)

		return err
	}

	return nil
}
