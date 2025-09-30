package handlers

import (
	"fmt"

	"github.com/asfung/elara/internal/models"
	"github.com/asfung/elara/internal/services"
	"github.com/labstack/echo/v4"
)

type WalletHandler struct {
	*Handler
	walletService services.WalletService
}

func NewWalletHandler(walletService services.WalletService) *WalletHandler {
	return &WalletHandler{walletService: walletService}
}

func (h *WalletHandler) CreateWallet(c echo.Context) error {
	payload := new(models.AddWalletRequest)
	if err := h.BindBodyRequest(c, payload); err != nil {
		return models.SendBadRequestResponse(c, err.Error())
	}

	validationErrors := h.ValidateBodyRequest(payload)
	if validationErrors != nil {
		return models.SendFailedValidationResponse(c, validationErrors)
	}

	wallet, err := h.walletService.CreateWallet(*payload)
	if err != nil {
		return models.SendInternalServerErrorResponse(c, err.Error())
	}
	return models.SendSuccessResponse(c, "berhasil membuat wallet baru", models.ToWalletResponse(wallet))
}

func (h *WalletHandler) UpdateWallet(c echo.Context) error {
	payload := new(models.UpdateWalletRequest)
	if err := h.BindBodyRequest(c, payload); err != nil {
		return models.SendBadRequestResponse(c, err.Error())
	}

	validationErrors := h.ValidateBodyRequest(payload)
	if validationErrors != nil {
		return models.SendFailedValidationResponse(c, validationErrors)
	}

	wallet, err := h.walletService.UpdateWallet(*payload)
	if err != nil {
		return models.SendInternalServerErrorResponse(c, err.Error())
	}

	return models.SendSuccessResponse(c, "berhasil update wallet", models.ToWalletResponse(wallet))
}

func (h *WalletHandler) GetWalletById(c echo.Context) error {
	id := c.Param("id")
	wallet, err := h.walletService.GetWalletById(id)
	if err != nil {
		return models.SendInternalServerErrorResponse(c, err.Error())
	}
	return models.SendSuccessResponse(c, "berhasil mengambil wallet", models.ToWalletResponse(wallet))
}

func (h *WalletHandler) DeleteWallet(c echo.Context) error {
	id := c.Param("id")
	if err := h.walletService.DeleteWallet(id); err != nil {
		return models.SendInternalServerErrorResponse(c, err.Error())
	}
	return models.SendSuccessResponse(c, "berhasil menghapus wallet", nil)
}

func (h *WalletHandler) GetWalletByUserId(c echo.Context) error {
	userId := c.Param("userId")
	wallet, err := h.walletService.GetWalletByUserId(userId)
	if err != nil {
		return models.SendInternalServerErrorResponse(c, err.Error())
	}
	return models.SendSuccessResponse(c, "berhasil mengambil wallet", models.ToWalletResponse(wallet))
}

func (h *WalletHandler) UpdateWalletBalance(c echo.Context) error {
	id := c.Param("id")
	var isCredit bool
	if isCreditStr := c.QueryParam("is_credit"); isCreditStr != "" {
		fmt.Sscanf(isCreditStr, "%t", &isCredit)
	}
	payload := new(models.UpdateWalletRequest)
	if err := h.BindBodyRequest(c, payload); err != nil {
		return models.SendBadRequestResponse(c, err.Error())
	}

	validationErrors := h.ValidateBodyRequest(payload)
	if validationErrors != nil {
		return models.SendFailedValidationResponse(c, validationErrors)
	}

	_, err := h.walletService.UpdateWalletBalance(id, payload.Balance, isCredit)
	if err != nil {
		return models.SendInternalServerErrorResponse(c, err.Error())
	}

	return models.SendSuccessResponse(c, "berhasil update wallet balance", nil)
}
