package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/maxzhovtyj/financeApp-server/internal/models"
	"github.com/maxzhovtyj/financeApp-server/pkg/logger"
	"net/http"
)

type signInUserInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) initUsersRoutes(group *echo.Group) {
	users := group.Group("/users")
	{
		users.POST("/sign-up", h.signUp)
	}
}

func (h *Handler) signUp(ctx echo.Context) error {
	var input models.User

	if err := ctx.Bind(&input); err != nil {
		return err
	}

	id, err := h.service.Users.SignUp(input)
	if err != nil {
		logger.Error(err)
		return err
	}

	return ctx.String(http.StatusOK, id.String())
}

func (h *Handler) signIn(ctx echo.Context) error {
	var input signInUserInput

	if err := ctx.Bind(&input); err != nil {
		return err
	}

	accessToken, refreshToken, err := h.service.Users.SignIn(input.Email, input.Password)
	if err != nil {
		logger.Error(err)
		return err
	}

	return ctx.JSON(http.StatusOK, map[string]any{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	})
}
