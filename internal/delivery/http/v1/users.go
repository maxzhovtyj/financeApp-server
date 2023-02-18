package v1

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/maxzhovtyj/financeApp-server/internal/models"
	"github.com/maxzhovtyj/financeApp-server/pkg/logger"
	"net/http"
)

const signUpUrl = "/sign-up"

type signInUserInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) initUsersRoutes(group *echo.Group) {
	users := group.Group("/users")
	{
		users.POST(signUpUrl, h.signUp)
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

	err = h.service.Users.SignUp(ctx.Request().Context(), input)
	if err != nil {
		logger.Error(err)

		if errors.Is(err, models.ErrUserAlreadyExists) {
			newErrorResponse(ctx, http.StatusBadRequest, err.Error())
			return err
		}

		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return err
	}

	return ctx.NoContent(http.StatusCreated)
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
		if errors.Is(err, models.ErrUserNotFound) {
			newErrorResponse(ctx, http.StatusBadRequest, err.Error())
			return err
		}
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return err
	}

	return ctx.JSON(http.StatusOK, map[string]any{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	})
}
