package handlers

import (
	"net/http"
	"time"

	"github.com/asfung/elara/internal/entities"
	"github.com/asfung/elara/internal/models"
	"github.com/asfung/elara/internal/services"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	*Handler
	authService services.AuthService
}

func NewAuthHandler(authService services.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Login(c echo.Context) error {
	var req models.LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	accessToken, refreshToken, err := h.authService.Login(req)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
	}

	cookie := &http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		Path:     "/api/v1/auth/refresh",
		Expires:  time.Now().Add(30 * 24 * time.Hour), // 30 days
	}
	c.SetCookie(cookie)

	data := models.AuthResponse{
		AccessToken:          accessToken,
		AccessTokenFormatted: "Bearer " + accessToken,
		// RefreshToken: refreshToken,
		ExpiresAt: (24 * time.Hour * 7),
	}

	response := models.ApiResponse{
		Success: true,
		Data:    data,
	}
	return c.JSON(http.StatusOK, response)
}

func (h *AuthHandler) Register(c echo.Context) error {
	payload := new(models.AddUserRequest)

	if err := h.BindBodyRequest(c, payload); err != nil {
		return models.SendBadRequestResponse(c, err.Error())
	}

	validateErros := h.ValidateBodyRequest(payload)
	if validateErros != nil {
		return models.SendFailedValidationResponse(c, validateErros)
	}

	user, err := h.authService.Register(*payload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, user)
}

func (h *AuthHandler) RefreshToken(c echo.Context) error {
	cookie, err := c.Cookie("refresh_token")
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "missing refresh token"})
	}

	refreshToken := cookie.Value

	authResp, err := h.authService.RefreshToken(models.RefreshTokenRequest{RefreshToken: refreshToken})
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
	}

	newCookie := &http.Cookie{
		Name:     "refresh_token",
		Value:    authResp.RefreshToken,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		Path:     "/api/v1/auth/refresh",
		Expires:  time.Now().Add(30 * 24 * time.Hour),
	}
	c.SetCookie(newCookie)

	return c.JSON(http.StatusOK, models.AuthResponse{
		AccessToken:          authResp.AccessToken,
		AccessTokenFormatted: "Bearer " + authResp.AccessToken,
		ExpiresAt:            15 * time.Minute,
	})
}

func (h *AuthHandler) Logout(c echo.Context) error {
	cookie, err := c.Cookie("refresh_token")
	if err != nil {
		return models.SendUnauthorizedResponse(c, "no refresh token cookie")
	}

	refreshToken := cookie.Value
	_, err = h.authService.Verify(refreshToken)
	if err != nil {
		return models.SendUnauthorizedResponse(c, "invalid refresh token revoked")
	}

	if err := h.authService.Logout(refreshToken); err != nil {
		return models.SendUnauthorizedResponse(c, err.Error())
	}

	expiredCookie := &http.Cookie{
		Name:     "refresh_token",
		Value:    "",
		Path:     "/api/v1/auth/refresh",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		Expires:  time.Unix(0, 0), // expired date
		MaxAge:   -1,
	}
	c.SetCookie(expiredCookie)

	return c.JSON(http.StatusOK, map[string]string{"message": "logout successful"})
}

func (h *AuthHandler) Authenticated(c echo.Context) error {
	user := c.Get("user").(*entities.User)
	return c.JSON(http.StatusOK, models.ToUserResponse(*user))
}
