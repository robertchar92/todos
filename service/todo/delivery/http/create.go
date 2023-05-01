package http

import (
	"fmt"
	"net/http"
	"todo/service/todo/delivery/http/request"
	response_util "todo/utils/response_utils"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Create(c *gin.Context) {
	var req request.TodoCreateRequest
	if err := c.ShouldBind(&req); err != nil {
		m := ""

		if req.Title == "" {
			m = "title cannot be null"
		}

		if req.ActivityGroupID == 0 {
			m = "activity_group_id cannot be null"
		}

		c.JSON(http.StatusBadRequest, response_util.ErrorResponse{
			Status:  http.StatusText(http.StatusBadRequest),
			Message: m,
		})
		return
	}

	todoM, err := h.todoUsecase.Create(req)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, response_util.ErrorResponse{
			Status:  http.StatusText(http.StatusUnprocessableEntity),
			Message: fmt.Sprint(err),
		})
		return
	}

	c.JSON(http.StatusCreated, response_util.ShowResponse{
		Status:  "Success",
		Message: "Success",
		Data:    todoM,
	})
}
