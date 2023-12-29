package service

import (
	"errors"
	"recruit/dao/mysql"
	"recruit/model"
)

type UpdateJobSeekService struct {
	Post *model.JobSeek
}

/*
求职中，协商中，已完成
已完成的条目不可再更改状态，因为要统计成交次数，如果可随意更改，则成交次数无法统计
*/
func (updateJobSeekService *UpdateJobSeekService) UpdateJsStatusService() (err error) {
	if updateJobSeekService.Post.Status != "已完成" &&
		updateJobSeekService.Post.Status != "协商中" { //此处status为将要修改为的状态
		return errors.New("更新后的状态只能为协商中或已完成")
	}
	status, err := mysql.JobSeekStatus(updateJobSeekService.Post.JsId)
	if err != nil {
		return err
	}
	if status == "已完成" { //该status为条目原有状态
		return errors.New("已完成的条目不可修改状态")
	}
	err = mysql.UpdateJsStatus(updateJobSeekService.Post)
	return
}

/*
已完成的条目不可修改
*/
func (updateJobSeekService *UpdateJobSeekService) UpdateJsDetailsService() (err error) {
	status, err := mysql.JobSeekStatus(updateJobSeekService.Post.JsId)
	if err != nil {
		return err
	}
	if status == "已完成" { //该status为条目原有状态
		return errors.New("已完成的条目不可修改内容")
	}
	err = mysql.UpdateJsDetails(updateJobSeekService.Post)
	return
}
