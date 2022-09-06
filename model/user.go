package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email               string     `json:"email" gorm:"unique;not null"`
	Name                string     `json:"name" gorm:"unique;not null"`
	Password            string     `json:"-" gorm:"not null"`
	RoleID              uint       `json:"role_id"`
	Role                *Role      `json:"role" gorm:"-"`
	AreaID              *uint      `json:"area_id"`
	Area                *Area      `json:"area" gorm:"-"`
	Questions           []Question `json:"-" gorm:"foreignKey:UserRefer"`
	QuestionsToResponse []Question `json:"-" gorm:"foreignKey:UserResponsible"`
}
