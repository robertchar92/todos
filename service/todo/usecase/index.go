package usecase

import (
	"todo/models"
	"todo/utils/errors"
	"todo/utils/request_util"
	response_util "todo/utils/response_utils"
)

func (u *TodoUsecase) Index(paginationConfig request_util.PaginationConfig) ([]models.Todo, response_util.PaginationMeta, error) {
	meta := response_util.PaginationMeta{
		Offset: paginationConfig.Offset(),
		Limit:  paginationConfig.Limit(),
		Total:  0,
	}

	todos, err := u.todoRepo.FindAll(paginationConfig)
	if err != nil {

		err := errors.ErrUnprocessableEntity
		err.Message = "Todo items not found!"

		return nil, meta, err
	}

	total, err := u.todoRepo.Count(paginationConfig)
	if err != nil {
		return nil, meta, err
	}

	meta.Total = total

	return todos, meta, nil
}
