package v1

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/maxzhovtyj/financeApp-server/internal/models"
	"net/http"
)

const usersSignUpUrl = "/sign-up"
const usersSignInUrl = "/sign-in"

type signInUserInput struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (h *Handler) initUsersRoutes(group *echo.Group) {
	users := group.Group("/users")
	{
		users.POST(usersSignUpUrl, h.signUp)
		users.POST(usersSignInUrl, h.signIn)
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
