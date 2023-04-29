package http

import (
	"fmt"
	"net/http"
	"todo/service/activity/delivery/http/request"
	response_util "todo/utils/response_utils"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Create(c *gin.Context) {
	var req request.ActivityCreateRequest
	if err := c.ShouldBind(&req); err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypeBind)
		return
	}

	activitieM, err := h.activityUsecase.Create(req)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, response_util.ErrorResponse{
			Status:  http.StatusText(http.StatusUnprocessableEntity),
			Message: fmt.Sprint(err),
		})
		return
	}

	c.JSON(http.StatusOK, response_util.ShowResponse{
		Status:  "Success",
		Message: "Success",
		Data:    activitieM,
	})
}
