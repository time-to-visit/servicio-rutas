package entity

import (
	"time"

	"gorm.io/gorm"
)

type Resource struct {
	Model
	UrlResource string                `gorm:"column:url_resource;type:varchar(255);not null" json:"url_resource" validate:"required"`
	State       string                `gorm:"column:state;type:varchar(255);not null" json:"state" validate:"required"`
	IDStep      int64                 `gorm:"column:step_id;type:int(11);not null" json:"step_id" valiate:"required" `
	Steps       *StepsWithoutValidate `gorm:"joinForeignKey:step_id;foreignKey:id;references:IDStep" json:"steps,omitempty"`
	IDUser      int64                 `gorm:"column:user_id;type:int(11);not null" json:"user_id" valiate:"required" `
}

func (m Resource) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return
}

func (m Resource) BeforeUpdate(tx *gorm.DB) (err error) {
	m.UpdatedAt = time.Now()
	return
}
