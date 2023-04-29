package http

import (
	"todo/app/middleware"
	"todo/service/activity"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	activityUsecase activity.Usecase
}

func New(activityUC activity.Usecase) *Handler {
	return &Handler{
		activityUsecase: activityUC,
	}
}

func (h *Handler) Register(r *gin.Engine, m *middleware.Middleware) {
	route := r.Group("/activity-groups")
	{
		route.GET("", h.Index)
		route.GET("/:id", h.Show)
		route.POST("/", h.Create)
		route.PATCH("/:id", h.Update)
		route.DELETE("/:id", h.Delete)
	}

}
