package service

import (
	"recruit/dao/mysql"
)

type Search struct {
	Uid uint16 `json:"uid"`
	Key string `json:"key"`
}

type SearchSerializer struct {
	JsId    uint16 `db:"jsid"`
	Status  string `db:"status" json:"status"`
	JobType string `db:"job_type" json:"jobType"`
	PubId   uint16 `db:"pub_id" json:"pubId"`
}

func (search *Search) SearchService() (jobSeek []SearchSerializer, err error) {
	tmpJobSeek, err := mysql.FullTextSearch(search.Key)
	m := 0
	for _, jobSeekPost := range tmpJobSeek {
		if search.Uid != jobSeekPost.PubId {
			jobSeek[m].JobType = jobSeekPost.JobType
			jobSeek[m].JsId = jobSeekPost.JsId
			jobSeek[m].PubId = jobSeekPost.PubId
			jobSeek[m].Status = jobSeekPost.Status
			m++
		}
	}
	return
}
