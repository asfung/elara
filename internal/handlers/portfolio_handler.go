package handlers

import (
	"net/http"

	"github.com/asfung/elara/internal/models"
	"github.com/asfung/elara/internal/services"
	"github.com/labstack/echo/v4"
)

type PortfolioHandler struct {
	*Handler
	portfolioService services.PortfolioService
}

func NewPortfolioHandler(portfolioService services.PortfolioService) *PortfolioHandler {
	return &PortfolioHandler{
		portfolioService: portfolioService,
	}
}

func (h *PortfolioHandler) CreatePortfolio(c echo.Context) error {
	payload := new(models.AddAPortfolioRequest)
	if err := h.BindBodyRequest(c, payload); err != nil {
		return models.SendBadRequestResponse(c, err.Error())
	}

	validationErrors := h.ValidateBodyRequest(payload)
	if validationErrors != nil {
		return models.SendFailedValidationResponse(c, validationErrors)
	}

	portfolio, err := h.portfolioService.CreatePortfolio(*payload)
	if err != nil {
		return models.SendInternalServerErrorResponse(c, err.Error())
	}

	return models.SendResponse(c, true, "membuat portfolio", portfolio, http.StatusCreated)
}

func (h *PortfolioHandler) UpdatePortfolio(c echo.Context) error {
	payload := new(models.UpdatePortfolioRequest)
	if err := h.BindBodyRequest(c, payload); err != nil {
		return models.SendBadRequestResponse(c, err.Error())
	}

	validationErrors := h.ValidateBodyRequest(payload)
	if validationErrors != nil {
		return models.SendFailedValidationResponse(c, validationErrors)
	}

	portfolio, err := h.portfolioService.UpdatePortfolio(*payload)
	if err != nil {
		return models.SendInternalServerErrorResponse(c, err.Error())
	}

	return models.SendSuccessResponse(c, "update portfolio", portfolio)
}

func (h *PortfolioHandler) GetPortfolioById(c echo.Context) error {
	id := c.Param("id")
	portfolio, err := h.portfolioService.GetPortfolioById(id)
	if err != nil {
		return models.SendInternalServerErrorResponse(c, err.Error())
	}

	return models.SendSuccessResponse(c, "portfolio data", portfolio)
}

func (h *PortfolioHandler) DeletePortfolio(c echo.Context) error {
	id := c.Param("id")

	if err := h.portfolioService.DeletePortfolio(id); err != nil {
		return models.SendInternalServerErrorResponse(c, err.Error())
	}

	return models.SendSuccessResponse(c, "portfolio terhapus", nil)
}
