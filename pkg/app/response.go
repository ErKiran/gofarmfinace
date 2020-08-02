package app

import (
	"farmfinance/pkg/codes"

	"github.com/gin-gonic/gin"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code   int         `json:"code,omitempty"`
	Msg    string      `json:"message,omitempty"`
	Errors []string    `json:"errors,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}

func (g *Gin) Response(code uint8, data interface{}, err ...string) {
	httpCode, message := codes.GetMsg(code)
	g.C.JSON(httpCode, Response{
		Msg:    message,
		Errors: err,
		Data:   data,
	})
	return
}
