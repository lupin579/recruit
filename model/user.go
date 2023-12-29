package model

type User struct {
	Uid            uint   `db:"uid" json:"uid"`
	Uname          string `db:"uname" json:"uname"`
	Identification string `db:"identification" json:"identification"`
	PhoneNumber    string `db:"phone_number" json:"phoneNumber"`
	QQNumber       string `db:"qq_number" json:"qqNumber "`
	WxNumber       string `db:"wx_number" json:"wxNumber"`
	SucceedTimes   uint16 `db:"succeed_times" json:"succeedTimes"`
}
