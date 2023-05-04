package handler

import (
	"service-routes/cmd/entry"
	"service-routes/internal/domain/usecase"
	"service-routes/internal/domain/validator"

	"github.com/labstack/echo/v4"
)

func NewResourceEntry(e *echo.Echo, resourceUseCase usecase.ResourcesUseCase, auth func(next echo.HandlerFunc) echo.HandlerFunc) *echo.Echo {
	resourceEntry := entry.NewResourceEntry(resourceUseCase)
	e.POST("/resource", resourceEntry.InsertResource, auth, validator.ValidateResource)
	e.DELETE("/resource/:ID", resourceEntry.DeleteResource, auth)
	return e
}
