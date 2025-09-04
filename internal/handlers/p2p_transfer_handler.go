package handlers

import (
	"github.com/asfung/elara/internal/models"
	"github.com/asfung/elara/internal/services"
	"github.com/labstack/echo/v4"
)

type P2PTransferHandler struct {
	*Handler
	p2pTransferService services.P2PTransferService
}

func NewP2PTransferHandler(p2pTransferService services.P2PTransferService) *P2PTransferHandler {
	return &P2PTransferHandler{p2pTransferService: p2pTransferService}
}

func (h *P2PTransferHandler) CreateP2PTransfer(c echo.Context) error {
	payload := new(models.AddP2PTransferRequest)
	if err := h.BindBodyRequest(c, payload); err != nil {
		return models.SendBadRequestResponse(c, err.Error())
	}

	validationErors := h.ValidateBodyRequest(payload)
	if validationErors != nil {
		return models.SendFailedValidationResponse(c, validationErors)
	}

	p2pTransfer, err := h.p2pTransferService.CreateP2PTransfer(*payload)
	if err != nil {
		return models.SendInternalServerErrorResponse(c, err.Error())
	}
	return models.SendSuccessResponse(c, "berhasil membuat p2p transfer", models.ToP2PTransferResponse(p2pTransfer))
}

func (h *P2PTransferHandler) Update2PTransfer(c echo.Context) error {
	payload := new(models.UpdateP2PTransferRequest)
	if err := h.BindBodyRequest(c, payload); err != nil {
		return models.SendBadRequestResponse(c, err.Error())
	}

	validationErors := h.ValidateBodyRequest(payload)
	if validationErors != nil {
		return models.SendFailedValidationResponse(c, validationErors)
	}

	p2pTransfer, err := h.p2pTransferService.UpdateP2PTransfer(*payload)
	if err != nil {
		return models.SendInternalServerErrorResponse(c, err.Error())
	}

	return models.SendSuccessResponse(c, "berhasil update p2p transfer", models.ToP2PTransferResponse(p2pTransfer))
}

func (h *P2PTransferHandler) GetP2PTransferById(c echo.Context) error {
	id := c.Param("id")
	p2pTransfer, err := h.p2pTransferService.GetP2PTransferById(id)
	if err != nil {
		return models.SendInternalServerErrorResponse(c, err.Error())
	}
	return models.SendSuccessResponse(c, "berhasil mengambil data", models.ToP2PTransferResponse(p2pTransfer))
}

func (h *P2PTransferHandler) DeleteP2PTransfer(c echo.Context) error {
	id := c.Param("id")
	if err := h.p2pTransferService.DeleteP2PTransfer(id); err != nil {
		return models.SendInternalServerErrorResponse(c, err.Error())
	}
	return models.SendSuccessResponse(c, "berhasil menghapus data", nil)
}
