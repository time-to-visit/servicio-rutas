package handler

import (
	"service-routes/cmd/entry"
	"service-routes/internal/domain/usecase"
	"service-routes/internal/domain/validator"

	"github.com/labstack/echo/v4"
)

func NewCommentEntry(e *echo.Echo, commentUseCase usecase.CommentsUseCase, auth func(next echo.HandlerFunc) echo.HandlerFunc) *echo.Echo {
	commentEntry := entry.NewCommentEntry(commentUseCase)
	e.POST("/routes/comment", commentEntry.RegisterComment, auth, validator.ValidateComment)
	e.GET("/routes/comment/:ID", commentEntry.FindComment, auth)
	e.DELETE("/routes/comment/:IDCOMMENT/:IDUSER", commentEntry.DeleteComment, auth)
	return e
}
