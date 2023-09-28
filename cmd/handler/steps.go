package handler

import (
	"service-routes/cmd/entry"
	"service-routes/internal/domain/usecase"
	"service-routes/internal/domain/validator"

	"github.com/labstack/echo/v4"
)

func NewStepsEntry(e *echo.Echo, stepUseCase usecase.StepsUseCase, auth func(next echo.HandlerFunc) echo.HandlerFunc) *echo.Echo {
	stepEntry := entry.NewStepEntry(stepUseCase)
	e.POST("/routes/step", stepEntry.InsertStep, auth, validator.ValidateSteps)
	e.DELETE("/routes/step/:ID", stepEntry.DeleteStep, auth)
	e.GET("/routes/step/route/:ID", stepEntry.FindStep, auth)
	e.GET("/routes/step/:ID", stepEntry.FindStepOne, auth)
	return e
}
