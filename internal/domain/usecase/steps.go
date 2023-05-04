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

type StepsUseCase struct {
	repoSteps repository.IRepositorySteps
	file      storage.IGCImageRepo
}

func NewStepsUseCase(repoSteps repository.IRepositorySteps, file storage.IGCImageRepo) StepsUseCase {
	return StepsUseCase{
		repoSteps,
		file,
	}
}

func (r *StepsUseCase) InsertStep(ctx context.Context, step entity.Steps) (interface{}, int) {
	pathname, err := r.file.SetFile(ctx, step.Cover, "step/step-%s.png")
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "hubo un problema con storage", nil), http.StatusBadRequest
	}
	step.Cover = pathname
	newStep, err := r.repoSteps.InsertStep(step)
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "hubo un problema al insertar el paso", nil), http.StatusBadRequest
	}
	return objectValues.NewResponseWithData(http.StatusOK, "ok", "paso insertada exitosamente", newStep), http.StatusOK
}

func (r *StepsUseCase) FindStep(idRoute int64) (interface{}, int) {
	steps, err := r.repoSteps.FindStep(idRoute)
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "problemas al buscar la informacion", nil), http.StatusBadRequest
	}
	return objectValues.NewResponseWithData(http.StatusOK, "ok", "sucess", steps), http.StatusOK
}

func (r *StepsUseCase) FindStepOne(idStep int64) (interface{}, int) {
	step, err := r.repoSteps.FindStepOne(idStep)
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "problemas al buscar la informacion", nil), http.StatusBadRequest
	}
	return objectValues.NewResponseWithData(http.StatusOK, "ok", "sucess", step), http.StatusOK

}

func (r *StepsUseCase) DeleteStep(ctx context.Context, idStep int64) (interface{}, int) {
	step, err := r.repoSteps.FindStepOne(idStep)
	if err != nil || step == nil || step.ID == 0 {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "no se encontro el steps", nil), http.StatusBadRequest
	}
	objectName := utils.ExtractObjectName(step.Cover)
	err = r.file.DeleteFile(ctx, "step/%s", objectName)
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "hubo un problema con storage", nil), http.StatusBadRequest
	}
	err = r.repoSteps.DeleteStep(idStep)
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "hubo un problema al eliminar el paso", nil), http.StatusBadRequest
	}
	return objectValues.NewResponseWithData(http.StatusOK, "ok", "paso eliminada exitosamente", nil), http.StatusOK
}
