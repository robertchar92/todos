package http

import (
	"fmt"
	"net/http"
	"strconv"
	"todo/utils/errors"
	response_util "todo/utils/response_utils"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Show(c *gin.Context) {
	todoID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		_ = c.Error(errors.ErrUnprocessableEntity).SetType(gin.ErrorTypePublic)
		return
	}

	todos, err := h.todoUsecase.Show(todoID)
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
		Data:    todos,
	})
}
