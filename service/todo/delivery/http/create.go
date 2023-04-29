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
		c.JSON(http.StatusBadRequest, response_util.ErrorResponse{
			Status:  http.StatusText(http.StatusBadRequest),
			Message: "title cannot be null",
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

	c.JSON(http.StatusOK, response_util.ShowResponse{
		Status:  "Success",
		Message: "Success",
		Data:    todoM,
	})
}
