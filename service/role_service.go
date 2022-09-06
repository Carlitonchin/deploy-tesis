package service

import (
	"context"

	"github.com/Carlitonchin/Backend-Tesis/model"
)

type roleService struct {
	RoleRepository model.RoleRepository
}

func NewRoleService(repo model.RoleRepository) model.RoleService {
	return &roleService{
		RoleRepository: repo,
	}
}

func (s *roleService) GetRoles(ctx context.Context) ([]model.Role, error) {
	return s.RoleRepository.GetRoles(ctx)
}

func (s *roleService) GetRoleByName(ctx context.Context, role_name string) (*model.Role, error) {
	return s.RoleRepository.GetRoleByName(ctx, role_name)
}
