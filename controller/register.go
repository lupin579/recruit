package controller

import (
	"fmt"
	"os"
	"recruit/pkg/code"
	"recruit/service"
	"recruit/settings"

	"github.com/gin-gonic/gin"
)

func RegisterMsg(c *gin.Context) {
	var phoneNumber string
	c.ShouldBindJSON(&phoneNumber)
	if _, err := service.MsgSender(phoneNumber); err != nil { //短信发送
		ResponseError(c, code.ServerBusy, err)
		return
	}
	ResponseSuccess(c, code.Success, "发送成功")
}

func Register(c *gin.Context) {
	registerData := new(service.Register)
	c.ShouldBindJSON(registerData)
	registerService := new(service.Register)
	uid, err := registerService.Register()
	if err != nil {
		ResponseError(c, code.ServerBusy, err)
		return
	}
	err = os.Mkdir(fmt.Sprintf("%s/commodities_%d\n", settings.StaticPath, *uid), 0777)
	if err != nil {
		ResponseError(c, code.ServerBusy, err)
		return
	}
	ResponseSuccess(c, code.Success, "注册成功")
}
