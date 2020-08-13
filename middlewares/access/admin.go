package access

import (
	user "farmfinance/service/User"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Message struct {
	Message string `json:"message"`
}

func OnlyAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			user = user.GetUserFromContext(c)
		)

		if user.Role != "Admin" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, Message{Message: "Only Admin access it"})
		}

		c.Next()
	}
}
