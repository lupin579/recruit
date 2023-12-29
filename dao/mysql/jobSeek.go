package mysql

import "recruit/model"

/*
status为已完成的不会被从数据库中取出
*/
func JobSeekPosts() (jobSeekPosts []*model.JobSeek, err error) {
	str := "select * from job_seek where status=? or status=?"
	err = db.Select(&jobSeekPosts, str, "求职中", "协商中")
	return
}

func JobSeekDetails(jsId uint16) (jobSeekDetails *model.JobSeek, err error) {
	str := "select * from job_seek where js_id=?"
	err = db.Select(jobSeekDetails, str, jsId)
	return
}

func JobSeekStatus(jsid uint16) (status string, err error) {
	str := "select status from job_seek where jsid=?"
	err = db.Get(&status, str, jsid)
	return
}

func UpdateJsStatus(post *model.JobSeek) (err error) {
	str := "update job_seek set status=? where jsid=?"
	db.Exec(str, post.Status, post.JsId)
	return
}

func UpdateJsDetails(post *model.JobSeek) (err error) {
	str := "update job_seek set job_type=?,details=? where jsid=?"
	db.Exec(str, post.Details, post.JobType, post.JsId)
	return
}

func PostJobSeek(post *model.JobSeek) (err error) {
	str := "insert job_seek(status,job_type,details,pub_id,pub_date) values(?,?,?,?,?)"
	_, err = db.Exec(str, post.Status, post.JobType, post.Details, post.PubId, post.PubDate)
	return
}
