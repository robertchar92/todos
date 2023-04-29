package request

import "todo/utils/request_util"

type ActivityCreateRequest struct {
	Title string `form:"title" json:"title" binding:"required"`
	Email string `form:"email" json:"email" binding:"required"`
}

type ActivityUpdateRequest struct {
	Title string  `form:"title" json:"title" binding:"required"`
	Email *string `form:"email" json:"email" binding:"omitempty"`
}

func NewActivityPaginationConfig(conditions map[string][]string) request_util.PaginationConfig {
	filterable := map[string]string{
		"id":    request_util.IdType,
		"title": request_util.StringType,
		"email": request_util.StringType,
	}
	return request_util.NewRequestPaginationConfig(conditions, filterable)
}
