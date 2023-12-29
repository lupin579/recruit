package service

import (
	"errors"
	"recruit/dao/mysql"
	"recruit/model"
	"time"
)

type PeopleSeekPage struct {
	PsId    uint16 `db:"psid" json:"psId"`
	Status  string `db:"status" json:"status"`
	JobType string `db:"job_type" json:"jobType"`
	PubId   uint16 `db:"pub_id" json:"pubId"`
}

type PeopleSeekDetailsService struct {
	PsId uint16 `db:"psid" json:"psId"`
}

type PostPeopleSeekService struct {
}

type UpdatePeopleSeekService struct {
	Post *model.PeopleSeek
}

func (peopleSeekPage *PeopleSeekPage) PeopleSeekPosts() (psList []*PeopleSeekPage, err error) {
	postList, err := mysql.GetPeopleSeekPosts(peopleSeekPage.PubId)
	if err != nil {
		return
	}
	j := 0
	for _, post := range postList {
		if post.PubId != peopleSeekPage.PubId {
			postList[j] = post
			j++
		}
	}
	realPList := postList[:j]
	for _, realP := range realPList {
		psl := PeopleSeekPage{
			PsId:    realP.PsId,
			Status:  realP.Status,
			JobType: realP.JobType,
			PubId:   realP.PubId,
		}
		psList = append(psList, &psl)
	}
	return
}

func (peopleSeekDetailsService *PeopleSeekDetailsService) PeopleSeekDetails() (ps *model.PeopleSeek, err error) {
	ps, err = mysql.PeopleSeekDetails(peopleSeekDetailsService.PsId)
	return
}

func (postPeopleSeekService *PostPeopleSeekService) PostPeopleSeekService(ps *model.PeopleSeek, pubId uint16) (err error) {
	ps.PubDate = time.Now().Format("2006/01/02 15:04")
	ps.PubId = pubId
	ps.Status = "求职中"
	err = mysql.PostPeopleSeek(ps)
	return
}

/*
求职中，协商中，已完成
已完成的条目不可再更改状态，因为要统计成交次数，如果可随意更改，则成交次数无法统计
*/
func (updatePeopleSeekService *UpdatePeopleSeekService) UpdateJsStatusService() (err error) {
	if updatePeopleSeekService.Post.Status != "已完成" &&
		updatePeopleSeekService.Post.Status != "协商中" { //此处status为将要修改为的状态
		return errors.New("更新后的状态只能为协商中或已完成")
	}
	status, err := mysql.PeopleSeekStatus(updatePeopleSeekService.Post.PsId)
	if err != nil {
		return err
	}
	if status == "已完成" { //该status为条目原有状态
		return errors.New("已完成的条目不可修改状态")
	}
	err = mysql.UpdatePsStatus(updatePeopleSeekService.Post)
	return
}

/*
已完成的条目不可修改
*/
func (updatePeopleSeekService *UpdatePeopleSeekService) UpdatePsDetailsService() (err error) {
	status, err := mysql.PeopleSeekStatus(updatePeopleSeekService.Post.PsId)
	if err != nil {
		return err
	}
	if status == "已完成" { //该status为条目原有状态
		return errors.New("已完成的条目不可修改内容")
	}
	err = mysql.UpdatePsDetails(updatePeopleSeekService.Post)
	return
}
