package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/maxzhovtyj/financeApp-server/internal/service"
	"github.com/maxzhovtyj/financeApp-server/pkg/auth"
	"net/http"
)

const userIdCtx = "userId"

type Handler struct {
	service      *service.Service
	tokenManager auth.TokenManager
}

func NewHandler(service *service.Service, tokenManager auth.TokenManager) *Handler {
	return &Handler{
		service:      service,
		tokenManager: tokenManager,
	}
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
