package user

import (
	"farmfinance/middlewares/auth"

	"github.com/gin-gonic/gin"
)

func GetUserFromContext(c *gin.Context) auth.User {
	return c.MustGet("User").(auth.User)
}
