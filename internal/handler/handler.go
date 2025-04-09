package handler

import (
	"todo-app/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(srv *service.Service) *Handler {
	return &Handler{service: srv}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.POST("/tasks", h.Create)
	router.GET("/tasks", h.GetAll)
	router.PUT("/tasks/:id", h.Update)
	router.DELETE("/tasks/:id", h.Delete)

	return router
}
