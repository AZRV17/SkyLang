package http

import (
	"github.com/AZRV17/Skylang/internal/config"
	v1 "github.com/AZRV17/Skylang/internal/delivery/http/v1"
	"github.com/AZRV17/Skylang/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service service.Service
	config  *config.Config
}

func NewHandler(service service.Service, config *config.Config) *Handler {
	return &Handler{
		service: service,
		config:  config,
	}
}

func (h *Handler) Init(r *gin.Engine) {
	v1 := v1.NewHandler(h.service, h.config)
	v1.Init(r)
}
