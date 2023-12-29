package mysql

import "recruit/model"

func FullTextSearch(key string) (jobSeek []*model.JobSeek, err error) {
	str := "select * from job_seek where job_type=? or details=?"
	err = db.Select(&jobSeek, str, key, key)
	return
}
