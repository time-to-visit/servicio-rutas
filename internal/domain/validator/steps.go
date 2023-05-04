package validator

import (
	"net/http"
	"service-routes/internal/domain/entity"
	validatorPer "service-routes/internal/infra/validation"

	validatorV "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func ValidateSteps(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		v := validatorPer.NewValidator()
		step := new(entity.Steps)

		_ = c.Bind(&step)
		if err := v.Struct(step); err != nil {
			errs := err.(validatorV.ValidationErrors)
			return c.JSON(http.StatusBadRequest, validatorPer.GenerateMessage(v, errs))
		}
		c.Set("step", step)
		return next(c)
	}
}
