package usecase

import (
	"context"
	"net/http"
	"service-routes/internal/domain/entity"
	objectValues "service-routes/internal/domain/object_values"
	"service-routes/internal/domain/repository"
	"service-routes/internal/infra/storage"
	"service-routes/internal/utils"
)

type RoutesUseCase struct {
	repoRoutes repository.IRepositoryRoutes
	file       storage.IGCImageRepo
}

func NewRoutesUseCase(repoRoutes repository.IRepositoryRoutes, file storage.IGCImageRepo) RoutesUseCase {
	return RoutesUseCase{
		repoRoutes,
		file,
	}
}

func (r *RoutesUseCase) RegisterRoute(ctx context.Context, route entity.Routes) (interface{}, int) {
	pathname, err := r.file.SetFile(ctx, route.Cover, "route/route-%s.png")
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "hubo un problema con storage", nil), http.StatusBadRequest
	}
	route.Cover = pathname
	newRoute, err := r.repoRoutes.InsertRoute(route)
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "hubo un problema al insertar la ruta", nil), http.StatusBadRequest
	}
	return objectValues.NewResponseWithData(http.StatusOK, "ok", "ruta insertada exitosamente", newRoute), http.StatusOK
}

func (r *RoutesUseCase) DeleteRoute(ctx context.Context, idRoute int64) (interface{}, int) {
	route, err := r.repoRoutes.FindRouteOne(idRoute)
	if err != nil || route == nil || route.ID == 0 {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "no se encontro el steps", nil), http.StatusBadRequest
	}
	objectName := utils.ExtractObjectName(route.Cover)
	err = r.file.DeleteFile(ctx, "route/%s", objectName)
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "hubo un problema con storage", nil), http.StatusBadRequest
	}
	err = r.repoRoutes.DelRoute(idRoute)
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "hubo un problema al eliminar la ruta", nil), http.StatusBadRequest
	}
	return objectValues.NewResponseWithData(http.StatusOK, "ok", "ruta eliminada exitosamente", nil), http.StatusOK
}

func (r *RoutesUseCase) FindRoute(filter map[string]interface{}) (interface{}, int) {
	routes, err := r.repoRoutes.FindRoute(filter)
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "problemas al buscar la informacion", nil), http.StatusBadRequest
	}
	return objectValues.NewResponseWithData(http.StatusOK, "ok", "sucess", routes), http.StatusOK
}

func (r *RoutesUseCase) FindRouteOne(idRoute int64) (interface{}, int) {
	route, err := r.repoRoutes.FindRouteOne(idRoute)
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "problemas al buscar la informacion", nil), http.StatusBadRequest
	}
	return objectValues.NewResponseWithData(http.StatusOK, "ok", "sucess", route), http.StatusOK
}
