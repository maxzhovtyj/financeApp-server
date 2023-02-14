package delivery

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	v1 "github.com/maxzhovtyj/financeApp-server/internal/delivery/http/v1"
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

	h.initAPI(router)

	return router
}

func (h *Handler) initAPI(router *echo.Echo) {
	handlerV1 := v1.NewHandler(h.service)
	api := router.Group("/api")
	{
		handlerV1.Init(api)
	}
}
