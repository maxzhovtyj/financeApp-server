package v1

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AppValidator struct {
	Validator *validator.Validate
}

func (av *AppValidator) Validate(i interface{}) error {
	if err := av.Validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return nil
}
