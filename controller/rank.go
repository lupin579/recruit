package controller

import (
	"github.com/gin-gonic/gin"
	"recruit/pkg/code"
	"recruit/service"
)

func Rank(c *gin.Context) {
	rankService := new(service.RankService)
	list, err := rankService.Rank()
	if err != nil {
		ResponseError(c, code.GetError, err)
	}
	ResponseSuccess(c, code.Success, list)
}
