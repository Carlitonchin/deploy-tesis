package handler

import (
	"net/http"

	"github.com/Carlitonchin/Backend-Tesis/model"
	"github.com/Carlitonchin/Backend-Tesis/model/apperrors"
	"github.com/gin-gonic/gin"
)

func (h *Handler) signout(ctx *gin.Context) {
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

	err := h.TokenService.SignOut(ctx, user.(*model.User).ID)
	if err != nil {
		ctx.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Sesion cerrada correctamente",
	})
}
