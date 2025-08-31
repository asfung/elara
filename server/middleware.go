package server

import (
	"errors"
	"net/http"
	"strings"
	"sync"

	"github.com/asfung/elara/internal/services"
	"github.com/charmbracelet/log"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// TEMPORARY
type logWriter struct{}

func (lw *logWriter) Write(p []byte) (n int, err error) {
	msg := string(p)

	if strings.Contains(msg, "status=5") {
		log.Error(msg)
	} else if strings.Contains(msg, "status=4") {
		log.Warn(msg)
	} else {
		log.Info(msg)
	}
	return len(p), nil
}

func BaseMiddleware(e *echo.Echo) {
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(RequestLoggerMiddleware)

	// e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins: []string{"*"},
	// 	AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodOptions},
	// 	AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderXRequestedWith, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	// }))
}

func AuthMiddleware(authService services.AuthService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		var doOnce sync.Once
		var registeredRoutes []*echo.Route

		return func(c echo.Context) error {
			doOnce.Do(func() {
				registeredRoutes = c.Echo().Routes()
			})

			// well, im tryna use https://github.com/asfung/TServer/blob/main/app/Http/Middleware/AccessControl.php
			// found this, https://github.com/labstack/echo/discussions/2081. not butter that i expected
			for _, r := range registeredRoutes {
				if r.Method == c.Request().Method && r.Path == c.Path() {
					if r.Name == "auth.refresh.token" {
						return next(c)
					}
					break
				}
			}

			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
				return echo.NewHTTPError(http.StatusUnauthorized, "Missing or invalid token")
			}

			token := strings.TrimPrefix(authHeader, "Bearer ")

			user, err := authService.Verify(token)
			if err != nil {
				// return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
				if errors.Is(err, jwt.ErrTokenExpired) {
					return c.JSON(http.StatusUnauthorized, map[string]interface{}{
						"message": "token is expired",
						"key":     "refresh-token",
					})
				}
				if errors.Is(err, jwt.ErrTokenMalformed) {
					return echo.NewHTTPError(http.StatusUnauthorized, "token is malformed")
				}
				if errors.Is(err, jwt.ErrTokenSignatureInvalid) {
					return echo.NewHTTPError(http.StatusUnauthorized, "invalid token signature")
				}
				if errors.Is(err, jwt.ErrTokenNotValidYet) {
					return echo.NewHTTPError(http.StatusUnauthorized, "token not valid yet")
				}
			}

			c.Set("user", user)
			return next(c)
		}
	}
}

func RequestLoggerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)

		status := c.Response().Status
		method := c.Request().Method
		uri := c.Request().RequestURI
		ip := c.RealIP()

		switch {
		case status >= 500:
			log.Error("request failed",
				"ip", ip,
				"method", method,
				"uri", uri,
				"status", status,
			)
		case status >= 400:
			log.Warn("client error",
				"ip", ip,
				"method", method,
				"uri", uri,
				"status", status,
			)
		default:
			log.Info("request success",
				"ip", ip,
				"method", method,
				"uri", uri,
				"status", status,
			)
		}

		return err
	}
}
