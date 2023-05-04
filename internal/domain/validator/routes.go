package validator

import (
	"net/http"
	"service-routes/internal/domain/entity"
	validatorPer "service-routes/internal/infra/validation"

	validatorV "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func ValidateRoutes(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		v := validatorPer.NewValidator()
		routes := new(entity.Routes)

		_ = c.Bind(&routes)
		if err := v.Struct(routes); err != nil {
			errs := err.(validatorV.ValidationErrors)
			return c.JSON(http.StatusBadRequest, validatorPer.GenerateMessage(v, errs))
		}
		c.Set("routes", routes)
		return next(c)
	}
}
