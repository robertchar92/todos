package request

import "todo/utils/request_util"

type TodoCreateRequest struct {
	Title           string  `form:"title" json:"title" binding:"required"`
	ActivityGroupID uint64  `form:"activity_group_id" json:"activity_group_id" binding:"required"`
	Priority        *string `form:"priority" json:"priority" binding:"omitempty"`
}

type TodoUpdateRequest struct {
	Title    *string `form:"title" json:"title" binding:"omitempty"`
	Priority *string `form:"priority" json:"priority" binding:"omitempty"`
	IsActive *bool   `form:"is_active" json:"is_active" binding:"omitempty"`
}

func NewTodoPaginationConfig(conditions map[string][]string) request_util.PaginationConfig {
	filterable := map[string]string{
		"id":                request_util.IdType,
		"activity_group_id": request_util.IdType,
		"title":             request_util.StringType,
		"priority":          request_util.StringType,
		"is_active":         request_util.BoolType,
	}
	return request_util.NewRequestPaginationConfig(conditions, filterable)
}
