package handler

import (
	"url-shortener/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.POST("/url", h.Create)
	router.GET("/:alias", h.Get)
	router.DELETE("/:alias", h.Delete)

	return router
}
