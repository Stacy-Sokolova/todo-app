package handler

import (
	"todo-app/internal/service"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service *service.Service
}

func NewHandler(srv *service.Service) *Handler {
	return &Handler{service: srv}
}

func (h *Handler) InitRoutes() *fiber.App {
	router := fiber.New()

	router.Post("/tasks", h.Create)
	router.Get("/tasks", h.GetAll)
	router.Put("/tasks/:id", h.Update)
	router.Delete("/tasks/:id", h.Delete)

	return router
}
