package controller

import (
	"recruit/model"
	"recruit/pkg/code"
	"recruit/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func JobSeekPage(c *gin.Context) {
	sid := c.Param("uid")
	uid, _ := strconv.ParseUint(sid, 10, 64)
	jobSeekPage := service.JobSeekPage{
		PubId: uint16(uid),
	}
	posts, err := jobSeekPage.JobSeekPosts()
	if err != nil {
		ResponseError(c, code.ServerBusy, err)
	}
	ResponseSuccess(c, code.Success, posts)
}

func JobSeekDetails(c *gin.Context) {
	jsId := c.Param("jsId")
	jobSeekDetailsService := new(service.JobSeekDetailsService)
	u16, _ := strconv.ParseUint(jsId, 10, 16)
	jobSeekDetailsService.JsId = uint16(u16)
	jobPost, err := jobSeekDetailsService.JobSeekDetails()
	if err != nil {
		ResponseError(c, code.GetError, err)
	}
	ResponseSuccess(c, code.Success, jobPost)
}

func PostJobSeek(c *gin.Context) {
	sid := c.Param("uid")
	uid, _ := strconv.ParseUint(sid, 10, 16)
	jobSeek := new(model.JobSeek)
	if err := c.ShouldBind(jobSeek); err != nil {
		ResponseError(c, code.BindError, err)
	}
	postedJobSeekService := new(service.PostJobSeekService)
	if err := postedJobSeekService.PostJobSeekService(jobSeek, uint16(uid)); err != nil {
		ResponseError(c, code.OperateFail, err)
	}
	ResponseSuccess(c, code.Success, "发布成功！")
}

func UpdateJsStatus(c *gin.Context) {
	modifiedJobPost := new(model.JobSeek)
	if err := c.ShouldBindJSON(modifiedJobPost); err != nil {
		ResponseError(c, code.BindError, err)
		return
	}
	updateJobSeekService := new(service.UpdateJobSeekService)
	updateJobSeekService.Post = modifiedJobPost
	if err := updateJobSeekService.UpdateJsStatusService(); err != nil {
		ResponseError(c, code.OperateFail, err)
		return
	}
	ResponseSuccess(c, code.Success, "更新成功")
}

func UpdateJsDetails(c *gin.Context) {
	modifiedJobPost := new(model.JobSeek)
	if err := c.ShouldBindJSON(modifiedJobPost); err != nil {
		ResponseError(c, code.BindError, err)
		return
	}
	updateJobSeekService := new(service.UpdateJobSeekService)
	if err := updateJobSeekService.UpdateJsDetailsService(); err != nil {
		ResponseError(c, code.OperateFail, err)
		return
	}
	ResponseSuccess(c, code.Success, "更新成功")
}
