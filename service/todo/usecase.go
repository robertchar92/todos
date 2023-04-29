package todo

import (
	"todo/models"
	"todo/service/todo/delivery/http/request"
	"todo/utils/request_util"
	response_util "todo/utils/response_utils"
)

type Usecase interface {
	Index(paginationConfig request_util.PaginationConfig) ([]models.Todo, response_util.PaginationMeta, error)
	Show(todoID uint64) (*models.Todo, error)
	Create(request request.TodoCreateRequest) (*models.Todo, error)
	Update(todoID uint64, request request.TodoUpdateRequest) (*models.Todo, error)
	Delete(todoID uint64) error
}
