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
	// route := r.Group("/todo")
	// {
	// route.GET("/:code", func(c *gin.Context) {
	// 	path := c.Param("code")
	// 	if path == "counter" {
	// 		h.UserCounter(c)
	// 	} else {
	// 		h.UserShow(c)
	// 	}
	// })
	// route.GET("/:code/example-needed-picture", h.UserExampleNeededPicture)
	// route.GET("", h.UserIndex)
	// route.POST("/folder-name", h.UserFolderPath)
	// route.POST("/upload-single-image", h.UploadSingleImage)
	// route.POST("/upload-single-add-image/:code", h.UploadSingleAddImage)
	// route.POST("/comment/:code", h.UserAddComment)
	// route.POST("/add-additional-picture/:code", h.UserAddAddAdditionalPicture)
	// }

}
