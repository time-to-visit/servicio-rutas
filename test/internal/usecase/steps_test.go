package usecase_test

import (
	"context"
	"errors"
	"net/http"
	"service-routes/internal/domain/entity"
	"service-routes/internal/domain/usecase"
	"service-routes/mocks"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/mock"
)

func Test_SeeStepOk(t *testing.T) {
	repoCat := mocks.NewIRepositorySteps(t)
	cats := []entity.Steps{
		{},
	}
	objectFile := mocks.NewIGCImageRepo(t)
	usecaseCat := usecase.NewStepsUseCase(repoCat, objectFile)

	repoCat.On("FindStep", mock.Anything).Return(&cats, nil)

	_, status := usecaseCat.FindStep(1)
	assert.Equal(t, http.StatusOK, status)
}

func Test_SeeStepErr(t *testing.T) {
	repoCat := mocks.NewIRepositorySteps(t)
	cats := []entity.Steps{
		{},
	}
	objectFile := mocks.NewIGCImageRepo(t)
	usecaseCat := usecase.NewStepsUseCase(repoCat, objectFile)

	repoCat.On("FindStep", mock.Anything).Return(&cats, errors.New("err"))

	_, status := usecaseCat.FindStep(1)
	assert.Equal(t, http.StatusBadRequest, status)
}

func Test_SeeStepOneOk(t *testing.T) {
	repoCat := mocks.NewIRepositorySteps(t)
	cats := entity.Steps{}

	objectFile := mocks.NewIGCImageRepo(t)
	usecaseCat := usecase.NewStepsUseCase(repoCat, objectFile)

	repoCat.On("FindStepOne", mock.Anything).Return(&cats, nil)

	_, status := usecaseCat.FindStepOne(1)
	assert.Equal(t, http.StatusOK, status)
}

func Test_SeeStepOneErr(t *testing.T) {
	repoCat := mocks.NewIRepositorySteps(t)
	cats := entity.Steps{}
	objectFile := mocks.NewIGCImageRepo(t)
	usecaseCat := usecase.NewStepsUseCase(repoCat, objectFile)

	repoCat.On("FindStepOne", mock.Anything).Return(&cats, errors.New("err"))

	_, status := usecaseCat.FindStepOne(1)
	assert.Equal(t, http.StatusBadRequest, status)
}

func Test_DeleteStepOk(t *testing.T) {
	cat := entity.Steps{
		Cover: "https://google.com/dasdasdasdas",
		Model: entity.Model{
			ID: 1,
		},
	}
	repoCat := mocks.NewIRepositorySteps(t)
	objectFile := mocks.NewIGCImageRepo(t)
	usecaseCat := usecase.NewStepsUseCase(repoCat, objectFile)
	repoCat.On("FindStepOne", mock.Anything).Return(&cat, nil)
	repoCat.On("DeleteStep", mock.Anything).Return(nil)
	objectFile.On("DeleteFile", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	_, status := usecaseCat.DeleteStep(context.Background(), 1)
	assert.Equal(t, status, http.StatusOK)
}

func Test_DeleteStepFindErr(t *testing.T) {
	cat := entity.Steps{
		Cover: "https://google.com/dasdasdasdas",
		Model: entity.Model{
			ID: 1,
		},
	}
	repoCat := mocks.NewIRepositorySteps(t)
	objectFile := mocks.NewIGCImageRepo(t)
	usecaseCat := usecase.NewStepsUseCase(repoCat, objectFile)
	repoCat.On("FindStepOne", mock.Anything).Return(&cat, errors.New(""))
	_, status := usecaseCat.DeleteStep(context.Background(), 1)
	assert.Equal(t, status, http.StatusBadRequest)
}

func Test_DeleteStepObjectErr(t *testing.T) {
	cat := entity.Steps{
		Cover: "https://google.com/dasdasdasdas",
		Model: entity.Model{
			ID: 1,
		},
	}
	repoCat := mocks.NewIRepositorySteps(t)
	objectFile := mocks.NewIGCImageRepo(t)
	usecaseCat := usecase.NewStepsUseCase(repoCat, objectFile)
	repoCat.On("FindStepOne", mock.Anything).Return(&cat, nil)
	objectFile.On("DeleteFile", mock.Anything, mock.Anything, mock.Anything).Return(errors.New(""))
	_, status := usecaseCat.DeleteStep(context.Background(), 1)
	assert.Equal(t, status, http.StatusBadRequest)

}

func Test_DeleteStepDeleteErr(t *testing.T) {
	cat := entity.Steps{
		Cover: "https://google.com/dasdasdasdas",

		Model: entity.Model{
			ID: 1,
		},
	}
	repoCat := mocks.NewIRepositorySteps(t)
	objectFile := mocks.NewIGCImageRepo(t)
	usecaseCat := usecase.NewStepsUseCase(repoCat, objectFile)
	repoCat.On("FindStepOne", mock.Anything).Return(&cat, nil)
	objectFile.On("DeleteFile", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	repoCat.On("DeleteStep", mock.Anything).Return(errors.New(""))
	_, status := usecaseCat.DeleteStep(context.Background(), 1)
	assert.Equal(t, status, http.StatusBadRequest)
}

func Test_InsertStepOk(t *testing.T) {
	cat := entity.Steps{
		Cover: "https://google.com/dasdasdasdas",

		Model: entity.Model{
			ID: 1,
		},
	}
	repoCat := mocks.NewIRepositorySteps(t)
	objectFile := mocks.NewIGCImageRepo(t)
	usecaseCat := usecase.NewStepsUseCase(repoCat, objectFile)
	repoCat.On("InsertStep", mock.Anything).Return(&cat, nil)
	objectFile.On("SetFile", mock.Anything, mock.Anything, mock.Anything).Return("", nil)
	_, status := usecaseCat.InsertStep(context.Background(), cat)
	assert.Equal(t, status, http.StatusOK)
}

func Test_InsertStepObjectErr(t *testing.T) {
	cat := entity.Steps{
		Cover: "https://google.com/dasdasdasdas",

		Model: entity.Model{
			ID: 1,
		},
	}
	repoCat := mocks.NewIRepositorySteps(t)
	objectFile := mocks.NewIGCImageRepo(t)
	usecaseCat := usecase.NewStepsUseCase(repoCat, objectFile)
	objectFile.On("SetFile", mock.Anything, mock.Anything, mock.Anything).Return("", errors.New(""))
	_, status := usecaseCat.InsertStep(context.Background(), cat)
	assert.Equal(t, status, http.StatusBadRequest)
}

func Test_InsertStepInsertErr(t *testing.T) {
	cat := entity.Steps{
		Cover: "https://google.com/dasdasdasdas",

		Model: entity.Model{
			ID: 1,
		},
	}
	repoCat := mocks.NewIRepositorySteps(t)
	objectFile := mocks.NewIGCImageRepo(t)
	usecaseCat := usecase.NewStepsUseCase(repoCat, objectFile)
	repoCat.On("InsertStep", mock.Anything).Return(&cat, errors.New(""))
	objectFile.On("SetFile", mock.Anything, mock.Anything, mock.Anything).Return("", nil)
	_, status := usecaseCat.InsertStep(context.Background(), cat)
	assert.Equal(t, status, http.StatusBadRequest)
}
