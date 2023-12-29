package mysql

func GetPasswordByPhoneNumber(phoneNumber string) (password string, err error) {
	str := "select password from user where phone_number = ?"
	err = db.Get(&password, str, phoneNumber)
	return
}

func GetUidByPhoneNumber(phoneNumber string) (uid string, err error) {
	str := "select uid from user where phone_number = ?"
	err = db.Get(&uid, str, phoneNumber)
	return
}

func GetPasswordByQQNumber(QQNumber string) (password string, err error) {
	str := "select password from user where qq_number = ?"
	err = db.Get(&password, str, QQNumber)
	return
}

func GetUidByQQNumber(QQNumber string) (uid string, err error) {
	str := "select uid from user where qq_number = ?"
	err = db.Get(&uid, str, QQNumber)
	return
}

func GetPasswordByWxNumber(WxNumber string) (password string, err error) {
	str := "select password from user where wx_number = ?"
	err = db.Get(&password, str, WxNumber)
	return
}

func GetUidByWxNumber(wxNumber string) (uid string, err error) {
	str := "select uid from user where wx_number = ?"
	err = db.Get(&uid, str, wxNumber)
	return
}
