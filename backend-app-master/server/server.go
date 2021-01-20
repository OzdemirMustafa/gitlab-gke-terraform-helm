package server

import (
	"backend-app/config"
	"fmt"
	"net/http"
	"time"

	"github.com/tylerb/graceful"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Server struct {
	E      *echo.Echo
	config *config.Config
}

func New(conf *config.Config) (*Server, error) {
	server := &Server{}

	e := echo.New()
	e.Use(middleware.CORS())

	server.E = e
	server.config = conf

	return server, nil
}

func (s *Server) Start() error {
	s.E.Server.Addr = fmt.Sprintf(":%d", s.config.Server.Port)
	s.E.GET("/health", s.healthControl)

	return graceful.ListenAndServe(s.E.Server, 10*time.Second)
}

func (s *Server) healthControl(context echo.Context) error {
	return context.NoContent(http.StatusOK)
}