package handler_test

import (
	"service-routes/cmd/handler"
	"service-routes/internal/domain/usecase"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func Test_HandlerResource(t *testing.T) {
	e := handler.NewResourceEntry(echo.New(), usecase.ResourcesUseCase{}, func(next echo.HandlerFunc) echo.HandlerFunc { return nil })
	assert.NotNil(t, e)
}
