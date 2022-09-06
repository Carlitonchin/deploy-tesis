package handler

import (
	"log"
	"net/http"

	"github.com/Carlitonchin/Backend-Tesis/model/apperrors"
	"github.com/gin-gonic/gin"
)

func (h *Handler) getAllRoles(ctx *gin.Context) {
	log.Print("entre al handler")
	roles, err := h.RoleService.GetRoles(ctx.Request.Context())
	log.Print("service layer passed")

	if err != nil {
		ctx.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"roles": roles,
	})
}

type updateRoleReq struct {
	UserId uint `json:"user_id" binding:"required"`
	RoleId uint `json:"role_id" binding:"required"`
}

func (h *Handler) updateUserRole(ctx *gin.Context) {
	var req updateRoleReq

	if ok := bindData(ctx, &req); !ok {
		return
	}

	err := h.UserService.AddRoleToUser(ctx.Request.Context(), req.UserId, req.RoleId)

	if err != nil {
		ctx.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}
