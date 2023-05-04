package validator

import (
	"net/http"
	"service-routes/internal/domain/entity"
	objectValues "service-routes/internal/domain/object_values"
	validatorPer "service-routes/internal/infra/validation"

	validatorV "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func ValidateComment(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		v := validatorPer.NewValidator()
		comment := new(entity.Comments)
		auth := c.Get("auth").(objectValues.Auth)
		_ = c.Bind(&comment)
		if err := v.Struct(comment); err != nil {
			errs := err.(validatorV.ValidationErrors)
			return c.JSON(http.StatusBadRequest, validatorPer.GenerateMessage(v, errs))
		}
		comment.NameUser = auth.Data.Nombres
		comment.IDUser = int64(auth.Data.ID)
		c.Set("comment", comment)
		return next(c)
	}
}
