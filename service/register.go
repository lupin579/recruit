package service

import (
	"crypto/rand"
	"errors"
	"math/big"
	"recruit/cache"
	"recruit/dao/mysql"
	"recruit/pkg/utils"
	"strings"
)

type Register struct {
	Uname          string `db:"uname" json:"uname"`
	PassWord       string `db:"password" json:"password"`
	Identification string `db:"identification" json:"identification"`
	PhoneNumber    string `db:"phone_number" json:"phoneNumber"`
	QQNumber       string `db:"qq_number" json:"qqNumber "`
	WxNumber       string `db:"wx_number" json:"wxNumber"`
	VeriCode       uint16 `json:"veriCode"`
}

func (register *Register) Register() (uid *int, err error) {
	if err, _ = cache.ValidateVeriCode(register.PhoneNumber, register.VeriCode); err != nil {
		return nil, err
	}
	identification := strings.Trim(register.Identification, " ") //清除用户不小心输入的空格
	if identification != "招聘者" || identification != "被招聘者" {
		return nil, errors.New("职位不正确")
	}
	num, err := mysql.UserCount()
	if err != nil {
		return nil, err
	}
	*uid = *num + 1001
	err, _ = mysql.UserDataSave(*uid, register.Uname, register.Identification, register.PhoneNumber, register.QQNumber, register.WxNumber)
	return
}

func MsgSender(phoneNumber string) (details string, err error) {

	num, _ := rand.Int(rand.Reader, big.NewInt(9999))
	if err, details = cache.MsgTimeout(phoneNumber, uint16(num.Uint64())); err != nil {
		return details, err
	}
	if err = utils.SendMessage(phoneNumber, uint16(num.Uint64())); err != nil {
		details = "验证码发送失败，请重试"
		cache.DeleteMsgCacheIfExist(phoneNumber)
		return
	}
	details = "发送成功！"
	return
}
