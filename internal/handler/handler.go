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

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.SignUp)
		auth.POST("/sign-in", h.SignIn)
	}

	api := router.Group("/tasks", h.UserIdentity)
	{
		api.POST("/", h.Create)
		api.GET("/", h.GetAll)
		api.PUT("/:id", h.Update)
		api.DELETE("/:id", h.Delete)
	}

	return router
}
