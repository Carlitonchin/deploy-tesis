package handler_utils

import (
	"github.com/Carlitonchin/Backend-Tesis/model"
	"github.com/Carlitonchin/Backend-Tesis/model/apperrors"
	"github.com/gin-gonic/gin"
)

func GetUser(ctx *gin.Context) (*model.User, error) {
	user, exists := ctx.Get("user")

	if !exists {
		message := "No se pudo extraer el usuario del contexto por algun motivo desconocido"
		type_error := apperrors.Internal

		err := apperrors.NewError(type_error, message)

		return nil, err
	}

	return user.(*model.User), nil
}

func SendErrorResponse(ctx *gin.Context, err error) {
	ctx.JSON(apperrors.Status(err), gin.H{
		"errror": err,
	})
}
