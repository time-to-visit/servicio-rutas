package handler

import (
	"service-routes/cmd/entry"
	"service-routes/internal/domain/usecase"
	"service-routes/internal/domain/validator"

	"github.com/labstack/echo/v4"
)

func NewRoutesEntry(e *echo.Echo, routeUseCase usecase.RoutesUseCase, auth func(next echo.HandlerFunc) echo.HandlerFunc) *echo.Echo {
	routeEntry := entry.NewRouteEntry(routeUseCase)
	e.POST("/route", routeEntry.RegisterRoute, auth, validator.ValidateRoutes)
	e.DELETE("/route/:ID", routeEntry.DeleteRoute, auth)
	e.GET("/route", routeEntry.FindRoute, auth)
	e.GET("/route/:ID", routeEntry.FindRouteOne, auth)
	return e
}
