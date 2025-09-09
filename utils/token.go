package utils

import (
	"errors"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func GetBearerToken(c echo.Context) (string, error) {
	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("missing authorization header")
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
		return "", errors.New("invalid authorization header format")
	}

	return parts[1], nil
}

func MustGetBearerToken(c echo.Context) (string, bool) {
	token, err := GetBearerToken(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		return "", false
	}
	return token, true
}
