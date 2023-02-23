package v1

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/maxzhovtyj/financeApp-server/internal/models"
	"net/http"
)

const (
	usersUrl       = "/users"
	signUpUrl      = "/sign-up"
	signInUrl      = "/sign-in"
	authRefreshUrl = "/auth/refresh"
)

func (h *Handler) initUsersRoutes(group *echo.Group) {
	users := group.Group(usersUrl)
	{
		users.POST(signUpUrl, h.signUp)
		users.POST(signInUrl, h.signIn)
		users.POST(authRefreshUrl, h.userRefresh)
	}
}

func (h *Handler) signUp(ctx echo.Context) error {
	var input models.User

	if err := BindAndValidate(ctx, &input); err != nil {
		return newErrorResponse(ctx, http.StatusBadRequest, models.ErrInvalidInputBody)
	}

	err := h.service.Users.SignUp(ctx.Request().Context(), input)
	if err != nil {
		if errors.Is(err, models.ErrUserAlreadyExists) {
			return newErrorResponse(ctx, http.StatusBadRequest, models.ErrUserAlreadyExists)
		}

		return newErrorResponse(ctx, http.StatusInternalServerError, err)
	}

	return ctx.NoContent(http.StatusCreated)
}

type signInUserInput struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (h *Handler) signIn(ctx echo.Context) error {
	var input signInUserInput

	if err := BindAndValidate(ctx, &input); err != nil {
		return newErrorResponse(ctx, http.StatusBadRequest, models.ErrInvalidInputBody)
	}

	accessToken, refreshToken, err := h.service.Users.SignIn(ctx.Request().Context(), input.Email, input.Password)
	if err != nil {
		if errors.Is(err, models.ErrUserNotFound) {
			return newErrorResponse(ctx, http.StatusBadRequest, models.ErrUserNotFound)
		}

		return newErrorResponse(ctx, http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, map[string]any{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	})
}

func (h *Handler) userRefresh(ctx echo.Context) error {
	// TODO
	return newErrorResponse(ctx, http.StatusNotImplemented, errors.New("not implemented"))
}
