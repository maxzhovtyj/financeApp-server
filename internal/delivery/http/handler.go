package delivery

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	v1 "github.com/maxzhovtyj/financeApp-server/internal/delivery/http/v1"
	"github.com/maxzhovtyj/financeApp-server/internal/service"
	"github.com/maxzhovtyj/financeApp-server/pkg/auth"
)

type Handler struct {
	service      *service.Service
	validator    v1.AppValidator
	tokenManager auth.TokenManager
}

func New(service *service.Service, tokenManager auth.TokenManager) *Handler {
	return &Handler{
		service:      service,
		tokenManager: tokenManager,
	}
}

func (h *Handler) Init() *echo.Echo {
	router := echo.New()

	router.Use(middleware.Logger())
	router.Use(middleware.Recover())

	v := validator.New()
	router.Validator = &v1.AppValidator{Validator: v}

	h.initAPI(router)

	return router
}

func (h *Handler) initAPI(router *echo.Echo) {
	handlerV1 := v1.NewHandler(h.service, h.tokenManager)
	api := router.Group("/api")
	{
		handlerV1.Init(api)
	}
}
