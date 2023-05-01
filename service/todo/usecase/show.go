package usecase

import (
	"fmt"
	"todo/models"
	"todo/utils/errors"
)

func (u *TodoUsecase) Show(todoID uint64) (*models.Todo, error) {
	todoM, err := u.todoRepo.FindByID(todoID)
	if err != nil {
		err := errors.ErrUnprocessableEntity
		err.Message = fmt.Sprintf("Todo with ID %d Not Found", todoID)

		return nil, err
	}

	return todoM, nil
}
