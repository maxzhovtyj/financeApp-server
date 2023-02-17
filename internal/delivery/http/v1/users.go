package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/maxzhovtyj/financeApp-server/internal/models"
	"github.com/maxzhovtyj/financeApp-server/pkg/logger"
	"net/http"
)

type signInUserInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) initUsersRoutes(group *echo.Group) {
	users := group.Group("/users")
	{
		users.POST("/sign-up", h.signUp)
		users.POST("/sign-in", h.signIn)
	}
}

func (h *Handler) signUp(ctx echo.Context) error {
	var input models.User

	if err := ctx.Bind(&input); err != nil {
		logger.Error(err)
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return err
	}

	err := ctx.Validate(&input)
	if err != nil {
		return err
	}

	id, err := h.service.Users.SignUp(ctx.Request().Context(), input)
	if err != nil {
		logger.Error(err)
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return err
	}

	return ctx.String(http.StatusOK, id.String())
}

func (h *Handler) signIn(ctx echo.Context) error {
	var input signInUserInput

	if err := ctx.Bind(&input); err != nil {
		logger.Error(err)
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return err
	}

	accessToken, refreshToken, err := h.service.Users.SignIn(ctx.Request().Context(), input.Email, input.Password)
	if err != nil {
		logger.Error(err)
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return err
	}

	return ctx.JSON(http.StatusOK, map[string]any{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	})
}
