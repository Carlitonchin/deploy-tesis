package handler

import (
	"fmt"
	"net/http"

	"github.com/Carlitonchin/Backend-Tesis/model"
	"github.com/Carlitonchin/Backend-Tesis/model/apperrors"
	"github.com/gin-gonic/gin"
)

func (s *Handler) me(ctx *gin.Context) {
	user, exists := ctx.Get("user")

	if !exists {
		message := "No se pudo extraer el usuario del contexto por algun motivo desconocido"
		type_error := apperrors.Internal

		err := apperrors.NewError(type_error, message)

		ctx.JSON(err.Status(), gin.H{
			"error": err,
		})

		return
	}

	// user exists in context
	user_id := user.(*model.User).ID

	u, err := s.UserService.GetById(ctx.Request.Context(), user_id)

	if err != nil {
		message := fmt.Sprintf("No fue posible encontrar al usuario con id=%v", user_id)
		type_error := apperrors.NotFound

		err := apperrors.NewError(type_error, message)

		ctx.JSON(err.Status(), gin.H{
			"error": err,
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"user": u,
	})
}
