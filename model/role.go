package model

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name  string `json:"name" gorm:"unique;not null"`
	Users []User `json:"-" gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
