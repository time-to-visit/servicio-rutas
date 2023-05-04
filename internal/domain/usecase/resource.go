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

type ResourcesUseCase struct {
	repoResources repository.IRepositoryResource
	file          storage.IGCImageRepo
}

func NewResourcesUseCase(repoResources repository.IRepositoryResource, file storage.IGCImageRepo) ResourcesUseCase {
	return ResourcesUseCase{
		repoResources,
		file,
	}
}

func (r *ResourcesUseCase) InsertResource(ctx context.Context, resource entity.Resource) (interface{}, int) {
	pathname, err := r.file.SetFile(ctx, resource.UrlResource, "resource-route/resorce-%s.png")
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "hubo un problema con storage", nil), http.StatusBadRequest
	}
	resource.UrlResource = pathname
	newResource, err := r.repoResources.AddResource(resource)
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "hubo un problema al insertar el recurso", nil), http.StatusBadRequest
	}
	return objectValues.NewResponseWithData(http.StatusOK, "ok", "recurso insertada exitosamente", newResource), http.StatusOK
}

func (r *ResourcesUseCase) DeleteResource(ctx context.Context, idResource int64) (interface{}, int) {
	resource, err := r.repoResources.FindOneResource(idResource)
	if err != nil || resource == nil || resource.ID == 0 {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "no se encontro el resource", nil), http.StatusBadRequest
	}
	objectName := utils.ExtractObjectName(resource.UrlResource)
	err = r.file.DeleteFile(ctx, "resource-route/%s", objectName)
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "hubo un problema con storage", nil), http.StatusBadRequest
	}
	err = r.repoResources.DelResource(idResource)
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "hubo un problema al eliminar el recurso", nil), http.StatusBadRequest
	}
	return objectValues.NewResponseWithData(http.StatusOK, "ok", "recurso eliminada exitosamente", nil), http.StatusOK
}
