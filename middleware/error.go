package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/towelong/egret/errors"
	customValidator "github.com/towelong/egret/pkg/validate"
	"strings"
)

func Error(ctx *gin.Context) {
	ctx.Next()

	if len(ctx.Errors) > 0 {
		e := ctx.Errors.Last()
		switch err := e.Err.(type) {
		case *errors.Error:
			ctx.JSON(err.Status, err)
		case validator.ValidationErrors:
			wrapError(ctx, err)
		case *validator.ValidationErrors:
			wrapError(ctx, *err)
		default:
			fromError := errors.Unknown
			ctx.JSON(fromError.Status, fromError)
		}
	}
}

func wrapError(ctx *gin.Context, err validator.ValidationErrors) {
	mapErrors := make(map[string]string)
	var (
		errString string
		ce        *errors.Error
	)
	for _, v := range err {
		errString = v.Translate(customValidator.Trans)
		filedName := strings.ToLower(v.StructField())
		mapErrors[filedName] = errString
	}
	ce = errors.ParamsError
	if len(err) > 1 {
		ce.Message = mapErrors
	} else {
		ce.Message = errString
	}
	ctx.JSON(ce.Status, ce)
}
