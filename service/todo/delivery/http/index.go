package http

import (
	"fmt"
	"net/http"
	"todo/service/todo/delivery/http/request"
	response_util "todo/utils/response_utils"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Index(c *gin.Context) {
	todos, todoPagination, err := h.todoUsecase.Index(request.NewTodoPaginationConfig(c.Request.URL.Query()))
	if err != nil {
		c.JSON(http.StatusNotFound, response_util.ErrorResponse{
			Status:  http.StatusText(http.StatusNotFound),
			Message: fmt.Sprint(err),
		})
		return
	}

	c.JSON(http.StatusOK, response_util.IndexResponse{
		Status:  "Success",
		Message: "Success",
		Data:    todos,
		Meta:    todoPagination,
	})
}
