package apperrors

import (
	"errors"
	"net/http"
)

type Type string

const (
	Authorization   Type = "Autorización"
	BadRequest      Type = "Request Incorrecto"
	Conflict        Type = "Conflicto"
	Internal        Type = "Interno"
	NotFound        Type = "No encontrado"
	PayloadTooLarge Type = "Archivo demasiado grande"
	TimeOut         Type = "Error de Tiempo"
)

type Error struct {
	Type    Type   `json:"type"`
	Message string `json:"message"`
}

func (s *Error) Error() string {
	return s.Message
}

func (s *Error) Status() int {
	switch s.Type {
	case Authorization:
		return http.StatusUnauthorized
	case BadRequest:
		return http.StatusBadRequest
	case Conflict:
		return http.StatusConflict
	case Internal:
		return http.StatusInternalServerError
	case NotFound:
		return http.StatusNotFound
	case PayloadTooLarge:
		return http.StatusRequestEntityTooLarge
	case TimeOut:
		return http.StatusRequestTimeout
	default:
		return http.StatusInternalServerError
	}
}

func Status(err error) int {
	var e *Error
	if errors.As(err, &e) {
		return e.Status()
	}

	return http.StatusInternalServerError
}

func NewError(t Type, reason string) *Error {
	return &Error{
		Type:    t,
		Message: reason,
	}
}
