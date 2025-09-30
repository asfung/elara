package handlers

import (
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/asfung/elara/internal/entities"
	"github.com/asfung/elara/internal/models"
	"github.com/asfung/elara/internal/services"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type AuthHandler struct {
	*Handler
	authService services.AuthService
	userService services.UserService
	otpService  services.OTPService
}

func NewAuthHandler(authService services.AuthService, userService services.UserService, otpService services.OTPService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		userService: userService,
		otpService:  otpService,
	}
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

	validationErrors := h.ValidateBodyRequest(payload)
	if validationErrors != nil {
		return models.SendFailedValidationResponse(c, validationErrors)
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

	newRefreshCookie := &http.Cookie{
		Name:     "refresh_token",
		Value:    authResp.RefreshToken,
		HttpOnly: true,
		// Secure:   true,
		Secure:   false,
		SameSite: http.SameSiteStrictMode,
		// Path:     "/api/v1/auth/refresh",
		Path:    "/",
		Expires: time.Now().Add(30 * 24 * time.Hour),
	}
	c.SetCookie(newRefreshCookie)

	newAccessCookie := &http.Cookie{
		Name:     "access_token",
		Value:    authResp.RefreshToken,
		HttpOnly: true,
		// Secure:   true,
		Secure:   false,
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
		Expires:  time.Now().Add(15 * time.Minute), // 15 minutes
	}
	c.SetCookie(newAccessCookie)

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

	accessTokenCookie := &http.Cookie{
		Name:     "refresh_token",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		Expires:  time.Unix(0, 0), // expired date
		MaxAge:   -1,
	}
	c.SetCookie(accessTokenCookie)

	refreshTokenCookies := &http.Cookie{
		Name:     "access_token",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		Expires:  time.Unix(0, 0), // expired date
		MaxAge:   -1,
	}
	c.SetCookie(refreshTokenCookies)

	return c.JSON(http.StatusOK, map[string]string{"message": "logout successful"})
}

func (h *AuthHandler) Authenticated(c echo.Context) error {
	user := c.Get("user").(*entities.User)
	return c.JSON(http.StatusOK, models.ToUserResponse(*user))
}

// ============================================================
// ========================= NEW AUTH =========================
// ============================================================

func (h *AuthHandler) CheckEmail(c echo.Context) error {
	type request struct {
		Email string `json:"email" validate:"required,email"`
	}

	req := new(request)
	if err := h.BindBodyRequest(c, req); err != nil {
		return models.SendBadRequestResponse(c, err.Error())
	}

	validateErrors := h.ValidateBodyRequest(req)
	if validateErrors != nil {
		return models.SendFailedValidationResponse(c, validateErrors)
	}

	user, err := h.authService.GetUserByEmail(req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			res := map[string]interface{}{
				"continue_url": os.Getenv("CLIENT_URL") + "/create-account/password",
				"email":        req.Email,
			}
			return c.JSON(http.StatusOK, res)
		}
		return models.SendInternalServerErrorResponse(c, err.Error())
	}
	if user.UserID != uuid.Nil {
		res := map[string]interface{}{
			"continue_url": os.Getenv("CLIENT_URL") + "/log-in/password",
			"email":        req.Email,
		}
		return c.JSON(http.StatusOK, res)
	}

	// idk what de fuck its gonna be
	return c.JSON(http.StatusOK, nil)
}

