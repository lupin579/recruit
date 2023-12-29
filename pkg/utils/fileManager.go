package utils

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"mime/multipart"
	"recruit/settings"
	"strings"
)

func SaveImage(c *gin.Context, file *multipart.FileHeader, postTime int64, name string, uid int) (fileName *string, err error) {
	//判断文件格式是否为jpg或png
	tp := strings.Split(file.Filename, ".")[1]
	if tp != "jpg" || tp != "png" {
		err = errors.New("invalid type")
		log.Printf("saveAvator:%s", err.Error())
		return nil, err
	}

	//重命名用户上传物品图片，便于管理
	file.Filename = fmt.Sprintf("%s_%s_%s.%s", uid, name, postTime, tp) //文件重命名：用户ID_物品名_上传时间.文件格式

	//规定文件大小不可大于3MB
	if file.Size > 3*1024*1024 {
		err = errors.New("overwhelming file! file size cannot exceed 3MB")
		log.Printf("saveAvator:%s", err.Error())
		return nil, err
	}

	//保存至指定路径(static/commodities_${uid}/)
	absPath := fmt.Sprintf("%s/%s_%s", settings.StaticPath, "commodities", uid)
	err = c.SaveUploadedFile(file, absPath)
	if err != nil {
		log.Printf("saveAvator.SaveUploadedFile:%s", err.Error())
		return nil, err
	}
	wholeName := fmt.Sprintf("%s/%s", absPath, file.Filename)
	return &(wholeName), nil
}
