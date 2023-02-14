package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/maxzhovtyj/financeApp-server/internal/service"
	"net/http"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Init(api *echo.Group) {
	v1 := api.Group("/v1")
	{
		v1.GET("/ping", func(ctx echo.Context) error {
			err := ctx.String(http.StatusOK, "pong")
			if err != nil {
				return err
			}

			return nil
		})

		h.initUsersRoutes(v1)
	}
}
