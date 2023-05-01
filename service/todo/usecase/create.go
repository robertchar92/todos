package usecase

import (
	"fmt"
	"todo/models"
	"todo/service/todo/delivery/http/request"
	"todo/utils/errors"
)

func (u *TodoUsecase) Create(request request.TodoCreateRequest) (*models.Todo, error) {
	activityM, _ := u.activityRepo.FindByID(request.ActivityGroupID)
	if activityM == nil {
		err := errors.ErrNotFound
		err.Message = fmt.Sprintf("Error create todo items for Activity Group ID %d.", request.ActivityGroupID)

		return nil, err
	}

	todoM := &models.Todo{
		Title:           request.Title,
		ActivityGroupID: request.ActivityGroupID,
	}

	if (request.Priority != nil && (*request.Priority != models.TodoPriorityVeryHigh &&
		*request.Priority != models.TodoPriorityHigh &&
		*request.Priority != models.TodoPriorityMedium &&
		*request.Priority != models.TodoPriorityLow &&
		*request.Priority != models.TodoPriorityVeryLow)) || request.Priority == nil {
		todoM.Priority = models.TodoPriorityVeryHigh
	}

	todoM.IsActive = true

	err := u.todoRepo.Insert(todoM, nil)
	if err != nil {
		err := errors.ErrUnprocessableEntity
		err.Message = "Error inserting todo."

		return nil, err
	}

	return todoM, nil
}
