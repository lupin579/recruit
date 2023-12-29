package model

type PeopleSeek struct {
	PsId    uint16 `db:"psid"`
	Status  string `db:"status" json:"status"`
	JobType string `db:"job_type" json:"jobType"`
	Details string `db:"details" json:"details"`
	PubId   uint16 `db:"pub_id" json:"pubId"`
	PubDate string `db:"pub_date" json:"publishedDate"`
}
