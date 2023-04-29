package usecase

import (
	"fmt"
	"todo/models"
	"todo/service/todo/delivery/http/request"
	"todo/utils/errors"

	"github.com/jinzhu/copier"
)

func (u *TodoUsecase) Update(todoID uint64, request request.TodoUpdateRequest) (*models.Todo, error) {
	todoM, err := u.todoRepo.FindByID(todoID)
	if err != nil {
		err := errors.ErrUnprocessableEntity
		err.Message = fmt.Sprintf("Activity Group with ID %d Not Found", todoID)

		return nil, err
	}

	_ = copier.Copy(todoM, &request)

	if todoM.Priority != models.TodoPriorityVeryHigh &&
		todoM.Priority != models.TodoPriorityHigh &&
		todoM.Priority != models.TodoPriorityMedium &&
		todoM.Priority != models.TodoPriorityLow &&
		todoM.Priority != models.TodoPriorityVeryLow {
		todoM.Priority = models.TodoPriorityVeryHigh
	}

	err = u.todoRepo.Update(todoM, nil)
	if err != nil {
		err := errors.ErrUnprocessableEntity
		err.Message = "Error Updating Todo."

		return nil, err
	}

	return todoM, nil
}
