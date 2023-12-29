package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"recruit/model"
	"recruit/pkg/code"
	"recruit/pkg/utils"
	"recruit/service"
	"strconv"
	"time"
)

func removeImage(fileName *string) {
	err := os.Remove(*fileName)
	if err != nil {
		log.Println("删除失败")
	}
}

func PostCommodities(c *gin.Context) {
	//绑定表单数据
	form, err := c.MultipartForm()
	if err != nil {
		ResponseError(c, code.BindError, err)
		return
	}

	//数据写入至结构体Comomdity
	commodity := new(model.Commodity)
	commodity.OwnerID, err = strconv.Atoi(form.Value["ownerID"][0])
	if err != nil {
		ResponseError(c, code.BindError, err)
		return
	}
	commodity.Price, err = strconv.Atoi(form.Value["price"][0])
	if err != nil {
		ResponseError(c, code.BindError, err)
		return
	}
	commodity.CmdtName = form.Value["cmdtName"][0]
	commodity.Description = &form.Value["description"][0]
	commodity.OnSale = 0

	//文件写入
	postTime := time.Now().Unix()
	fileHeader := form.File["image"][0]
	fileName, err := utils.SaveImage(c, fileHeader, postTime, commodity.CmdtName, commodity.OwnerID)
	if err != nil {
		ResponseError(c, code.ServerBusy, err)
		return
	}

	//将商品信息存入数据库
	commodityService := new(service.CommodityService)
	err = commodityService.PostCommodities()
	if err != nil {
		removeImage(fileName)
		ResponseError(c, code.OperateFail, err)
		return
	}

	ResponseSuccess(c, code.Success, nil)
}

func GetMyCommodities(c *gin.Context) {
	uid := c.Query("uid")
	commodity := new(service.CommodityService)
	list, err := commodity.GetMyCommodities(uid)
	if err != nil {
		ResponseError(c, code.GetError, err)
		return
	}
	ResponseSuccess(c, code.Success, list)
}

func GetCommodityDetails(c *gin.Context) {
	cid := c.Query("cid")
	commodity := new(service.CommodityService)
	var err error
	commodity.OwnerID, err = strconv.Atoi(cid)
	if err != nil {
		ResponseError(c, code.ServerBusy, err)
		return
	}
	cmdtS, err := commodity.GetCommodityDetails()
	if err != nil {
		ResponseError(c, code.GetError, err)
		return
	}
	ResponseSuccess(c, code.Success, cmdtS)
}

func UpdateCommodity(c *gin.Context) {
	commodityService := new(service.CommodityService)
	err := c.ShouldBindJSON(commodityService)
	if err != nil {
		ResponseError(c, code.BindError, err)
		return
	}
	err = commodityService.UpdateCommodity()
	if err != nil {
		ResponseError(c, code.OperateFail, err)
		return
	}
	ResponseSuccess(c, code.Success, nil)
}

func UpdateCmdtImage(c *gin.Context) {
	cid := c.Query("cmdtID")
	fileHeader, err := c.FormFile("image")
	if err != nil {
		ResponseError(c, code.BindError, err)
		return
	}

	cmdt := new(service.CommodityService)
	cmdt.CmdtID, err = strconv.Atoi(cid)
	if err != nil {
		ResponseError(c, code.ServerBusy, err)
		return
	}
	err = cmdt.UpdateCmdtImage(c, fileHeader)
	if err != nil {
		ResponseError(c, code.OperateFail, err)
		return
	}
	ResponseSuccess(c, code.Success, nil)
}

/*
developing(under development)开发中
*/
func GetCommodities(c *gin.Context) {
	cmdt := new(service.CommodityService)
	resList, err := cmdt.GetCommodities()
	if err != nil {
		ResponseError(c, code.GetError, err)
		return
	}
	ResponseSuccess(c, code.Success, resList)
}

func DeleteCommodity(c *gin.Context) {
	cid := c.Query("cid")
	cmdt := new(service.CommodityService)
	var err error
	cmdt.CmdtID, err = strconv.Atoi(cid)
	if err != nil {
		ResponseError(c, code.ServerBusy, err)
		return
	}
	err = cmdt.DeleteCommodity()
	if err != nil {
		ResponseError(c, code.OperateFail, err)
		return
	}
	ResponseSuccess(c, code.Success, nil)
}
