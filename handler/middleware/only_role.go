package middleware

import (
	"os"

	"github.com/Carlitonchin/Backend-Tesis/handler/handler_utils"
	"github.com/Carlitonchin/Backend-Tesis/model/apperrors"
	"github.com/gin-gonic/gin"
)

func OnlyRoles(roles []string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, err := handler_utils.GetUser(ctx)
		if err != nil {
			ctx.JSON(apperrors.Status(err), gin.H{
				"error": err,
			})
			ctx.Abort()
			return
		}

		fail := true

		for _, role := range roles {
			if user.Role.Name == os.Getenv(role) {
				fail = false
			}
		}

		if fail {
			type_error := apperrors.Authorization
			message := "No tiene permisos para acceder a este recurso"

			e := apperrors.NewError(type_error, message)
			ctx.JSON(e.Status(), gin.H{
				"error": e,
			})

			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
