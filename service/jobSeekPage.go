package service

import (
	"math/rand"
	"recruit/dao/mysql"
	"time"
)

type JobSeekPage struct {
	JsId    uint16 `db:"jsid" json:"jsId"`
	Status  string `db:"status" json:"status"`
	JobType string `db:"job_type" json:"jobType"`
	PubId   uint16 `db:"pub_id" json:"pubId"`
}

func (jobSeekPage *JobSeekPage) JobSeekPosts() (returnedJobSeekPosts []JobSeekPage, err error) {
	jobSeekPosts, err := mysql.JobSeekPosts()
	if err != nil {
		return nil, err
	}
	m := 0
	var tmpReturnedJobSeekPosts []JobSeekPage
	for _, jobSeekPost := range jobSeekPosts {
		if jobSeekPage.PubId != jobSeekPost.PubId { //不会将pubid为自己pubid的帖子返回给自己,此处jobSeekPage中的PubId指的是登陆者的uid以防止前面说的情况出现
			tmpReturnedJobSeekPosts[m].JobType = jobSeekPost.JobType
			tmpReturnedJobSeekPosts[m].JsId = jobSeekPost.JsId
			tmpReturnedJobSeekPosts[m].PubId = jobSeekPost.PubId
			tmpReturnedJobSeekPosts[m].Status = jobSeekPost.Status
			m++
		}
	}
	r := rand.New(rand.NewSource(time.Now().Unix())) //实现随机顺序返回
	randnum := r.Perm(len(returnedJobSeekPosts))
	for i := 0; i < len(randnum); i++ {
		returnedJobSeekPosts[i] = tmpReturnedJobSeekPosts[randnum[i]]
	}
	return
}
