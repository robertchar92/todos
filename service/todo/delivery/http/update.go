package http

import (
	"fmt"
	"net/http"
	"strconv"
	"todo/service/todo/delivery/http/request"
	response_util "todo/utils/response_utils"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Update(c *gin.Context) {
	todoID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, response_util.ErrorResponse{
			Status:  http.StatusText(http.StatusBadRequest),
			Message: "title cannot be null",
		})
		return
	}

	var req request.TodoUpdateRequest
	if err := c.ShouldBind(&req); err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypeBind)
		return
	}

	activitieM, err := h.todoUsecase.Update(todoID, req)
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
