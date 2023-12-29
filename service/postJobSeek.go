package service

import (
	"recruit/dao/mysql"
	"recruit/model"
	"time"
)

type PostJobSeekService struct {
}

func (PostJobSeekService *PostJobSeekService) PostJobSeekService(postedJobSeek *model.JobSeek, uid uint16) (err error) {
	postedJobSeek.PubDate = time.Now().Format("2006/01/02 15:04")
	postedJobSeek.PubId = uid
	postedJobSeek.Status = "求职中"
	err = mysql.PostJobSeek(postedJobSeek)
	return
}
