package service

import (
	"recruit/dao/mysql"
	"recruit/model"
)

type JobSeekDetailsService struct {
	JsId uint16 `db:"jsid"`
}

func (jobSeekDetailsService *JobSeekDetailsService) JobSeekDetails() (jobSeek *model.JobSeek, err error) {
	jobSeek, err = mysql.JobSeekDetails(jobSeekDetailsService.JsId)
	return
}
