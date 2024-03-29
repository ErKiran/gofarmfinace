package api

import (
	"farmfinance/pkg/app"
	"farmfinance/pkg/codes"
	"farmfinance/pkg/utils"
	user "farmfinance/service/User"

	"github.com/gin-gonic/gin"
)

type FarmForm struct {
	Name     string  `json:"name" binding:"required"`
	Location string  `json:"location"`
	FarmSize float64 `json:"farmSize"`
	FarmType string  `json:"farmType"`
}

func CreateFarm(c *gin.Context) {
	var (
		appGin = app.Gin{C: c}
		form   FarmForm

		user = user.GetUserFromContext(c)
	)

	if err := app.BindAndValid(c, &form); err != nil {
		appGin.Response(codes.InvalidFormData, nil, err...)
		return
	}

	utils.Pretty("form", form)

	utils.Pretty("User", user)
}
