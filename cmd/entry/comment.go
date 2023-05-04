package entry

import (
	"service-routes/internal/domain/entity"
	"service-routes/internal/domain/usecase"
	"strconv"

	"github.com/labstack/echo/v4"
)

type commentEntry struct {
	commentCaseuse usecase.CommentsUseCase
}

func NewCommentEntry(commentCaseuse usecase.CommentsUseCase) *commentEntry {
	return &commentEntry{
		commentCaseuse,
	}
}

func (r *commentEntry) RegisterComment(context echo.Context) error {
	comment := context.Get("comment").(*entity.Comments)
	response, status := r.commentCaseuse.RegisterComment(*comment)
	return context.JSON(status, response)
}

func (r *commentEntry) DeleteComment(context echo.Context) error {
	idComment, _ := strconv.Atoi(context.Param("IDCOMMENT"))
	idUser, _ := strconv.Atoi(context.Param("IDUSER"))
	response, status := r.commentCaseuse.DeleteComment(int64(idComment), int64(idUser))
	return context.JSON(status, response)
}

func (r *commentEntry) FindComment(context echo.Context) error {
	idRoute, _ := strconv.Atoi(context.Param("ID"))
	response, status := r.commentCaseuse.FindComment(int64(idRoute))
	return context.JSON(status, response)
}
