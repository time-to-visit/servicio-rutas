package entity

import (
	"time"

	"gorm.io/gorm"
)

type Comments struct {
	Model
	Description string                 `gorm:"column:description;type:varchar(255);not null" json:"description" validate:"required"`
	Star        string                 `gorm:"column:star;type:varchar(255);not null" json:"star" validate:"required"`
	IDUser      int64                  `gorm:"column:user_id;type:int(11);not null" json:"user_id" `
	NameUser    string                 `gorm:"column:name_user;type:varchar(255);not null" json:"name_user"`
	State       string                 `gorm:"column:state;type:varchar(255);not null" json:"state" validate:"required"`
	IDRoutes    int64                  `gorm:"column:routes_id;type:int(11);not null" json:"routes_id" validate:"required"`
	Routes      *RoutesWithoutValidate `gorm:"joinForeignKey:routes_id;foreignKey:id;references:IDRoutes" json:"routes,omitempty"`
}

func (m Comments) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return
}

func (m Comments) BeforeUpdate(tx *gorm.DB) (err error) {
	m.UpdatedAt = time.Now()
	return
}
