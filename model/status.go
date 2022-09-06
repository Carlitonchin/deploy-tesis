package model

import "gorm.io/gorm"

type Status struct {
	gorm.Model
	Description string     `json:"description" gorm:"unique;not null"`
	Questions   []Question `json:"-" gorm:"foreignKey:StatusId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
