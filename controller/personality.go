package controller

import (
	"errors"
	"fmt"
	"log"
	"mime/multipart"
	"recruit/pkg/code"
	"recruit/service"
	"recruit/settings"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func saveAvator(c *gin.Context, file *multipart.FileHeader, uid string) (err error) {
	//判断文件格式是否为jpg或png
	tp := strings.Split(file.Filename, ".")[1]
	if tp != "jpg" || tp != "png" {
		err = errors.New("invalid type")
		log.Printf("saveAvator:%s", err.Error())
		return err
	}

	//重命名用户上传头像，便于管理
	createDate := time.Now().Unix()
	file.Filename = fmt.Sprintf("%s_%s.%s", uid, createDate, tp) //文件重命名：用户ID_修改头像日期.文件格式

	//规定文件大小不可大于3MB
	if file.Size > 3*1024*1024 {
		err = errors.New("overwhelming file! file size cannot exceed 3MB")
		log.Printf("saveAvator:%s", err.Error())
		return err
	}

	//保存至指定路径
	err = c.SaveUploadedFile(file, settings.StaticPath)
	if err != nil {
		log.Printf("saveAvator.SaveUploadedFile:%s", err.Error())
		return err
	}
	return
}

func PostAvator(c *gin.Context) {
	//获取用户uid和图片fileHeader
	uid := c.Param("uid")
	file, err := c.FormFile("avator")
	if err != nil {
		log.Printf("Personality.FormFile:%s", err.Error())
		ResponseError(c, code.BindError, err)
	}

	//保存图片
	err = saveAvator(c, file, uid)
	if err != nil {
		log.Printf("Personality.saveAvator:%s", err.Error())
		ResponseError(c, code.ServerBusy, err)
	}

	ResponseSuccess(c, code.Success, nil)
}

func Me(c *gin.Context) {
	uid := c.Query("uid")
	personality := new(service.Personality)
	var err error
	personality.Uid, err = strconv.Atoi(uid)
	if err != nil {
		ResponseError(c, code.ServerBusy, err)
	}
	if err = personality.Me(); err != nil {
		ResponseError(c, code.GetError, err)
	}
	ResponseSuccess(c, code.Success, personality)
}
