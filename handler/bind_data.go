package handler

import (
	"github.com/Carlitonchin/Backend-Tesis/model/apperrors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type invalidArgument struct {
	Field string `json:"field"`
	Value string `json:"value"`
	Tag   string `json:"tag"`
	Param string `json:"param"`
}

func bindData(ctx *gin.Context, req interface{}) bool {
	if err := ctx.ShouldBind(req); err != nil {
		var invalidArgs []invalidArgument

		if errs, ok := err.(validator.ValidationErrors); ok {
			for _, err := range errs {

				invalidArgs = append(invalidArgs, invalidArgument{
					Field: err.Field(),
					Tag:   err.Tag(),
					Param: err.Param(),
				})

			}

			message := "Request invalido, mirar 'invalidArgs'"
			type_error := apperrors.BadRequest

			err := apperrors.NewError(type_error, message)

			ctx.JSON(err.Status(), gin.H{
				"error":       err,
				"invalidArgs": invalidArgs,
			})

			return false
		}

		message := "Error de servidor"
		type_error := apperrors.Internal
		err := apperrors.NewError(type_error, message)

		ctx.JSON(err.Status(), gin.H{
			"error": err,
		})

		return false
	}

	return true
}