func (h *AuthHandler) CreateAccount(c echo.Context) error {
	type request struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=8"`
	}
	req := new(request)
	if err := h.BindBodyRequest(c, req); err != nil {
		return models.SendBadRequestResponse(c, err.Error())
	}

	validateErrors := h.ValidateBodyRequest(req)
	if validateErrors != nil {
		return models.SendFailedValidationResponse(c, validateErrors)
	}

	userExist, _ := h.authService.GetUserByEmail(req.Email)
	if userExist.UserID != uuid.Nil {
		res := map[string]interface{}{
			"continue_url": os.Getenv("CLIENT_URL") + "/log-in/password",
		}
		return c.JSON(http.StatusOK, res)
	}

	user, err := h.authService.CreateAccountWithPassword(req.Email, req.Password)
	if err != nil {
		return models.SendInternalServerErrorResponse(c, err.Error())
	}

	data := map[string]interface{}{
		"continue_url": os.Getenv("CLIENT_URL") + "/email-verification",
		"user_id":      user.UserID,
	}
	// make it redirect to the http://localhost:6060/email-verification
	// location := os.Getenv("CLIENT_URL") + "/email-verification"
	// c.Redirect(http.StatusFound, location)

	return models.SendResponse(c, true, "acccount created, please verify OTP", data, http.StatusCreated)
}

func (h *AuthHandler) VerifyPassword(c echo.Context) error {
	type request struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=8"`
	}
	req := new(request)
	if err := h.BindBodyRequest(c, req); err != nil {
		return models.SendBadRequestResponse(c, err.Error())
	}

	validateErrors := h.ValidateBodyRequest(req)
	if validateErrors != nil {
		return models.SendFailedValidationResponse(c, validateErrors)
	}

	isValid, err := h.authService.VerifyPassword(req.Email, req.Password)
	if err != nil {
		return models.SendUnauthorizedResponse(c, err.Error())
	}

	if isValid {
		user, _ := h.authService.GetUserByEmail(req.Email)
		res := map[string]interface{}{
			"continue_url": os.Getenv("CLIENT_URL") + "/email-verification",
			"user_id":      user.UserID.String(),
		}
		return c.JSON(http.StatusOK, res)
	}

	return models.SendResponse(c, false, "Login Failed", nil, http.StatusUnauthorized)
}

func (h *AuthHandler) VerifyOTP(c echo.Context) error {
	type request struct {
		UserID string `json:"user_id" validate:"required"`
		OTP    string `json:"otp" validate:"required"`
	}
	req := new(request)
	if err := h.BindBodyRequest(c, req); err != nil {
		return models.SendBadRequestResponse(c, err.Error())
	}

	validateErrors := h.ValidateBodyRequest(req)
	if validateErrors != nil {
		return models.SendFailedValidationResponse(c, validateErrors)
	}

	isValid, err := h.otpService.VerifyOTP(req.UserID, req.OTP)
	if err != nil {
		return models.SendResponse(c, false, err.Error(), nil, http.StatusBadRequest)
	}

	// make a response isValid
	if isValid {
		user, _ := h.userService.GetUserByUserId(req.UserID)
		accessToken, refreshToken, err := h.authService.CreateTokensForUser(user)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
		}
		refreshCookie := &http.Cookie{
			Name:     "refresh_token",
			Value:    refreshToken,
			HttpOnly: true,
			// Secure:   true,
			Secure:   false,
			SameSite: http.SameSiteStrictMode,
			// Path:     "/api/v1/auth/refresh",
			Path:    "/",
			Expires: time.Now().Add(30 * 24 * time.Hour), // 30 days
		}
		c.SetCookie(refreshCookie)

		accessCookie := &http.Cookie{
			Name:     "access_token",
			Value:    accessToken,
			HttpOnly: true,
			// Secure:   true,
			Secure:   false,
			SameSite: http.SameSiteStrictMode,
			Path:     "/",
			Expires:  time.Now().Add(15 * time.Minute), // 15 minutes
		}
		c.SetCookie(accessCookie)

		data := models.AuthResponse{
			AccessToken:          accessToken,
			AccessTokenFormatted: "Bearer " + accessToken,
			ExpiresAt:            (24 * time.Hour * 7),
		}
		return models.SendSuccessResponse(c, "account verified, you can log in now", data)
	}

	return models.SendErrorResponse(c, "account not verified", http.StatusUnauthorized)
}
