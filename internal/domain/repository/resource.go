package repository

import (
	"service-routes/internal/domain/entity"

	"gorm.io/gorm"
)

type IRepositoryResource interface {
	AddResource(resource entity.Resource) (*entity.Resource, error)
	DelResource(idResource int64) error
	FindOneResource(idResource int64) (*entity.Resource, error)
}

func NewRepositoryResource(db *gorm.DB) IRepositoryResource {
	return &repositoryResource{
		db,
	}
}

type repositoryResource struct {
	db *gorm.DB
}

func (r *repositoryResource) AddResource(resource entity.Resource) (*entity.Resource, error) {
	err := r.db.Create(&resource).Error
	return &resource, err
}

func (r *repositoryResource) FindOneResource(idResource int64) (*entity.Resource, error) {
	var resource entity.Resource
	err := r.db.First(&resource, idResource).Error
	return &resource, err
}

func (r *repositoryResource) DelResource(idResource int64) error {
	return r.db.Delete(entity.Resource{}, idResource).Error
}
