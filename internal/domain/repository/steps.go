package repository

import (
	"service-routes/internal/domain/entity"

	"gorm.io/gorm"
)

type IRepositorySteps interface {
	InsertStep(step entity.Steps) (*entity.Steps, error)
	FindStep(idRoute int64) (*[]entity.Steps, error)
	FindStepOne(idStep int64) (*entity.Steps, error)
	DeleteStep(idStep int64) error
}

func NewRepositorySteps(db *gorm.DB) IRepositorySteps {
	return &repositorySteps{
		db,
	}
}

type repositorySteps struct {
	db *gorm.DB
}

func (r *repositorySteps) InsertStep(step entity.Steps) (*entity.Steps, error) {
	err := r.db.Create(&step).Error
	return &step, err
}

func (r *repositorySteps) FindStep(idRoute int64) (*[]entity.Steps, error) {
	var steps []entity.Steps
	err := r.db.Where("id_route = ?", idRoute).Find(&steps).Error
	return &steps, err
}

func (r *repositorySteps) FindStepOne(idStep int64) (*entity.Steps, error) {
	var step entity.Steps
	err := r.db.First(&step, idStep).Error
	return &step, err
}

func (r *repositorySteps) DeleteStep(idStep int64) error {
	return r.db.Delete(entity.Steps{}, idStep).Error
}
