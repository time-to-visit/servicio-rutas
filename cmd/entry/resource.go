package entry

import (
	"service-routes/internal/domain/entity"
	"service-routes/internal/domain/usecase"
	"strconv"

	"github.com/labstack/echo/v4"
)

type resourceEntry struct {
	resourceCaseuse usecase.ResourcesUseCase
}

func NewResourceEntry(resourceCaseuse usecase.ResourcesUseCase) *resourceEntry {
	return &resourceEntry{
		resourceCaseuse,
	}
}

func (r *resourceEntry) InsertResource(context echo.Context) error {
	resource := context.Get("resource").(*entity.Resource)
	response, status := r.resourceCaseuse.InsertResource(context.Request().Context(), *resource)
	return context.JSON(status, response)
}

func (r *resourceEntry) DeleteResource(context echo.Context) error {
	idResource, _ := strconv.Atoi(context.Param("ID"))
	response, status := r.resourceCaseuse.DeleteResource(context.Request().Context(), int64(idResource))
	return context.JSON(status, response)
}
