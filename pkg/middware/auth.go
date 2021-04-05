package middware

import (
	"github.com/KenLu0929/flowControlTask/config/auth"
	"github.com/KenLu0929/flowControlTask/pkg/models"

	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func VerifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")
		token := strings.Split(authorization, "Bearer ")[1]

		account, id, username, err := auth.AuthRequired(token)

		if err != nil {
			c.JSON(http.StatusUnauthorized, models.Result{
				Success: false,
				Error:   err.Error(),
			})
			c.Abort()
			return
		}

		//update token
		newToken, err := auth.SignToken(account, id, username)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.Result{
				Success: false,
				Error:   err.Error(),
			})
			c.Abort()
			return
		}

		c.Header("Authorization", newToken)
		c.Set("account", account)
		c.Set("id", id)
		c.Set("username", username)
		c.Next()
	}
}
