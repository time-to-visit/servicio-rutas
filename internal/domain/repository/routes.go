package repository

import (
	"service-routes/internal/domain/entity"
	"service-routes/internal/domain/utils"

	"gorm.io/gorm"
)

type IRepositoryRoutes interface {
	InsertRoute(routes entity.Routes) (*entity.Routes, error)
	DelRoute(idRoute int64) error
	FindRoute(filter map[string]interface{}) (*[]entity.Routes, error)
	FindRouteOne(idRoute int64) (*entity.Routes, error)
}

func NewRepositoryRoutes(db *gorm.DB) IRepositoryRoutes {
	return &repositoryRoutes{
		db,
	}
}

type repositoryRoutes struct {
	db *gorm.DB
}

func (r *repositoryRoutes) InsertRoute(routes entity.Routes) (*entity.Routes, error) {
	err := r.db.Create(&routes).Error
	return &routes, err
}
func (r *repositoryRoutes) DelRoute(idRoute int64) error {
	err := r.db.Delete(entity.Routes{}, idRoute).Error
	return err
}
func (r *repositoryRoutes) FindRoute(filter map[string]interface{}) (*[]entity.Routes, error) {
	var route []entity.Routes
	command, request := utils.GetWhere(filter)
	err := r.db.Preload("Steps").Where(command, request...).Find(&route).Error
	return &route, err
}
func (r *repositoryRoutes) FindRouteOne(idRoute int64) (*entity.Routes, error) {
	var route entity.Routes
	err := r.db.First(&route, idRoute).Error
	return &route, err
}
