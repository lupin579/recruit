package service

import (
	"errors"
	"log"
	"recruit/dao/mysql"
)

type Login struct {
	AccountNumber string `json:"accountNumber"`
	Password      string `json:"password"`
	LoginMethod   string `json:"loginMethod"`
}

var LoginFunc = map[string]func(string, string) (string, error){
	"PhoneNumber": phoneNumberLogin,
	"QQNumber":    qqNumberLogin,
	"WxNumber":    wxNumberLogin,
}

func (login *Login) MulMethodsLogin() (details string, err error) {
	details, err = LoginFunc[login.LoginMethod](login.AccountNumber, login.Password)
	return
}

func phoneNumberLogin(phoneNumber string, password string) (details string, err error) {
	realPassWord, err := mysql.GetPasswordByPhoneNumber(phoneNumber)
	if err != nil {
		return "获取密码失败", err
	}
	if realPassWord != password {
		return "密码错误", errors.New("密码错误")
	}
	uid, err := mysql.GetUidByPhoneNumber(phoneNumber)
	if err != nil {
		log.Print("phone登录uid获取\n")
		return "serverBusy", err
	}
	return uid, nil
}

func qqNumberLogin(QQNumber string, password string) (details string, err error) {
	realPassWord, err := mysql.GetPasswordByQQNumber(QQNumber)
	if err != nil {
		return "获取密码失败", err
	}
	if realPassWord != password {
		return "密码错误", errors.New("密码错误")
	}
	uid, err := mysql.GetUidByQQNumber(QQNumber)
	if err != nil {
		log.Print("qq登录uid获取\n")
		return "serverBusy", err
	}
	return uid, nil
}

func wxNumberLogin(WxNumber string, password string) (details string, err error) {
	realPassWord, err := mysql.GetPasswordByWxNumber(WxNumber)
	if err != nil {
		return "获取密码失败", err
	}
	if realPassWord != password {
		return "密码错误", errors.New("密码错误")
	}
	uid, err := mysql.GetUidByPhoneNumber(WxNumber)
	if err != nil {
		log.Print("wx登录uid获取\n")
		return "serverBusy", err
	}
	return uid, nil
}
