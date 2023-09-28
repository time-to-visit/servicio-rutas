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

func Test_DeleteResourceOk(t *testing.T) {
	cat := entity.Resource{
		UrlResource: "https://google.com/dasdasdasdas",

		Model: entity.Model{
			ID: 1,
		},
	}
	repoCat := mocks.NewIRepositoryResource(t)
	objectFile := mocks.NewIGCImageRepo(t)
	usecaseCat := usecase.NewResourcesUseCase(repoCat, objectFile)
	repoCat.On("FindOneResource", mock.Anything).Return(&cat, nil)
	repoCat.On("DelResource", mock.Anything).Return(nil)
	objectFile.On("DeleteFile", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	_, status := usecaseCat.DeleteResource(context.Background(), 1)
	assert.Equal(t, status, http.StatusOK)
}

func Test_DeleteResourceFindErr(t *testing.T) {
	cat := entity.Resource{
		UrlResource: "https://google.com/dasdasdasdas",

		Model: entity.Model{
			ID: 1,
		},
	}
	repoCat := mocks.NewIRepositoryResource(t)
	objectFile := mocks.NewIGCImageRepo(t)
	usecaseCat := usecase.NewResourcesUseCase(repoCat, objectFile)
	repoCat.On("FindOneResource", mock.Anything).Return(&cat, errors.New(""))
	_, status := usecaseCat.DeleteResource(context.Background(), 1)
	assert.Equal(t, status, http.StatusBadRequest)
}

func Test_DeleteResourceObjectErr(t *testing.T) {
	cat := entity.Resource{
		UrlResource: "https://google.com/dasdasdasdas",

		Model: entity.Model{
			ID: 1,
		},
	}
	repoCat := mocks.NewIRepositoryResource(t)
	objectFile := mocks.NewIGCImageRepo(t)
	usecaseCat := usecase.NewResourcesUseCase(repoCat, objectFile)
	repoCat.On("FindOneResource", mock.Anything).Return(&cat, nil)
	objectFile.On("DeleteFile", mock.Anything, mock.Anything, mock.Anything).Return(errors.New(""))
	_, status := usecaseCat.DeleteResource(context.Background(), 1)
	assert.Equal(t, status, http.StatusBadRequest)

}

func Test_DeleteResourceDeleteErr(t *testing.T) {
	cat := entity.Resource{
		UrlResource: "https://google.com/dasdasdasdas",

		Model: entity.Model{
			ID: 1,
		},
	}
	repoCat := mocks.NewIRepositoryResource(t)
	objectFile := mocks.NewIGCImageRepo(t)
	usecaseCat := usecase.NewResourcesUseCase(repoCat, objectFile)
	repoCat.On("FindOneResource", mock.Anything).Return(&cat, nil)
	objectFile.On("DeleteFile", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	repoCat.On("DelResource", mock.Anything).Return(errors.New(""))
	_, status := usecaseCat.DeleteResource(context.Background(), 1)
	assert.Equal(t, status, http.StatusBadRequest)
}

func Test_InsertResourceOk(t *testing.T) {
	cat := entity.Resource{
		UrlResource: "https://google.com/dasdasdasdas",

		Model: entity.Model{
			ID: 1,
		},
	}
	repoCat := mocks.NewIRepositoryResource(t)
	objectFile := mocks.NewIGCImageRepo(t)
	usecaseCat := usecase.NewResourcesUseCase(repoCat, objectFile)
	repoCat.On("AddResource", mock.Anything).Return(&cat, nil)
	objectFile.On("SetFile", mock.Anything, mock.Anything, mock.Anything).Return("", nil)
	_, status := usecaseCat.InsertResource(context.Background(), cat)
	assert.Equal(t, status, http.StatusOK)
}

func Test_InsertResourceObjectErr(t *testing.T) {
	cat := entity.Resource{
		UrlResource: "https://google.com/dasdasdasdas",

		Model: entity.Model{
			ID: 1,
		},
	}
	repoCat := mocks.NewIRepositoryResource(t)
	objectFile := mocks.NewIGCImageRepo(t)
	usecaseCat := usecase.NewResourcesUseCase(repoCat, objectFile)
	objectFile.On("SetFile", mock.Anything, mock.Anything, mock.Anything).Return("", errors.New(""))
	_, status := usecaseCat.InsertResource(context.Background(), cat)
	assert.Equal(t, status, http.StatusBadRequest)
}

func Test_InsertResourceInsertErr(t *testing.T) {
	cat := entity.Resource{
		UrlResource: "https://google.com/dasdasdasdas",

		Model: entity.Model{
			ID: 1,
		},
	}
	repoCat := mocks.NewIRepositoryResource(t)
	objectFile := mocks.NewIGCImageRepo(t)
	usecaseCat := usecase.NewResourcesUseCase(repoCat, objectFile)
	repoCat.On("AddResource", mock.Anything).Return(&cat, errors.New(""))
	objectFile.On("SetFile", mock.Anything, mock.Anything, mock.Anything).Return("", nil)
	_, status := usecaseCat.InsertResource(context.Background(), cat)
	assert.Equal(t, status, http.StatusBadRequest)
}
