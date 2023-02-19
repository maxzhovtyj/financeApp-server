package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/maxzhovtyj/financeApp-server/pkg/logger"
)

type Error struct {
	Message string `json:"message"`
}

func newErrorResponse(ctx echo.Context, code int, err error) error {
	ctx.Response().Status = code

	_, responseErr := ctx.Response().Write([]byte(err.Error()))
	if responseErr != nil {
		return err
	}

	// Logging error
	logger.Error(err)

	return err
}
