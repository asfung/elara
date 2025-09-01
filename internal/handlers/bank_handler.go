package handlers

import (
	"net/http"

	"github.com/asfung/elara/internal/models"
	"github.com/asfung/elara/internal/services"
	"github.com/labstack/echo/v4"
)

type BankHandler struct {
	*Handler
	bankService services.BankService
}

func NewBankHandler(bankService services.BankService) *BankHandler {
	return &BankHandler{bankService: bankService}
}

func (h *BankHandler) CreateBank(c echo.Context) error {
	payload := new(models.AddBankRequest)
	if err := h.BindBodyRequest(c, payload); err != nil {
		return models.SendBadRequestResponse(c, err.Error())
	}

	validationErrors := h.ValidateBodyRequest(payload)
	if validationErrors != nil {
		return models.SendFailedValidationResponse(c, validationErrors)
	}

	bank, err := h.bankService.CreateBank(*payload)
	if err != nil {
		return models.SendInternalServerErrorResponse(c, err.Error())
	}

	return models.SendSuccessResponse(c, "Bank created successfully", models.ToBankResponse(bank))
}

func (h *BankHandler) UpdateBank(c echo.Context) error {
	id := c.Param("id")
	payload := new(models.UpdateBankRequest)
	if err := h.BindBodyRequest(c, payload); err != nil {
		return models.SendBadRequestResponse(c, err.Error())
	}

	validationErros := h.ValidateBodyRequest(payload)
	if validationErros != nil {
		return models.SendFailedValidationResponse(c, validationErros)
	}
	payload.ID = id

	bank, err := h.bankService.UpdateBank(*payload)
	if err != nil {
		return models.SendInternalServerErrorResponse(c, err.Error())
	}
	return models.SendSuccessResponse(c, "Bank berhaisl update", models.ToBankResponse(bank))
}

func (h *BankHandler) GetById(c echo.Context) error {
	id := c.Param("id")
	bank, err := h.bankService.GetBankById(id)
	if err != nil {
		return models.SendErrorResponse(c, "Bank Not Found", http.StatusNotFound)
	}
	return models.SendSuccessResponse(c, "Get Bank Successful", models.ToBankResponse(bank))
}

func (h *BankHandler) DeleteBank(c echo.Context) error {
	id := c.Param("id")

	if err := h.bankService.DeleteBank(id); err != nil {
		return models.SendErrorResponse(c, err.Error(), http.StatusInternalServerError)
	}
	return models.SendSuccessResponse(c, "berhasil menghapus bank", nil)
}
