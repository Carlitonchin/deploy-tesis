package repository

import (
	"context"
	"fmt"

	"github.com/Carlitonchin/Backend-Tesis/model"
	"github.com/Carlitonchin/Backend-Tesis/model/apperrors"
	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) model.UserRepository {

	return &userRepository{
		DB: db,
	}
}

func (s *userRepository) Create(ctx context.Context, user *model.User) error {
	result := s.DB.Create(user)

	return result.Error
}

func (s *userRepository) FindById(ctx context.Context, id uint) (*model.User, error) {
	var user model.User

	err := s.DB.First(&user, id).Error

	if err != nil {
		return nil, err
	}

	role, err := s.getRoleById(ctx, user.RoleID)

	if err != nil {
		return nil, err
	}

	user.Role = &model.Role{
		Name: role.Name,
	}

	return &user, nil
}

func (s *userRepository) getRoleById(ctx context.Context, id uint) (*model.Role, error) {
	var role *model.Role

	err := s.DB.First(&role, id).Error

	if err != nil {
		return nil, err
	}

	return role, err
}

func (s *userRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	tx := s.DB.Where("email = ?", email).First(&user)

	if tx.Error != nil {
		return nil, tx.Error
	}

	role, err := s.getRoleById(ctx, user.RoleID)

	if err != nil {
		return nil, err
	}

	user.Role = &model.Role{
		Name: role.Name,
	}

	return &user, nil
}

func (s *userRepository) GetAllUsers(ctx context.Context) ([]model.User, error) {
	var users []model.User

	rows, err := s.DB.Table("users").Select("users.id, users.name, users.email, users.role_id, roles.name").Joins("left join roles on users.role_id = roles.id").Rows()
	for rows.Next() {
		var user_id uint
		var user_name string
		var user_email string

		var role_id uint
		var role_name string

		rows.Scan(&user_id, &user_name, &user_email, &role_id, &role_name)
		role := &model.Role{
			Name: role_name,
		}
		role.ID = role_id

		u := model.User{
			Email:  user_email,
			Name:   user_name,
			RoleID: role_id,
			Role:   role,
		}
		u.ID = user_id
		users = append(users, u)
	}

	return users, err
}

func (s *userRepository) AddRoleToUser(ctx context.Context, user_id uint, role_id uint) error {
	err := s.DB.First(&model.Role{}, role_id).Error

	if err != nil {
		type_error := apperrors.Conflict
		message := fmt.Sprintf("El rol con identificador '%v' no existe", role_id)

		e := apperrors.NewError(type_error, message)
		return e
	}

	err = s.DB.Model(&model.User{}).Where("id = ?", user_id).Update("role_id", role_id).Error

	if err != nil {
		type_error := apperrors.Conflict
		message := fmt.Sprintf("El usuario con identificador '%v' no existe", user_id)

		e := apperrors.NewError(type_error, message)
		return e
	}

	return nil
}

func (s *userRepository) UpdateUserArea(ctx context.Context, user_id uint, area_id uint) error {
	err := s.DB.Model(&model.User{}).Where("id = ?", user_id).Update("area_id", area_id).Error

	if err != nil {
		type_error := apperrors.Conflict
		message := fmt.Sprintf("No se pudo encontrar el usuario con id %v o el area con id %v", user_id, area_id)
		err = apperrors.NewError(type_error, message)
	}

	return err
}
