package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HelloHandler interface {
	Helllo(c echo.Context) error
}

type helloHandler struct {
}

func NewHelloHandler() HelloHandler {
	return &helloHandler{}
}

func (h *helloHandler) Helllo(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, %s!"+c.Param("name"))
}
