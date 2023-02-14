package server

import (
	"github.com/labstack/echo/v4"
	"github.com/maxzhovtyj/financeApp-server/internal/config"
)

type Server struct {
	echo    *echo.Echo
	address string
}

func NewServer(cfg *config.Config, e *echo.Echo) *Server {
	e.Server.MaxHeaderBytes = cfg.HTTP.MaxHeaderMegabytes << 20
	e.Server.ReadTimeout = cfg.HTTP.ReadTimeout
	e.Server.WriteTimeout = cfg.HTTP.WriteTimeout

	return &Server{
		echo:    e,
		address: ":" + cfg.HTTP.Port,
	}
}

func (s *Server) Run() error {
	return s.echo.Start(s.address)
}
