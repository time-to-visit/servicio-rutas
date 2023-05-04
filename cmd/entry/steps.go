package entry

import (
	"service-routes/internal/domain/entity"
	"service-routes/internal/domain/usecase"
	"strconv"

	"github.com/labstack/echo/v4"
)

type stepEntry struct {
	stepCaseuse usecase.StepsUseCase
}

func NewStepEntry(stepCaseuse usecase.StepsUseCase) *stepEntry {
	return &stepEntry{
		stepCaseuse,
	}
}

func (r *stepEntry) InsertStep(context echo.Context) error {
	step := context.Get("step").(*entity.Steps)
	response, status := r.stepCaseuse.InsertStep(context.Request().Context(), *step)
	return context.JSON(status, response)
}

func (r *stepEntry) FindStep(context echo.Context) error {
	idRoute, _ := strconv.Atoi(context.Param("ID"))
	response, status := r.stepCaseuse.FindStep(int64(idRoute))
	return context.JSON(status, response)
}

func (r *stepEntry) FindStepOne(context echo.Context) error {
	idStep, _ := strconv.Atoi(context.Param("ID"))
	response, status := r.stepCaseuse.FindStepOne(int64(idStep))
	return context.JSON(status, response)
}

func (r *stepEntry) DeleteStep(context echo.Context) error {
	idStep, _ := strconv.Atoi(context.Param("ID"))
	response, status := r.stepCaseuse.DeleteStep(context.Request().Context(), int64(idStep))
	return context.JSON(status, response)
}
