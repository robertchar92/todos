package http

import (
	"fmt"
	"net/http"
	"strconv"
	"todo/service/activity/delivery/http/request"
	response_util "todo/utils/response_utils"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Update(c *gin.Context) {
	activityID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, response_util.ErrorResponse{
			Status:  http.StatusText(http.StatusBadRequest),
			Message: "ID is not a valid",
		})
		return
	}

	var req request.ActivityUpdateRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, response_util.ErrorResponse{
			Status:  http.StatusText(http.StatusBadRequest),
			Message: "title cannot be null",
		})
		return
	}

	activitieM, err := h.activityUsecase.Update(activityID, req)
	if err != nil {
		c.JSON(http.StatusNotFound, response_util.ErrorResponse{
			Status:  http.StatusText(http.StatusNotFound),
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
