package repository

import (
	"context"
	"fmt"

	"github.com/Carlitonchin/Backend-Tesis/model"
	"github.com/Carlitonchin/Backend-Tesis/model/apperrors"
	"gorm.io/gorm"
)

type areaRepository struct {
	DB *gorm.DB
}

func NewAreaRepository(db *gorm.DB) model.AreaRepository {
	return &areaRepository{
		DB: db,
	}
}

func (s *areaRepository) CreateArea(ctx context.Context, area *model.Area) (*model.Area, error) {
	err := s.DB.Create(area).Error

	if err != nil {
		type_error := apperrors.Conflict
		message := fmt.Sprintf("Ya existe un area con nombre: '%v'", area.Name)

		err = apperrors.NewError(type_error, message)
	}

	return area, err
}
