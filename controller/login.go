package controller

import (
	"recruit/pkg/code"
	"recruit/pkg/utils"
	"recruit/service"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	loginService := new(service.Login)
	if err := c.ShouldBindJSON(loginService); err != nil {
		ResponseError(c, code.ServerBusy, err)
		return
	}
	userID, err := loginService.MulMethodsLogin()
	if err != nil {
		ResponseError(c, code.WrongPassWord, err)
		return
	}
	token, err := utils.JWTIssue(userID)
	if err != nil {
		ResponseError(c, code.OperateFail, err)
	}
	c.Header("Authorization", token)
	ResponseSuccess(c, code.Success, userID)
}
