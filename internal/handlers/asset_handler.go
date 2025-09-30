package handlers

import (
	"net/http"

	"github.com/asfung/elara/internal/models"
	"github.com/asfung/elara/internal/services"
	"github.com/labstack/echo/v4"
)

type AssetHandler struct {
	*Handler
	assetService services.AssetService
}

func NewAssetHandler(assetService services.AssetService) *AssetHandler {
	return &AssetHandler{
		assetService: assetService,
	}
}

func (h *AssetHandler) CreateAsset(c echo.Context) error {
	payload := new(models.AddAssetRequest)
	if err := h.BindBodyRequest(c, payload); err != nil {
		return models.SendBadRequestResponse(c, err.Error())
	}

	validationErrors := h.ValidateBodyRequest(payload)
	if validationErrors != nil {
		return models.SendFailedValidationResponse(c, validationErrors)
	}

	asset, err := h.assetService.CreateAsset(*payload)
	if err != nil {
		return models.SendInternalServerErrorResponse(c, err.Error())
	}

	return models.SendResponse(c, true, "membuat asset", asset, http.StatusCreated)
}

func (h *AssetHandler) UpdateAsset(c echo.Context) error {
	payload := new(models.UpdateAssetRequest)
	if err := h.BindBodyRequest(c, payload); err != nil {
		return models.SendBadRequestResponse(c, err.Error())
	}

	validationErrors := h.ValidateBodyRequest(payload)
	if validationErrors != nil {
		return models.SendFailedValidationResponse(c, validationErrors)
	}

	asset, err := h.assetService.UpdateAsset(*payload)
	if err != nil {
		return models.SendInternalServerErrorResponse(c, err.Error())
	}

	return models.SendSuccessResponse(c, "mengupdate asset", asset)
}

func (h *AssetHandler) GetAssetById(c echo.Context) error {
	id := c.Param("id")
	asset, err := h.assetService.GetAssetById(id)
	if err != nil {
		return models.SendInternalServerErrorResponse(c, err.Error())
	}

	return models.SendSuccessResponse(c, "asset data", asset)
}

func (h *AssetHandler) DeleteAsset(c echo.Context) error {
	id := c.Param("id")
	if err := h.assetService.DeleteAsset(id); err != nil {
		return models.SendInternalServerErrorResponse(c, err.Error())
	}
	return models.SendSuccessResponse(c, "asset terhapus", nil)
}
