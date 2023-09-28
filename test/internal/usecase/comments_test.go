package usecase_test

import (
	"errors"
	"net/http"
	"service-routes/internal/domain/entity"
	"service-routes/internal/domain/usecase"
	"service-routes/mocks"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/mock"
)

func Test_InsertCommentsOk(t *testing.T) {
	repoCat := mocks.NewIRepositoryComments(t)
	cats := entity.Comments{}
	usecaseCat := usecase.NewCommentsUseCase(repoCat)

	repoCat.On("RegisterComment", mock.Anything).Return(&cats, nil)

	_, status := usecaseCat.RegisterComment(cats)
	assert.Equal(t, http.StatusOK, status)
}

func Test_InsertCommentsErr(t *testing.T) {
	repoCat := mocks.NewIRepositoryComments(t)
	cats := entity.Comments{}
	usecaseCat := usecase.NewCommentsUseCase(repoCat)

	repoCat.On("RegisterComment", mock.Anything).Return(&cats, errors.New(""))

	_, status := usecaseCat.RegisterComment(cats)
	assert.Equal(t, http.StatusBadRequest, status)
}

func Test_FindCommentsOk(t *testing.T) {
	repoCat := mocks.NewIRepositoryComments(t)
	usecaseCat := usecase.NewCommentsUseCase(repoCat)
	cats := []entity.Comments{}

	repoCat.On("FindComment", mock.Anything).Return(&cats, nil)

	_, status := usecaseCat.FindComment(1)
	assert.Equal(t, http.StatusOK, status)
}

func Test_FindCommentsErr(t *testing.T) {
	repoCat := mocks.NewIRepositoryComments(t)
	usecaseCat := usecase.NewCommentsUseCase(repoCat)
	cats := []entity.Comments{}

	repoCat.On("FindComment", mock.Anything).Return(&cats, errors.New(""))

	_, status := usecaseCat.FindComment(1)
	assert.Equal(t, http.StatusBadRequest, status)
}

func Test_DeleteCommentsOk(t *testing.T) {
	repoCat := mocks.NewIRepositoryComments(t)
	usecaseCat := usecase.NewCommentsUseCase(repoCat)

	repoCat.On("DeleteComment", mock.Anything, mock.Anything).Return(errors.New(""))

	_, status := usecaseCat.DeleteComment(1, 1)
	assert.Equal(t, http.StatusBadRequest, status)
}

func Test_DeleteCommentsErr(t *testing.T) {
	repoCat := mocks.NewIRepositoryComments(t)
	usecaseCat := usecase.NewCommentsUseCase(repoCat)

	repoCat.On("DeleteComment", mock.Anything, mock.Anything).Return(errors.New(""))

	_, status := usecaseCat.DeleteComment(1, 1)
	assert.Equal(t, http.StatusBadRequest, status)
}
