package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"mime/multipart"
	"os"
	"recruit/dao/mysql"
	"recruit/model"
	"recruit/pkg/utils"
	"recruit/serializer"
	"recruit/settings"
	"strconv"
	"strings"
	"time"
)

type CommodityService struct {
	CmdtID      int     `json:"cmdtID"`
	CmdtName    string  `json:"cmdtName"`
	Price       int     `json:"price"`
	Quantity    int     `json:"quantity"`
	OnSale      int8    `json:"onSale"`
	OwnerID     int     `json:"ownerID"`
	Description *string `json:"description"`
}

func (commodityService *CommodityService) PostCommodities() (err error) {
	commodity := model.Commodity(*commodityService)
	err = mysql.PostCommodities(&commodity)
	if err != nil {
		return err
	}
	return nil
}

func getCommodities(uid int) (list []*string, err error) {
	absPath := fmt.Sprintf("%s/%s_%s", settings.StaticPath, "commodities", uid)
	fileInfo, err := ioutil.ReadDir(absPath)
	if err != nil {
		return nil, err
	}
	for _, file := range fileInfo {
		name := file.Name()
		list = append(list, &name)
	}
	return list, nil
}

func (commodityService *CommodityService) GetMyCommodities(uid string) (serialList []*serializer.CmdtSerial, err error) {
	cmdtList, err := mysql.GetMyCommodities(uid)
	if err != nil {
		log.Printf("GetMyCommodities:%s\n", err.Error())
		return nil, err
	}
	iuid, err := strconv.Atoi(uid)
	if err != nil {
		log.Printf("GetMyCommodities.Atoi:%s\n", err.Error())
		return nil, err
	}
	paths, err := getCommodities(iuid)
	if err != nil {
		log.Printf("GetMyCommodities.getCommodities:%s\n", err.Error())
		return nil, err
	}
	for _, cmdt := range cmdtList {
		var realPath *string

		j := 0
		for _, path := range paths {
			if !strings.Contains(*path, cmdt.CmdtName) {
				paths[j] = path
				j++
			} else {
				realPath = path
			}
		}
		paths = paths[:j]

		serial := serializer.CmdtSerial{
			ImagePath: realPath,
			Commodity: *cmdt,
		}
		serialList = append(serialList, &serial)
	}
	return serialList, nil
}

func (commodityService *CommodityService) GetCommodityDetails() (cmdtS *serializer.CmdtSerial, err error) {
	cmdt, err := mysql.GetCommodity(commodityService.OwnerID)
	if err != nil {
		log.Printf("GetCommodityDetails.GetCommodity:%s\n", err.Error())
		return
	}
	cmdtS.Commodity = *cmdt

	//
	usersCmdtList, err := getCommodities(commodityService.CmdtID)
	if err != nil {
		log.Printf("GetCommodityDetails.getCommodities:%s\n", err.Error())
		return
	}
	var realName *string
	for _, ucmdt := range usersCmdtList {
		if strings.Contains(*ucmdt, cmdt.CmdtName) {
			realName = ucmdt
			break
		}
	}
	//找到该文件(未找到文件名为空)
	if realName != nil {
		tPath := fmt.Sprintf("/static/commodities_%d/%s", commodityService.OwnerID, *realName)
		cmdtS.ImagePath = &tPath
	}

	return
}

func (commodity *CommodityService) UpdateCommodity() (err error) {
	err = mysql.UpdateCommodity(model.Commodity(*commodity))
	return
}

func (commodity *CommodityService) UpdateCmdtImage(c *gin.Context, header *multipart.FileHeader) (err error) {
	cmdt, err := mysql.GetCommodity(commodity.CmdtID)
	if err != nil {
		return
	}

	dirPath := fmt.Sprintf("%s/commodity_%d", settings.StaticPath, cmdt.OwnerID)
	fileHeaders, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return
	}
	for _, fileHeader := range fileHeaders {
		if strings.Split(fileHeader.Name(), "_")[1] == cmdt.CmdtName {
			os.Remove(fmt.Sprintf("%s/%s", dirPath, fileHeader.Name()))
		}
	}

	_, err = utils.SaveImage(c, header, time.Now().Unix(), cmdt.CmdtName, cmdt.OwnerID)
	if err != nil {
		log.Printf("updateCmdtImage.SaveImage:%s\n", err.Error())
	}
	return
}

func (cmdt *CommodityService) GetCommodities() (cmdtList []*serializer.MainPageSerial, err error) {
	cmdtList, err = mysql.GetCommodities()
	if err != nil {
		return
	}
	for _, cmdt := range cmdtList {
		files, err := ioutil.ReadDir(fmt.Sprintf("%s/commodities_%s", settings.StaticPath, cmdt.Uid))
		if err != nil {
			log.Printf("service.GetCommodities:%s\n", err.Error())
			return nil, err
		}
		for _, file := range files {
			if strings.Split(file.Name(), "_")[1] == cmdt.CmdtName {
				imagePath := fmt.Sprintf("%s/commodities_%s/%s", settings.StaticPath, cmdt.Uid, file.Name())
				cmdt.ImagePath = &imagePath
			}
		}
	}
	return
}

func (cmdt *CommodityService) DeleteCommodity() (err error) {
	cname, err := mysql.GetCmdtName(cmdt.CmdtID)
	if err != nil {
		return
	}
	ownerID, err := mysql.GetOwnerID(cmdt.CmdtID)
	if err != nil {
		return
	}
	err = mysql.DeleteCommodity(cmdt.CmdtID)
	if err != nil {
		return
	}
	fileInfos, err := ioutil.ReadDir(fmt.Sprintf("%s/commodities_%s", settings.StaticPath, ownerID))
	if err != nil {
		return
	}
	for _, file := range fileInfos {
		if strings.Split(file.Name(), "_")[1] == *cname {
			if err := os.Remove(fmt.Sprintf("%s/commodities_%s/%s", settings.StaticPath, ownerID, file.Name())); err != nil {
				log.Printf("Service.DeleteCommodity.RemoveFileFailed:%s\n", err.Error())
			}
		}
	}
	return
}
