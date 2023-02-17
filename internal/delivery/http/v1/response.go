package v1

import (
	"github.com/labstack/echo/v4"
)

type Error struct {
	Message string `json:"message"`
}

func newErrorResponse(ctx echo.Context, code int, message string) {
	e := ctx.JSON(code, Error{Message: message})
	if e != nil {
		return
	}
}
