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

func Test_SeeRoutesOk(t *testing.T) {
	repoCat := mocks.NewIRepositoryRoutes(t)
	cats := []entity.Routes{
		{},
	}
	objectFile := mocks.NewIGCImageRepo(t)
	usecaseCat := usecase.NewRoutesUseCase(repoCat, objectFile)

	repoCat.On("FindRoute", mock.Anything).Return(&cats, nil)

	_, status := usecaseCat.FindRoute(map[string]interface{}{})
	assert.Equal(t, http.StatusOK, status)
}

func Test_SeeRoutesErr(t *testing.T) {
	repoCat := mocks.NewIRepositoryRoutes(t)
	cats := []entity.Routes{
		{},
	}
	objectFile := mocks.NewIGCImageRepo(t)
	usecaseCat := usecase.NewRoutesUseCase(repoCat, objectFile)

	repoCat.On("FindRoute", mock.Anything).Return(&cats, errors.New("err"))

	_, status := usecaseCat.FindRoute(map[string]interface{}{})
	assert.Equal(t, http.StatusBadRequest, status)
}

func Test_SeeRoutesOneOk(t *testing.T) {
	repoCat := mocks.NewIRepositoryRoutes(t)
	cats := entity.Routes{}

	objectFile := mocks.NewIGCImageRepo(t)
	usecaseCat := usecase.NewRoutesUseCase(repoCat, objectFile)

	repoCat.On("FindRouteOne", mock.Anything).Return(&cats, nil)

	_, status := usecaseCat.FindRouteOne(1)
	assert.Equal(t, http.StatusOK, status)
}

func Test_SeeRoutesOneErr(t *testing.T) {
	repoCat := mocks.NewIRepositoryRoutes(t)
	cats := entity.Routes{}
	objectFile := mocks.NewIGCImageRepo(t)
	usecaseCat := usecase.NewRoutesUseCase(repoCat, objectFile)

	repoCat.On("FindRouteOne", mock.Anything).Return(&cats, errors.New("err"))

	_, status := usecaseCat.FindRouteOne(1)
	assert.Equal(t, http.StatusBadRequest, status)
}

func Test_DeleteRoutesOk(t *testing.T) {
	cat := entity.Routes{
		Cover: "https://google.com/dasdasdasdas",
		Model: entity.Model{
			ID: 1,
		},
	}
	repoCat := mocks.NewIRepositoryRoutes(t)
	objectFile := mocks.NewIGCImageRepo(t)
	usecaseCat := usecase.NewRoutesUseCase(repoCat, objectFile)
	repoCat.On("FindRouteOne", mock.Anything).Return(&cat, nil)
	repoCat.On("DelRoute", mock.Anything).Return(nil)
	objectFile.On("DeleteFile", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	_, status := usecaseCat.DeleteRoute(context.Background(), 1)
	assert.Equal(t, status, http.StatusOK)
}

func Test_DeleteRoutesFindErr(t *testing.T) {
	cat := entity.Routes{
		Cover: "https://google.com/dasdasdasdas",
		Model: entity.Model{
			ID: 1,
		},
	}
	repoCat := mocks.NewIRepositoryRoutes(t)
	objectFile := mocks.NewIGCImageRepo(t)
	usecaseCat := usecase.NewRoutesUseCase(repoCat, objectFile)
	repoCat.On("FindRouteOne", mock.Anything).Return(&cat, errors.New(""))
	_, status := usecaseCat.DeleteRoute(context.Background(), 1)
	assert.Equal(t, status, http.StatusBadRequest)
}

func Test_DeleteRoutesObjectErr(t *testing.T) {
	cat := entity.Routes{
		Cover: "https://google.com/dasdasdasdas",
		Model: entity.Model{
			ID: 1,
		},
	}
	repoCat := mocks.NewIRepositoryRoutes(t)
	objectFile := mocks.NewIGCImageRepo(t)
	usecaseCat := usecase.NewRoutesUseCase(repoCat, objectFile)
	repoCat.On("FindRouteOne", mock.Anything).Return(&cat, nil)
	objectFile.On("DeleteFile", mock.Anything, mock.Anything, mock.Anything).Return(errors.New(""))
	_, status := usecaseCat.DeleteRoute(context.Background(), 1)
	assert.Equal(t, status, http.StatusBadRequest)

}

func Test_DeleteRoutesDeleteErr(t *testing.T) {
	cat := entity.Routes{
		Cover: "https://google.com/dasdasdasdas",

		Model: entity.Model{
			ID: 1,
		},
	}
	repoCat := mocks.NewIRepositoryRoutes(t)
	objectFile := mocks.NewIGCImageRepo(t)
	usecaseCat := usecase.NewRoutesUseCase(repoCat, objectFile)
	repoCat.On("FindRouteOne", mock.Anything).Return(&cat, nil)
	objectFile.On("DeleteFile", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	repoCat.On("DelRoute", mock.Anything).Return(errors.New(""))
	_, status := usecaseCat.DeleteRoute(context.Background(), 1)
	assert.Equal(t, status, http.StatusBadRequest)
}

func Test_InsertRoutesOk(t *testing.T) {
	cat := entity.Routes{
		Cover: "https://google.com/dasdasdasdas",

		Model: entity.Model{
			ID: 1,
		},
	}
	repoCat := mocks.NewIRepositoryRoutes(t)
	objectFile := mocks.NewIGCImageRepo(t)
	usecaseCat := usecase.NewRoutesUseCase(repoCat, objectFile)
	repoCat.On("InsertRoute", mock.Anything).Return(&cat, nil)
	objectFile.On("SetFile", mock.Anything, mock.Anything, mock.Anything).Return("", nil)
	_, status := usecaseCat.RegisterRoute(context.Background(), cat)
	assert.Equal(t, status, http.StatusOK)
}

func Test_InsertRoutesObjectErr(t *testing.T) {
	cat := entity.Routes{
		Cover: "https://google.com/dasdasdasdas",

		Model: entity.Model{
			ID: 1,
		},
	}
	repoCat := mocks.NewIRepositoryRoutes(t)
	objectFile := mocks.NewIGCImageRepo(t)
	usecaseCat := usecase.NewRoutesUseCase(repoCat, objectFile)
	objectFile.On("SetFile", mock.Anything, mock.Anything, mock.Anything).Return("", errors.New(""))
	_, status := usecaseCat.RegisterRoute(context.Background(), cat)
	assert.Equal(t, status, http.StatusBadRequest)
}

func Test_InsertRoutesInsertErr(t *testing.T) {
	cat := entity.Routes{
		Cover: "https://google.com/dasdasdasdas",

		Model: entity.Model{
			ID: 1,
		},
	}
	repoCat := mocks.NewIRepositoryRoutes(t)
	objectFile := mocks.NewIGCImageRepo(t)
	usecaseCat := usecase.NewRoutesUseCase(repoCat, objectFile)
	repoCat.On("InsertRoute", mock.Anything).Return(&cat, errors.New(""))
	objectFile.On("SetFile", mock.Anything, mock.Anything, mock.Anything).Return("", nil)
	_, status := usecaseCat.RegisterRoute(context.Background(), cat)
	assert.Equal(t, status, http.StatusBadRequest)
}
