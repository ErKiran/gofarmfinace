package api

import (
	"farmfinance/pkg/app"
	"farmfinance/pkg/codes"
	"farmfinance/pkg/utils"

	"github.com/gin-gonic/gin"
)

type FarmTypeForm struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func CreateFarmType(c *gin.Context) {
	var (
		appGin = app.Gin{C: c}
		form   FarmTypeForm
	)

	if err := app.BindAndValid(c, &form); err != nil {
		appGin.Response(codes.InvalidFormData, nil, err...)
		return
	}

	utils.Pretty("form", form)
}
