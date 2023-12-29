package service

import "recruit/dao/mysql"

type Personality struct {
	Uid           int    `json:"uid"`
	Uname         string `json:"uname"`
	Identificaton string `json:"identificaton"`
	PhoneNumber   string `json:"phoneNumber"`
	QQNumber      string `json:"QQNumber"`
	WXNumber      string `json:"WXNumber"`
	SucceedTimes  int    `json:"succeedTimes"`
}

func (person *Personality) Me() (err error) {
	user, err := mysql.Me(person.Uid)
	if err != nil {
		return
	}
	person.Uname = user.Uname
	person.Identificaton = user.Identification
	person.PhoneNumber = user.PhoneNumber
	person.QQNumber = user.QQNumber
	person.WXNumber = user.WxNumber
	person.SucceedTimes = int(user.SucceedTimes)
	return
}
