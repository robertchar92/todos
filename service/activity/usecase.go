package activity

import (
	"todo/models"
	"todo/service/activity/delivery/http/request"
	"todo/utils/request_util"
	response_util "todo/utils/response_utils"
)

type Usecase interface {
	Index(paginationConfig request_util.PaginationConfig) ([]models.Activity, response_util.PaginationMeta, error)
	Show(activityID uint64) (*models.Activity, error)
	Create(request request.ActivityCreateRequest) (*models.Activity, error)
	Update(activityID uint64, request request.ActivityUpdateRequest) (*models.Activity, error)
	Delete(activityID uint64) error
}
