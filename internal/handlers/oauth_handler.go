package handlers

import (
	"net/http"
	"time"

	"github.com/asfung/elara/internal/models"
	"github.com/asfung/elara/internal/services"
	"github.com/labstack/echo/v4"
	"github.com/markbates/goth/gothic"
)

type OAuthHandler struct {
	*Handler
	authService services.AuthService
	frontendURL string
}

func NewOAuthHandler(authService services.AuthService, frontendURL string) *OAuthHandler {
	return &OAuthHandler{
		authService: authService,
		frontendURL: frontendURL,
	}
}

func (h *OAuthHandler) BeginAuth(c echo.Context) error {
	provider := c.Param("provider")
	if provider == "" {
		return models.SendErrorResponse(c, "missing provider", http.StatusBadRequest)
	}

	q := c.Request().URL.Query()
	q.Set("provider", provider)
	c.Request().URL.RawQuery = q.Encode()

	gothic.BeginAuthHandler(c.Response(), c.Request())
	return nil
}

func (h *OAuthHandler) Callback(c echo.Context) error {
	provider := c.Param("provider")
	if provider == "" {
		return models.SendErrorResponse(c, "missing provider", http.StatusBadRequest)
	}

	q := c.Request().URL.Query()
	q.Set("provider", provider)
	c.Request().URL.RawQuery = q.Encode()

	gUser, err := gothic.CompleteUserAuth(c.Response(), c.Request())
	if err != nil {
		return models.SendUnauthorizedResponse(c, err.Error())
	}

	accessToken, refreshToken, err := h.authService.OAuthLoginFromGothUser(gUser)
	if err != nil {
		return models.SendInternalServerErrorResponse(c, err.Error())
	}

	cookie := &http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		Path:     "/api/v1/auth/refresh",
		Expires:  time.Now().Add(30 * 24 * time.Hour),
	}

	c.SetCookie(cookie)

	redirectUrl := h.frontendURL
	if redirectUrl == "" {
		redirectUrl = "http://localhost:6060"
	}

	location := redirectUrl + "/auth/oauth/callback#access_token=" + accessToken
	return c.Redirect(http.StatusFound, location)
}
