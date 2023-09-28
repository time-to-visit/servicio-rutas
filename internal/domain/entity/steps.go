package entity

import (
	"time"

	"gorm.io/gorm"
)

type Steps struct {
	Model
	Name        string  `gorm:"column:name;type:varchar(255);not null" json:"name" validate:"required"`
	IdSites     int64   `gorm:"column:sites_id;type:int(11);not null" json:"sites_id" valiate:"required" `
	NameSites   string  `gorm:"column:name_sites;type:varchar(255);not null" json:"name_sites" validate:"required"`
	Cover       string  `gorm:"column:cover;type:varchar(255);not null" json:"cover" validate:"required"`
	Direction   string  `gorm:"column:direction;type:varchar(255);not null" json:"direction" validate:"required"`
	Latitud     string  `gorm:"column:latitud;type:varchar(255);not null" json:"latitud" validate:"required"`
	Longitud    string  `gorm:"column:longitud;type:varchar(255);not null" json:"longitud" validate:"required"`
	Description string  `gorm:"column:description;type:text;size:65535;not null" json:"description" validate:"required"`
	State       string  `gorm:"column:state;type:varchar(255);not null" json:"state" validate:"required"`
	IDRoutes    int64   `gorm:"column:routes_id;type:int(11);not null" json:"routes_id" validate:"required" `
	Routes      *Routes `gorm:"joinForeignKey:routes_id;foreignKey:id;references:IDRoutes" json:"routes,omitempty"`
}

type StepsWithoutValidate struct {
	Model
	Name        string `gorm:"column:name;type:varchar(255);not null" json:"name" validate:"required"`
	IdSites     int64  `gorm:"column:sites_id;type:int(11);not null" json:"sites_id" valiate:"required" `
	NameSites   string `gorm:"column:name_sites;type:varchar(255);not null" json:"name_sites" validate:"required"`
	Cover       string `gorm:"column:cover;type:varchar(255);not null" json:"cover" validate:"required"`
	Direction   string `gorm:"column:direction;type:varchar(255);not null" json:"direction" validate:"required"`
	Latitud     string `gorm:"column:latitud;type:varchar(255);not null" json:"latitud" validate:"required"`
	Longitud    string `gorm:"column:longitud;type:varchar(255);not null" json:"longitud" validate:"required"`
	Description string `gorm:"column:description;type:text;size:65535;not null" json:"description" validate:"required"`
	State       string `gorm:"column:state;type:varchar(255);not null" json:"state" validate:"required"`
	IDRoutes    int64  `gorm:"column:routes_id;type:int(11);not null" json:"routes_id" valiate:"required" `
}

func (m StepsWithoutValidate) TableName() string {
	return "steps"
}

func (m Steps) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return
}

func (m Steps) BeforeUpdate(tx *gorm.DB) (err error) {
	m.UpdatedAt = time.Now()
	return
}
