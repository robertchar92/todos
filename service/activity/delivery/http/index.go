package http

import (
	"fmt"
	"net/http"
	"todo/service/activity/delivery/http/request"
	response_util "todo/utils/response_utils"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Index(c *gin.Context) {
	activities, activityPagination, err := h.activityUsecase.Index(request.NewActivityPaginationConfig(c.Request.URL.Query()))
	if err != nil {
		// _ = c.Error(err).SetType(gin.ErrorTypePublic)

		c.JSON(http.StatusNotFound, response_util.ErrorResponse{
			Status:  http.StatusText(http.StatusNotFound),
			Message: fmt.Sprint(err),
		})
		return
	}

	c.JSON(http.StatusOK, response_util.IndexResponse{
		Status:  "Success",
		Message: "Success",
		Data:    activities,
		Meta:    activityPagination,
	})
}
