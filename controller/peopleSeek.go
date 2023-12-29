package controller

import (
	"github.com/gin-gonic/gin"
	"recruit/model"
	"recruit/pkg/code"
	"recruit/service"
	"strconv"
)

/*
招人主页面
*/
func PeopleSeekPage(c *gin.Context) {
	sid := c.Param("uid")
	uid, _ := strconv.ParseUint(sid, 10, 64)
	peopleSeekPage := service.PeopleSeekPage{
		PubId: uint16(uid),
	}
	posts, err := peopleSeekPage.PeopleSeekPosts()
	if err != nil {
		ResponseError(c, code.ServerBusy, err)
	}
	ResponseSuccess(c, code.Success, posts)
}

/*
招人帖详情
*/
func PeopleSeekDetails(c *gin.Context) {
	psId := c.Param("psId")
	peopleSeekDetailsService := new(service.PeopleSeekDetailsService)
	u16, _ := strconv.ParseUint(psId, 10, 16)
	peopleSeekDetailsService.PsId = uint16(u16)
	peoplePost, err := peopleSeekDetailsService.PeopleSeekDetails()
	if err != nil {
		ResponseError(c, code.GetError, err)
	}
	ResponseSuccess(c, code.Success, peoplePost)
}

/*
发布招人帖
*/
func PostPeopleSeek(c *gin.Context) {
	sid := c.Param("uid")
	uid, _ := strconv.ParseUint(sid, 10, 16)
	peopleSeek := new(model.PeopleSeek)
	if err := c.ShouldBind(peopleSeek); err != nil {
		ResponseError(c, code.BindError, err)
	}
	postedPeopleSeekService := new(service.PostPeopleSeekService)
	if err := postedPeopleSeekService.PostPeopleSeekService(peopleSeek, uint16(uid)); err != nil {
		ResponseError(c, code.OperateFail, err)
	}
	ResponseSuccess(c, code.Success, "发布成功！")
}

/*
更新招人帖状态（招人中or协商中or已完成）
*/
func UpdatePsStatus(c *gin.Context) {
	modifiedPeoplePost := new(model.PeopleSeek)
	if err := c.ShouldBindJSON(modifiedPeoplePost); err != nil {
		ResponseError(c, code.BindError, err)
		return
	}
	updatePeopleSeekService := new(service.UpdatePeopleSeekService)
	if err := updatePeopleSeekService.UpdateJsStatusService(); err != nil {
		ResponseError(c, code.OperateFail, err)
		return
	}
	ResponseSuccess(c, code.Success, "更新成功")
}

/*
更新招人帖详细信息（不包含帖子ID）
*/
func UpdatePsDetails(c *gin.Context) {
	modifiedPeoplePost := new(model.PeopleSeek)
	if err := c.ShouldBindJSON(modifiedPeoplePost); err != nil {
		ResponseError(c, code.BindError, err)
		return
	}
	updatePeopleSeekService := new(service.UpdatePeopleSeekService)
	if err := updatePeopleSeekService.UpdatePsDetailsService(); err != nil {
		ResponseError(c, code.OperateFail, err)
		return
	}
	ResponseSuccess(c, code.Success, "更新成功")
}
