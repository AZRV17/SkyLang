package v1

import (
	"github.com/AZRV17/Skylang/internal/config"
	"github.com/AZRV17/Skylang/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service service.Service
	config  *config.Config
}

func NewHandler(service service.Service, cfg *config.Config) *Handler {
	return &Handler{
		service: service,
		config:  cfg,
	}
}

func (h *Handler) Init(r *gin.Engine) {
	h.initUserRoutes(r)
	h.initExerciseRoutes(r)
	h.initLectureRoutes(r)
	h.initCourseRoutes(r)
	h.initCommentRoutes(r)
	h.initAuthorRequestRoutes(r)
}
