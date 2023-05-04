package repository

import (
	"service-routes/internal/domain/entity"

	"gorm.io/gorm"
)

type IRepositoryComments interface {
	RegisterComment(comment entity.Comments) (*entity.Comments, error)
	DeleteComment(idComment int64, idUser int64) error
	FindComment(idRoute int64) (*[]entity.Comments, error)
}

func NewRepositoryComments(db *gorm.DB) IRepositoryComments {
	return &repositoryComments{
		db,
	}
}

type repositoryComments struct {
	db *gorm.DB
}

func (r *repositoryComments) RegisterComment(comment entity.Comments) (*entity.Comments, error) {
	err := r.db.Create(&comment).Error
	return &comment, err
}

func (r *repositoryComments) DeleteComment(idComment int64, idUser int64) error {
	return r.db.Where("id = ? and user_id = ?", idComment, idUser).Error
}

func (r *repositoryComments) FindComment(idRoute int64) (*[]entity.Comments, error) {
	var comments []entity.Comments
	err := r.db.Where("routes_id = ?", idRoute).Find(&comments).Error
	return &comments, err
}
