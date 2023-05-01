package http

import (
	"todo/app/middleware"
	"todo/service/todo"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	todoUsecase todo.Usecase
}

func New(todoUC todo.Usecase) *Handler {
	return &Handler{
		todoUsecase: todoUC,
	}
}

func (h *Handler) Register(r *gin.Engine, m *middleware.Middleware) {
	route := r.Group("/todo-items")
	{
		route.GET("", h.Index)
		route.GET("/:id", h.Show)
		route.POST("", h.Create)
		route.PATCH("/:id", h.Update)
		route.DELETE("/:id", h.Delete)
	}

}
