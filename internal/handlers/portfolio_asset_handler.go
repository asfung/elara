package handlers

import (
	"github.com/asfung/elara/internal/models"
	"github.com/asfung/elara/internal/services"
	"github.com/labstack/echo/v4"
)

type PortfolioAssetHandler struct {
	*Handler
	portfolioAssetService services.PortfolioAssetService
}

func NewPortfolioAssetHandler(portfolioAssetService services.PortfolioAssetService) *PortfolioAssetHandler {
	return &PortfolioAssetHandler{
		portfolioAssetService: portfolioAssetService,
	}
}

func (h *PortfolioAssetHandler) CreatePortfolioAsset(c echo.Context) error {
	payload := new(models.AddAPortfolioAssetRequest)
	if err := h.BindBodyRequest(c, payload); err != nil {
		return models.SendBadRequestResponse(c, err.Error())
	}

	validationErrors := h.ValidateBodyRequest(payload)
	if validationErrors != nil {
		return models.SendFailedValidationResponse(c, validationErrors)
	}

	portfolioAsset, err := h.portfolioAssetService.CreatePortfolioAsset(*payload)
	if err != nil {
		return models.SendInternalServerErrorResponse(c, err.Error())
	}

	return models.SendSuccessResponse(c, "membuat portfolio", portfolioAsset)
}

func (h *PortfolioAssetHandler) UpdatePortfolioAsset(c echo.Context) error {
	payload := new(models.UpdatePortfolioAssetRequest)
	if err := h.BindBodyRequest(c, payload); err != nil {
		return models.SendBadRequestResponse(c, err.Error())
	}

	validationErrors := h.ValidateBodyRequest(payload)
	if validationErrors != nil {
		return models.SendFailedValidationResponse(c, validationErrors)
	}

	portfolioAsset, err := h.portfolioAssetService.UpdatePortfolioAsset(*payload)
	if err != nil {
		return models.SendInternalServerErrorResponse(c, err.Error())
	}

	return models.SendSuccessResponse(c, "update portfolio asset", portfolioAsset)
}

func (h *PortfolioAssetHandler) GetPortfolioAssetById(c echo.Context) error {
	id := c.Param("id")
	portfolioAsset, err := h.portfolioAssetService.GetPortfolioAssetById(id)
	if err != nil {
		return models.SendInternalServerErrorResponse(c, err.Error())
	}

	return models.SendSuccessResponse(c, "portfolio asset data", portfolioAsset)
}

func (h *PortfolioAssetHandler) DeletePortfolioAsset(c echo.Context) error {
	id := c.Param("id")
	if err := h.portfolioAssetService.DeletePortfolioAsset(id); err != nil {
		return models.SendInternalServerErrorResponse(c, err.Error())
	}

	return models.SendSuccessResponse(c, "portfolio asset terhapus", nil)
}
