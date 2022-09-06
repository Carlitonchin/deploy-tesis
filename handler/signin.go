package handler

import (
	"log"
	"net/http"

	"github.com/Carlitonchin/Backend-Tesis/model"
	"github.com/Carlitonchin/Backend-Tesis/model/apperrors"
	"github.com/gin-gonic/gin"
)

type siginReq struct {
	Email string `json:"email" binding:"required,email"`
	Pass  string `json:"pass" binding:"required,gte=6,lte=30"`
}

func (h *Handler) signin(ctx *gin.Context) {
	var req siginReq

	if ok := bindData(ctx, &req); !ok {
		return // response returned inside of bindData function
	}

	u := &model.User{
		Email:    req.Email,
		Password: req.Pass,
	}

	err := h.UserService.SignIn(ctx.Request.Context(), u)

	if err != nil {
		log.Println("signin fail")
		ctx.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})

		return
	}

	// user singin succesfully
	tokens, err := h.TokenService.GetNewPairFromUser(ctx.Request.Context(), u, "")

	if err != nil {
		log.Println("Token generation fail")
		ctx.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"tokens": tokens,
	})

}
