package model

import "gorm.io/gorm"

type Question struct {
	gorm.Model
	Text            string  `json:"text" gorm:"not null"`
	UserRefer       uint    `json:"user_id"`
	StatusId        uint    `json:"status_id"`
	Status          *Status `json:"status" gorm:"-"`
	AreaID          *uint   `json:"area_id"`
	UserResponsible *uint   `json:"user_responsible_id"`
	Response        string  `json:"response"`
}
