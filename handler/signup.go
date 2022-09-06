package handler

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Carlitonchin/Backend-Tesis/model"
	"github.com/Carlitonchin/Backend-Tesis/model/apperrors"
	"github.com/gin-gonic/gin"
)

type signup_req struct {
	Email    string `json:"email" binding:"required,email"`
	Pass     string `json:"pass" binding:"required,gte=6,lte=30"`
	UserName string `json:"name" binding:"required"`
	Worker   string `json:"worker" binding:"required"`
}

func (s *Handler) signUp(ctx *gin.Context) {
	var req signup_req

	if ok := bindData(ctx, &req); !ok {
		return // error handled at bindData function
	}

	u := &model.User{
		Email:    req.Email,
		Password: req.Pass,
		Name:     req.UserName,
	}

	default_role_name := os.Getenv("ROLE_DEFAULT_WORKER")

	if req.Worker == "0" {
		default_role_name = os.Getenv("ROLE_DEFAULT_STUDENT")
	}

	role, err := s.RoleService.GetRoleByName(ctx.Request.Context(), default_role_name)

	if err != nil {
		type_error := apperrors.Internal
		message := fmt.Sprintf("Error buscando el rol %s en la base de datos", default_role_name)

		e := apperrors.NewError(type_error, message)

		ctx.JSON(e.Status(), gin.H{
			"error": e,
		})

		return
	}

	u.RoleID = role.ID
	u.Role = &model.Role{
		Name: role.Name,
	}

	err = s.UserService.SignUp(ctx.Request.Context(), u)

	if err != nil {
		ctx.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})

		return
	}

	token, err := s.TokenService.GetNewPairFromUser(ctx.Request.Context(), u, "")

	if err != nil {
		ctx.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})

		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"tokens": token,
	})
}
