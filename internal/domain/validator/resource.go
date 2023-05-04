package validator

import (
	"net/http"
	"service-routes/internal/domain/entity"
	validatorPer "service-routes/internal/infra/validation"

	validatorV "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func ValidateResource(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		v := validatorPer.NewValidator()
		resource := new(entity.Resource)

		_ = c.Bind(&resource)
		if err := v.Struct(resource); err != nil {
			errs := err.(validatorV.ValidationErrors)
			return c.JSON(http.StatusBadRequest, validatorPer.GenerateMessage(v, errs))
		}
		c.Set("resource", resource)
		return next(c)
	}
}
