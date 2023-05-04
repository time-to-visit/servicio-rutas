package usecase

import (
	"net/http"
	"service-routes/internal/domain/entity"
	objectValues "service-routes/internal/domain/object_values"
	"service-routes/internal/domain/repository"
)

type CommentsUseCase struct {
	repoComments repository.IRepositoryComments
}

func NewCommentsUseCase(repoComments repository.IRepositoryComments) CommentsUseCase {
	return CommentsUseCase{
		repoComments,
	}
}

func (e *CommentsUseCase) RegisterComment(comment entity.Comments) (interface{}, int) {
	newComment, err := e.repoComments.RegisterComment(comment)
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "hubo un problema al insertar el comentario", nil), http.StatusBadRequest
	}
	return objectValues.NewResponseWithData(http.StatusOK, "ok", "comentario insertada exitosamente", newComment), http.StatusOK
}

func (e *CommentsUseCase) DeleteComment(idComment int64, idUser int64) (interface{}, int) {
	err := e.repoComments.DeleteComment(idComment, idUser)
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "hubo un problema al eliminar el comentario", nil), http.StatusBadRequest
	}
	return objectValues.NewResponseWithData(http.StatusOK, "ok", "coemntario eliminada exitosamente", nil), http.StatusOK
}

func (e *CommentsUseCase) FindComment(idRoute int64) (interface{}, int) {
	comments, err := e.repoComments.FindComment(idRoute)
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "problemas al buscar la informacion", nil), http.StatusBadRequest
	}
	return objectValues.NewResponseWithData(http.StatusOK, "ok", "sucess", comments), http.StatusOK
}
