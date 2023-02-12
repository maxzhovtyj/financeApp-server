package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	if err := e.Start(":8000"); err != nil {
		return
	}
}
