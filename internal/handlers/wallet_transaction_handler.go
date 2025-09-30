package handlers

import (
	"fmt"

	"github.com/asfung/elara/internal/models"
	"github.com/asfung/elara/internal/services"
	"github.com/labstack/echo/v4"
)

type WalletTransactionHandler struct {
	*Handler
	walletTransactionService services.WalletTransactionService
}

func NewWalletTransactionHandler(walletTransactionService services.WalletTransactionService) *WalletTransactionHandler {
	return &WalletTransactionHandler{walletTransactionService: walletTransactionService}
}

func (h *WalletTransactionHandler) CreateWalletTransaction(c echo.Context) error {
	payload := new(models.AddWalletTransactionRequest)
	if err := h.BindBodyRequest(c, payload); err != nil {
		return models.SendBadRequestResponse(c, err.Error())
	}

	validationErrors := h.ValidateBodyRequest(payload)
	if validationErrors != nil {
		return models.SendFailedValidationResponse(c, validationErrors)
	}

	walletTransaction, wallet, err := h.walletTransactionService.CreateWalletTransaction(*payload)
	if err != nil {
		return models.SendInternalServerErrorResponse(c, err.Error())
	}

	return models.SendSuccessResponse(c, "berhasil mendapatkan wallet transaction", models.ToWalletTransactionWithWalletResponse(walletTransaction, &wallet))
}

func (h *WalletTransactionHandler) UpdateWalletTransaction(c echo.Context) error {
	payload := new(models.UpdateWalletTransactionRequest)
	if err := h.BindBodyRequest(c, payload); err != nil {
		return models.SendBadRequestResponse(c, err.Error())
	}

	validationErrors := h.ValidateBodyRequest(payload)
	if validationErrors != nil {
		return models.SendFailedValidationResponse(c, validationErrors)
	}

	walletTransaction, wallet, err := h.walletTransactionService.UpdateWalletTransaction(*payload)
	if err != nil {
		return models.SendInternalServerErrorResponse(c, err.Error())
	}

	return models.SendSuccessResponse(c, "berhasil update wallet transaction", models.ToWalletTransactionWithWalletResponse(walletTransaction, &wallet))
}

func (h *WalletTransactionHandler) GetWalletTransactionById(c echo.Context) error {
	id := c.Param("id")
	walletTransaction, wallet, err := h.walletTransactionService.GetWalletTransactionById(id)
	if err != nil {
		return models.SendInternalServerErrorResponse(c, err.Error())
	}

	return models.SendSuccessResponse(c, "berhasil mendapatkan wallet transaction", models.ToWalletTransactionWithWalletResponse(walletTransaction, &wallet))
}

func (h *WalletTransactionHandler) DeleteWalletTransaction(c echo.Context) error {
	id := c.Param("id")
	if err := h.walletTransactionService.DeleteWalletTransaction(id); err != nil {
		return models.SendInternalServerErrorResponse(c, err.Error())
	}

	return models.SendSuccessResponse(c, "berhasil menghapus wallet transaction", nil)
}

func (h *WalletTransactionHandler) GetWalletTransactionByUserIdPaginated(c echo.Context) error {
	userId := c.Param("userId")
	page := 1
	pageSize := 10

	if p := c.QueryParam("page"); p != "" {
		fmt.Sscanf(p, "%d", &page)
	}
	if ps := c.QueryParam("page_size"); ps != "" {
		fmt.Sscanf(ps, "%d", &pageSize)
	}

	req := models.RequestParams{Page: page, PageSize: pageSize}
	result, err := h.walletTransactionService.GetWalletTransactionByUserIdPaginated(req, userId)
	if err != nil {
		return models.SendInternalServerErrorResponse(c, err.Error())
	}

	return models.SendPaginatedResponse(c, "berhasil mendapatkan wallet transaction", result.PaginatorInfo, result.Data)
}

func (h *WalletTransactionHandler) GetWalletTransactionByWalletIdPaginated(c echo.Context) error {
	id := c.Param("walletId")
	page := 1
	pageSize := 10

	if p := c.QueryParam("page"); p != "" {
		fmt.Sscanf(p, "%d", &page)
	}
	if ps := c.QueryParam("page_size"); ps != "" {
		fmt.Sscanf(ps, "%d", &pageSize)
	}

	req := models.RequestParams{Page: page, PageSize: pageSize}
	result, err := h.walletTransactionService.GetWalletTransactionByWalletIdPaginated(req, id)
	if err != nil {
		return models.SendInternalServerErrorResponse(c, err.Error())
	}

	return models.SendPaginatedResponse(c, "berhasil mendapatkan wallet transaction", result.PaginatorInfo, result.Data)
}
