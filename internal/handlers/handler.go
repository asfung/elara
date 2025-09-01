package handlers

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/asfung/elara/internal/models"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Handler struct {
	DB     *gorm.DB
	Logger echo.Logger
}

func (h *Handler) BindBodyRequest(c echo.Context, payload interface{}) error {
	if err := (&echo.DefaultBinder{}).BindBody(c, payload); err != nil {
		c.Logger().Error(err)
		return errors.New("failed to bind body, make sure you are sending a valid payload")
	}
	return nil
}

func (h *Handler) ValidateBodyRequest(payload interface{}) []*models.ValidationError {
	validate := validator.New()

	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	if err := validate.Struct(payload); err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			var validationErrors []*models.ValidationError
			for _, fe := range ve {
				validationErrors = append(validationErrors, &models.ValidationError{
					Error:     fmt.Sprintf("%s failed on the '%s' rule", fe.Field(), fe.Tag()),
					Key:       fe.Field(),
					Condition: fe.Tag(),
				})
			}
			return validationErrors
		}
		return []*models.ValidationError{{
			Error: err.Error(),
		}}
	}
	return nil
}
