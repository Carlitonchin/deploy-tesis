package handler

import (
	"net/http"

	"github.com/Carlitonchin/Backend-Tesis/handler/handler_utils"
	"github.com/Carlitonchin/Backend-Tesis/model"
	"github.com/gin-gonic/gin"
)

type addQuestionReq struct {
	Text string `json:"text" binding:"required"`
}

func (h *Handler) addQuestion(ctx *gin.Context) {
	var req addQuestionReq
	if ok := bindData(ctx, &req); !ok {
		return
	}

	user, err := handler_utils.GetUser(ctx)
	if err != nil {
		handler_utils.SendErrorResponse(ctx, err)

		return
	}

	question := &model.Question{
		Text:      req.Text,
		UserRefer: user.ID,
	}

	question, err = h.QuestionService.AddQuestion(ctx.Request.Context(), question)

	if err != nil {
		handler_utils.SendErrorResponse(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"question": question.ID,
	})
}

type clasifyReq struct {
	QuestionId uint `json:"question_id" binding:"required"`
	AreaId     uint `json:"area_id" binding:"required"`
}

func (h *Handler) clasifyQuestion(ctx *gin.Context) {
	var req clasifyReq

	if ok := bindData(ctx, &req); !ok {
		return
	}

	err := h.QuestionService.Clasify(ctx.Request.Context(), req.QuestionId, req.AreaId)

	if err != nil {
		handler_utils.SendErrorResponse(ctx, err)
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}

type questionReq struct {
	Question_id uint `json:"question_id" binding:"required"`
}

func (h *Handler) takeQuestion(ctx *gin.Context) {
	var req questionReq
	if ok := bindData(ctx, &req); !ok {
		return
	}

	user, err := handler_utils.GetUser(ctx)

	if err != nil {
		handler_utils.SendErrorResponse(ctx, err)
		return
	}

	err = h.QuestionService.TakeQuestion(ctx.Request.Context(), user, req.Question_id)

	if err != nil {
		handler_utils.SendErrorResponse(ctx, err)
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}

type responseQuestionReq struct {
	QuestionId uint   `json:"question_id" binding:"required"`
	Response   string `json:"response" binding:"required"`
}

func (h *Handler) responseQuestion(ctx *gin.Context) {
	var req responseQuestionReq
	if ok := bindData(ctx, &req); !ok {
		return
	}

	user, err := handler_utils.GetUser(ctx)
	if err != nil {
		handler_utils.SendErrorResponse(ctx, err)
		return
	}

	err = h.QuestionService.ResponseQuestion(ctx, user, req.QuestionId, req.Response)

	if err != nil {
		handler_utils.SendErrorResponse(ctx, err)
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}

func (h *Handler) upLevel(ctx *gin.Context) {
	var req questionReq
	if ok := bindData(ctx, &req); !ok {
		return
	}

	user, err := handler_utils.GetUser(ctx)
	if err != nil {
		handler_utils.SendErrorResponse(ctx, err)
		return
	}

	err = h.QuestionService.UpLevel(ctx.Request.Context(), user, req.Question_id)

	if err != nil {
		handler_utils.SendErrorResponse(ctx, err)
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}

func (h *Handler) upToAdmin(ctx *gin.Context) {
	var req questionReq
	if ok := bindData(ctx, &req); !ok {
		return
	}

	user, err := handler_utils.GetUser(ctx)
	if err != nil {
		handler_utils.SendErrorResponse(ctx, err)
		return
	}

	err = h.QuestionService.UpToAdmin(ctx.Request.Context(), user, req.Question_id)

	if err != nil {
		handler_utils.SendErrorResponse(ctx, err)
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}
