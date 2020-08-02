package app

import (
	"encoding/json"
	"farmfinance/pkg/codes"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	ut "github.com/go-playground/universal-translator"
)

var trans ut.Translator

func BindAndValid(c *gin.Context, form interface{}) []string {
	err := c.ShouldBind(form)
	if err != nil {
		var errors []string
		if reflect.TypeOf(err) == reflect.TypeOf(&json.UnmarshalTypeError{}) {
			errors = append(errors, "Form parsing error")
			return errors
		}
		for _, err := range err.(validator.ValidationErrors).Translate(trans) {
			errors = append(errors, err)
		}
		return errors
	}
	return nil
}

func BindParam(c *gin.Context, params interface{}) uint8 {
	if err := c.ShouldBindUri(params); err != nil {
		return codes.InvalidURIParam
	}
	return codes.Ok
}
