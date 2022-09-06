package handler

import (
	"net/http"

	"github.com/Carlitonchin/Backend-Tesis/handler/handler_utils"
	"github.com/Carlitonchin/Backend-Tesis/model"
	"github.com/gin-gonic/gin"
)

type areaReq struct {
	Name string `json:"name" binding:"required"`
}

func (h *Handler) addArea(ctx *gin.Context) {
	var req areaReq
	if ok := bindData(ctx, &req); !ok {
		return
	}

	area := &model.Area{
		Name: req.Name,
	}

	area, err := h.AreaService.AddArea(ctx.Request.Context(), area)

	if err != nil {
		handler_utils.SendErrorResponse(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"area": area.ID,
	})
}

type updateAreaReq struct {
	UserId uint `json:"user_id" binding:"required"`
	AreaId uint `json:"area_id" binding:"required"`
}

func (h *Handler) updateUserArea(ctx *gin.Context) {
	var req updateAreaReq

	if ok := bindData(ctx, &req); !ok {
		return
	}

	err := h.UserService.UpdateUserArea(ctx.Request.Context(), req.UserId, req.AreaId)

	if err != nil {
		handler_utils.SendErrorResponse(ctx, err)
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}
