package model

import "gorm.io/gorm"

type Area struct {
	gorm.Model
	Name      string     `json:"name" gorm:"unique;not null"`
	Users     []User     `json:"-" gorm:"foreignKey:AreaID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Questions []Question `json:"-" gorm:"foreignKey:AreaID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
