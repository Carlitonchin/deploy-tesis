package service

import (
	"context"

	"github.com/Carlitonchin/Backend-Tesis/model"
)

type areaService struct {
	repo model.AreaRepository
}

func NewAreaService(area_repo model.AreaRepository) model.AreaService {
	return &areaService{
		repo: area_repo,
	}
}

func (s *areaService) AddArea(ctx context.Context, area *model.Area) (*model.Area, error) {
	return s.repo.CreateArea(ctx, area)
}
