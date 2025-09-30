package handlers

import (
	"net/http"

	"github.com/asfung/elara/internal/models"
	"github.com/asfung/elara/internal/services"
	"github.com/labstack/echo/v4"
)

type BankAccountHandler struct {
	*Handler
	bankAccountService services.BankAccountService
}

func NewBankAccountHandler(bankAccountService services.BankAccountService) *BankAccountHandler {
	return &BankAccountHandler{
		bankAccountService: bankAccountService,
	}
}

func (h *BankAccountHandler) CreateBankAccount(c echo.Context) error {
	payload := new(models.AddBankAccountRequest)
	if err := h.BindBodyRequest(c, payload); err != nil {
		return models.SendBadRequestResponse(c, err.Error())
	}

	validationErrors := h.ValidateBodyRequest(payload)
	if validationErrors != nil {
		return models.SendFailedValidationResponse(c, validationErrors)
	}

	bankAccount, bank, err := h.bankAccountService.CreateBankAccount(*payload)
	if err != nil {
		return models.SendErrorResponse(c, err.Error(), http.StatusInternalServerError)
	}
	return models.SendSuccessResponse(c, "berhasil membuat bank account", models.ToBankAccountWithBankResponse(bankAccount, bank))
}

func (h *BankAccountHandler) UpdateBankAccount(c echo.Context) error {
	id := c.Param("id")
	payload := new(models.UpdateBankAccountRequest)
	if err := h.BindBodyRequest(c, payload); err != nil {
		return models.SendBadRequestResponse(c, err.Error())
	}

	validationErrors := h.ValidateBodyRequest(payload)
	if validationErrors != nil {
		return models.SendFailedValidationResponse(c, validationErrors)
	}
	payload.ID = id

	bankAccount, bank, err := h.bankAccountService.UpdateBankAccount(*payload)
	if err != nil {
		return models.SendErrorResponse(c, err.Error(), http.StatusInternalServerError)
	}
	return models.SendSuccessResponse(c, "berhasil update bank account", models.ToBankAccountWithBankResponse(bankAccount, bank))
}

func (h *BankAccountHandler) GetById(c echo.Context) error {
	id := c.Param("id")
	bankAccount, bank, err := h.bankAccountService.GetBankAccountById(id)
	if err != nil {
		return models.SendErrorResponse(c, err.Error(), http.StatusInternalServerError)
	}
	return models.SendSuccessResponse(c, "get Bank Account successful", models.ToBankAccountWithBankResponse(bankAccount, bank))
}

func (h *BankAccountHandler) DeleteBankAccount(c echo.Context) error {
	id := c.Param("id")
	if err := h.bankAccountService.DeleteBankAccount(id); err != nil {
		return models.SendErrorResponse(c, err.Error(), http.StatusInsufficientStorage)
	}
	return models.SendSuccessResponse(c, "berhasil delete bank account", nil)
}
