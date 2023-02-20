package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/maxzhovtyj/financeApp-server/pkg/logger"
)

func BindAndValidate(ctx echo.Context, i interface{}) error {
	if err := ctx.Bind(i); err != nil {
		logger.Error(err)
		return err
	}

	if err := ctx.Validate(i); err != nil {
		logger.Error(err)
		return err
	}

	return nil
}
