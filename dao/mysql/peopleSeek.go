package mysql

import "recruit/model"

func GetPeopleSeekPosts(pubId uint16) (postList []*model.PeopleSeek, err error) {
	str := "select * from people_seek where status=? or status=?"
	err = db.Select(&postList, str, "求职中", "协商中")
	return
}

func PeopleSeekDetails(psId uint16) (peopleSeekDetails *model.PeopleSeek, err error) {
	str := "select * from people_seek where ps_id=?"
	err = db.Select(peopleSeekDetails, str, psId)
	return
}

func PostPeopleSeek(post *model.PeopleSeek) (err error) {
	str := "insert people_seek(status,job_type,details,pub_id,pub_date) values(?,?,?,?,?)"
	_, err = db.Exec(str, post.Status, post.JobType, post.Details, post.PubId, post.PubDate)
	return
}

func PeopleSeekStatus(psid uint16) (status string, err error) {
	str := "select status from people_seek where psid=?"
	err = db.Get(&status, str, psid)
	return
}

func UpdatePsStatus(post *model.PeopleSeek) (err error) {
	str := "update people_seek set status=? where psid=?"
	db.Exec(str, post.Status, post.PsId)
	return
}

func UpdatePsDetails(post *model.PeopleSeek) (err error) {
	str := "update people_seek set details=? where psid=?"
	db.Exec(str, post.Details, post.PsId)
	return
}
