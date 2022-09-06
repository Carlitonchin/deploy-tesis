package service

import (
	"context"
	"fmt"
	"os"

	"github.com/Carlitonchin/Backend-Tesis/model"
	"github.com/Carlitonchin/Backend-Tesis/model/apperrors"
	"github.com/Carlitonchin/Backend-Tesis/some_utils"
)

type questionService struct {
	repo model.QuestionRepository
}

func NewQuestionService(question_repo model.QuestionRepository) model.QuestionService {
	return &questionService{
		repo: question_repo,
	}
}

func (s *questionService) AddQuestion(
	ctx context.Context, question *model.Question) (
	*model.Question, error) {

	status_send_id, err := some_utils.GetUintEnv("STATUS_SEND_CODE")
	if err != nil {
		type_error := apperrors.Internal
		message := "No se pudo leer el id del stado 'enviada'"

		e := apperrors.NewError(type_error, message)
		return nil, e
	}

	question.StatusId = status_send_id
	return s.repo.CreateQuestion(ctx, question)
}

func (s *questionService) Clasify(ctx context.Context, question_id uint, area_id uint) error {
	question, err := s.repo.GetById(ctx, question_id)
	if err != nil {
		return err
	}

	status_sended_id, err := some_utils.GetUintEnv("STATUS_SEND_CODE")
	if err != nil {
		return err
	}

	if question.StatusId != status_sended_id {
		type_error := apperrors.Conflict
		message := fmt.Sprintf("Solo se pueden clasificar de esa manera las preguntas con status_id = '%v'", status_sended_id)

		e := apperrors.NewError(type_error, message)
		return e
	}

	return s.repo.Clasify(ctx, question_id, area_id)
}

func (s *questionService) TakeQuestion(ctx context.Context, user *model.User, question_id uint) error {
	status_clasified_id, err := some_utils.GetUintEnv("STATUS_CLASIFIED1_CODE")

	if err != nil {
		return err
	}

	status_claisfied_id2, err := some_utils.GetUintEnv("STATUS_CLASIFIED2_CODE")

	if err != nil {
		return err
	}

	role_admin := os.Getenv("ROLE_ADMIN")
	role_specialist1 := os.Getenv("ROLE_SPECIALIST_LEVEL1")
	role_specialist2 := os.Getenv("ROLE_SPECIALIST_LEVEL2")

	question, err := s.repo.GetById(ctx, question_id)
	if err != nil {
		return err
	}

	can_take_specialist1 := user.Role.Name == role_specialist1 && question.StatusId != status_clasified_id
	can_take_specialist2 := user.Role.Name == role_specialist2 && question.StatusId != status_claisfied_id2

	if !can_take_specialist1 && !can_take_specialist2 {
		type_error := apperrors.Conflict
		message := fmt.Sprintf("%v no pueden tomar preguntas con status_id = '%v'", user.Role.Name, question.StatusId)
		err = apperrors.NewError(type_error, message)
		return err
	}

	if question.UserResponsible != nil {
		type_error := apperrors.Conflict
		message := fmt.Sprintf("La pregunta con id = '%v' ya fue tomada por alguien", question_id)
		err = apperrors.NewError(type_error, message)
		return err
	}

	if user.Role.Name != role_admin && (user.AreaID == nil || *user.AreaID != *question.AreaID) {
		type_error := apperrors.Conflict
		message := fmt.Sprintf("La pregunta con id = '%v' solo puede ser tomada por usuarios del area de id = '%v'", question_id, *question.AreaID)
		err = apperrors.NewError(type_error, message)
		return err
	}

	return s.repo.TakeQuestion(ctx, user.ID, question_id)
}

func (s *questionService) ResponseQuestion(ctx context.Context, user *model.User, question_id uint, response string) error {
	question, err := s.repo.GetById(ctx, question_id)

	if err != nil {
		return err
	}

	if *question.UserResponsible != user.ID {
		type_error := apperrors.Conflict
		message := fmt.Sprintf("El usuario con id = '%v' no es el responsable de la pregunta con id = '%v'", user.ID, question_id)
		err = apperrors.NewError(type_error, message)
		return err
	}

	return s.repo.ResponseQuestion(ctx, question_id, response)
}

func (s *questionService) UpLevel(ctx context.Context, user *model.User, question_id uint) error {
	question, err := s.repo.GetById(ctx, question_id)

	if err != nil {
		return err
	}

	status_clasified_id, err := some_utils.GetUintEnv("STATUS_CLASIFIED1_CODE")
	if err != nil {
		return err
	}

	if question.StatusId != status_clasified_id {
		type_error := apperrors.Conflict
		message := fmt.Sprintf("No se puede subir de nivel una pregunta con status_id = '%v'", question.StatusId)

		err = apperrors.NewError(type_error, message)
		return err
	}

	if *question.UserResponsible != user.ID {
		type_error := apperrors.Authorization
		message := fmt.Sprintf("El usuario con id = '%v' no es el responsable de la pregunta con id = '%v'", user.ID, question_id)
		err = apperrors.NewError(type_error, message)
		return err
	}

	return s.repo.UpLevel(ctx, question_id)
}

func (s *questionService) UpToAdmin(ctx context.Context, user *model.User, question_id uint) error {
	question, err := s.repo.GetById(ctx, question_id)

	if err != nil {
		return err
	}

	status_clasified2_id, err := some_utils.GetUintEnv("STATUS_CLASIFIED2_CODE")
	if err != nil {
		return err
	}

	if question.StatusId != status_clasified2_id {
		type_error := apperrors.Conflict
		message := fmt.Sprintf(
			"Las preguntas con status_id = '%v' no pueden ser elevadas al admin", question.StatusId)

		err = apperrors.NewError(type_error, message)
		return err
	}

	if question.UserResponsible == nil || *question.UserResponsible != user.ID {
		type_error := apperrors.Conflict
		message := fmt.Sprintf(
			"El usuario con id = '%v' no es responsable de la pregunta con id = '%v'",
			user.ID, question_id)

		err = apperrors.NewError(type_error, message)
		return err
	}

	return s.repo.UpToAdmin(ctx, question_id)
}
