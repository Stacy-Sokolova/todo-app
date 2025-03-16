package app

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	server *fiber.App
}

func (s *Server) Run(port string, srv *fiber.App) error {
	s = &Server{srv}
	return s.server.Listen(":" + port)
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.server.Shutdown()
}
