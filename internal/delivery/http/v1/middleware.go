package v1

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/maxzhovtyj/financeApp-server/internal/models"
	"net/http"
	"strings"
)

func (h *Handler) userIdentity(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		authHeader := ctx.Request().Header.Get("Authorization")

		if authHeader == "" {
			return newErrorResponse(ctx, http.StatusUnauthorized, models.ErrInvalidAuthorizationHeader)
		}

		authHeaderParts := strings.Split(authHeader, " ")

		if len(authHeaderParts) != 2 {
			return newErrorResponse(ctx, http.StatusUnauthorized, models.ErrInvalidAuthorizationHeader)
		}

		if !strings.EqualFold(authHeaderParts[0], "Bearer") {
			return newErrorResponse(ctx, http.StatusUnauthorized, models.ErrInvalidAuthorizationHeader)
		}

		if authHeaderParts[1] == "" {
			return newErrorResponse(ctx, http.StatusUnauthorized, models.ErrInvalidAuthorizationHeader)
		}

		userId, err := h.tokenManager.Parse(authHeaderParts[1])
		if err != nil {
			return newErrorResponse(ctx, http.StatusUnauthorized, models.ErrInvalidAuthorizationHeader)
		}

		ctx.Set(userIdCtx, userId)

		return next(ctx)
	}
}

func getUserIdFromContext(ctx echo.Context) (string, error) {
	userIdString, ok := ctx.Get(userIdCtx).(string)
	if !ok {
		return "", fmt.Errorf("failed to cast user id to string")
	}

	return userIdString, nil
}
