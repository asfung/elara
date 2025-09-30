package handlers

import (
	"net/http"

	"github.com/asfung/elara/internal/models"
	"github.com/asfung/elara/internal/services"
	"github.com/labstack/echo/v4"
)

type CardHandler struct {
	*Handler
	cardService services.CardService
}

func NewCardHandler(cardService services.CardService) *CardHandler {
	return &CardHandler{cardService: cardService}
}

func (h *CardHandler) CreateCard(c echo.Context) error {
	payload := new(models.AddCardRequest)
	if err := h.BindBodyRequest(c, payload); err != nil {
		return models.SendBadRequestResponse(c, err.Error())
	}

	validationErrors := h.ValidateBodyRequest(payload)
	if validationErrors != nil {
		return models.SendFailedValidationResponse(c, validationErrors)
	}

	card, err := h.cardService.CreateCard(*payload)
	if err != nil {
		return models.SendErrorResponse(c, err.Error(), http.StatusInternalServerError)
	}

	return models.SendSuccessResponse(c, "berhasil membuat card", models.ToCardResponse(card))
}

func (h *CardHandler) UpdateCard(c echo.Context) error {
	id := c.Param("id")
	payload := new(models.UpdateCardRequest)
	if err := h.BindBodyRequest(c, payload); err != nil {
		return models.SendBadRequestResponse(c, err.Error())
	}

	validationErrors := h.ValidateBodyRequest(payload)
	if validationErrors != nil {
		return models.SendFailedValidationResponse(c, validationErrors)
	}

	payload.ID = id
	card, err := h.cardService.UpdateCard(*payload)
	if err != nil {
		return models.SendErrorResponse(c, err.Error(), http.StatusInternalServerError)
	}
	return models.SendSuccessResponse(c, "berhasil update card", models.ToCardResponse(card))
}

func (h *CardHandler) GetCardById(c echo.Context) error {
	id := c.Param("id")
	card, err := h.cardService.GetCardById(id)
	if err != nil {
		return models.SendErrorResponse(c, err.Error(), http.StatusInternalServerError)
	}
	return models.SendSuccessResponse(c, "berhasil get card", models.ToCardResponse(card))
}

func (h *CardHandler) DeleteCard(c echo.Context) error {
	id := c.Param("id")
	if err := h.cardService.DeleteCard(id); err != nil {
		return models.SendErrorResponse(c, err.Error(), http.StatusInternalServerError)
	}
	return models.SendSuccessResponse(c, "berhasil menghapus card", nil)
}
