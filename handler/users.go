package handler

import (
	"net/http"

	"github.com/Carlitonchin/Backend-Tesis/model/apperrors"
	"github.com/gin-gonic/gin"
)

func (h *Handler) getAllUsers(ctx *gin.Context) {
	users, err := h.UserService.GetAllUsers(ctx)

	if err != nil {
		ctx.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}
