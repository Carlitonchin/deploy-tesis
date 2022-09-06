package service

import (
	"context"

	"github.com/Carlitonchin/Backend-Tesis/model"
	"github.com/Carlitonchin/Backend-Tesis/model/apperrors"
)

type userService struct {
	UserRepository model.UserRepository
}

type USConfig struct {
	UserRepository model.UserRepository
}

func NewUserService(c *USConfig) model.UserService {
	return &userService{
		UserRepository: c.UserRepository,
	}
}

func (s *userService) GetById(ctx context.Context, id uint) (*model.User, error) {
	return s.UserRepository.FindById(ctx, id)
}

func (s *userService) SignUp(ctx context.Context, user *model.User) error {

	hashed_pass, err := HashPass(user.Password)

	if err != nil {
		return err
	}

	user.Password = hashed_pass

	err = s.UserRepository.Create(ctx, user)

	return err
}

func (s *userService) SignIn(ctx context.Context, user *model.User) error {
	u, err := s.UserRepository.FindByEmail(ctx, user.Email)

	if err != nil {
		type_error := apperrors.Authorization
		message := "email no encontrado"

		e := apperrors.NewError(type_error, message)
		return e
	}

	match, err := comparePass(user.Password, u.Password)

	if err != nil {
		type_error := apperrors.Internal
		message := "Fallo inesperado mientras se verificaba la contraseña"

		e := apperrors.NewError(type_error, message)
		return e
	}

	if !match {
		type_error := apperrors.Authorization
		message := "Contraseña invalida"

		e := apperrors.NewError(type_error, message)
		return e
	}

	*user = *u

	return nil

}

func (s *userService) GetAllUsers(ctx context.Context) ([]model.User, error) {
	return s.UserRepository.GetAllUsers(ctx)
}

func (s *userService) AddRoleToUser(ctx context.Context, user_id uint, role_id uint) error {

	err := s.UserRepository.AddRoleToUser(ctx, user_id, role_id)

	return err
}

func (s *userService) UpdateUserArea(ctx context.Context, user_id uint, area_id uint) error {
	return s.UserRepository.UpdateUserArea(ctx, user_id, area_id)
}
