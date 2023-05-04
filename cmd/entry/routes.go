package entry

import (
	"service-routes/internal/domain/entity"
	"service-routes/internal/domain/usecase"
	"strconv"

	"github.com/labstack/echo/v4"
)

type routeEntry struct {
	routeCaseuse usecase.RoutesUseCase
}

func NewRouteEntry(routeCaseuse usecase.RoutesUseCase) *routeEntry {
	return &routeEntry{
		routeCaseuse,
	}
}

func (r *routeEntry) RegisterRoute(context echo.Context) error {
	routes := context.Get("routes").(*entity.Routes)
	response, status := r.routeCaseuse.RegisterRoute(context.Request().Context(), *routes)
	return context.JSON(status, response)
}

func (r *routeEntry) DeleteRoute(context echo.Context) error {
	id, _ := strconv.Atoi(context.Param("ID"))
	response, status := r.routeCaseuse.DeleteRoute(context.Request().Context(), int64(id))
	return context.JSON(status, response)
}

func (r *routeEntry) FindRoute(context echo.Context) error {
	filter := make(map[string]interface{})
	response, status := r.routeCaseuse.FindRoute(filter)
	return context.JSON(status, response)
}

func (r *routeEntry) FindRouteOne(context echo.Context) error {
	id, _ := strconv.Atoi(context.Param("ID"))
	response, status := r.routeCaseuse.FindRouteOne(int64(id))
	return context.JSON(status, response)
}
