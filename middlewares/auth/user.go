package auth

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type User struct {
	Email    string `json:"email"`
	ID       string `json:"id"`
	UserName string `json:"userName"`
	Role     string `json:"role"`
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Fucking Ass")
		var user User
		if err := c.BindHeader(&user); err != nil {
			fmt.Println("Err", err)
			return
		}
		c.Set("User", user)
		c.Next()
	}
}
