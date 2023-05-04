package entity

import (
	"time"

	"gorm.io/gorm"
)

type Routes struct {
	Model
	MunicipalitiesId   int64  `gorm:"column:municipalities_id;type:int(11);not null" json:"municipalities_id" validate:"required" `
	MuincipalitiesName string `gorm:"column:municipalities_name;type:varchar(255);not null" json:"municipalities_name" validate:"required"`
	DepartmentId       int64  `gorm:"column:department_id;type:int(11);not null" json:"department_id" validate:"required" `
	DepartmentName     string `gorm:"column:department_name;type:varchar(255);not null" json:"department_name" validate:"required"`
	Description        string `gorm:"column:description;type:varchar(255);not null" json:"description" validate:"required"`
	Name               string `gorm:"column:name;type:varchar(255);not null" json:"name" validate:"required"`
	Cover              string `gorm:"column:cover;type:varchar(255);not null" json:"cover" validate:"required"`
	State              string `gorm:"column:state;type:varchar(255);not null" json:"state" validate:"required"`
}

type RoutesWithoutValidate struct {
	Model
	MunicipalitiesId   int64  `gorm:"column:municipalities_id;type:int(11);not null" json:"municipalities_id" `
	MuincipalitiesName string `gorm:"column:municipalities_name;type:varchar(255);not null" json:"municipalities_name"`
	DepartmentId       int64  `gorm:"column:department_id;type:int(11);not null" json:"department_id" `
	DepartmentName     string `gorm:"column:department_name;type:varchar(255);not null" json:"department_name"`
	Description        string `gorm:"column:description;type:varchar(255);not null" json:"description" validate:"required"`
	Name               string `gorm:"column:name;type:varchar(255);not null" json:"name"`
	Cover              string `gorm:"column:cover;type:varchar(255);not null" json:"cover"`
	State              string `gorm:"column:state;type:varchar(255);not null" json:"state"`
}

func (m RoutesWithoutValidate) TableName() string {
	return "routes"
}

func (m Routes) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return
}

func (m Routes) BeforeUpdate(tx *gorm.DB) (err error) {
	m.UpdatedAt = time.Now()
	return
}
