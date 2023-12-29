package middlewares

import (
	"errors"
	"recruit/controller"
	"recruit/pkg/code"
	"recruit/pkg/utils"

	"github.com/gin-gonic/gin"
)

func ValidateToken(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		controller.ResponseError(c, code.WithoutToken, errors.New("未携带token"))
		c.Abort()
		return
	}
	if _, err := utils.JWTValidator(token); err != nil {
		controller.ResponseError(c, code.InvalidToken, err)
		c.Abort()
		return
	}
	c.Next()
}
