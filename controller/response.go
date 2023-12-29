package controller

import (
	"log"
	"recruit/pkg/code"
	"runtime"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func ResponseSuccess(ctx *gin.Context, errCode int, data interface{}) {
	ctx.JSON(200, Response{
		Code:    errCode,
		Message: code.Msg(errCode),
		Data:    data,
	})
}

func ResponseError(ctx *gin.Context, errCode int, err error) {
	var callerName string
	pc, _, _, ok := runtime.Caller(1)
	if ok {
		fc := runtime.FuncForPC(pc)
		callerName = fc.Name()
	} else {
		callerName = "notFound"
	}
	log.Printf("%s : %s\n", callerName, err.Error())
	ctx.JSON(500, Response{
		Code:    errCode,
		Message: code.Msg(errCode),
		Data:    nil,
	})
}
