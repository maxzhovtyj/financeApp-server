package delivery

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/maxzhovtyj/financeApp-server/internal/service"
)

type Handler struct {
	service *service.Service
}

func New(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Init() *echo.Echo {
	router := echo.New()

	router.Use(middleware.Recover())

	return router
}
