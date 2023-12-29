package cache

import (
	"errors"
	"strconv"
	"time"
)

func MsgTimeout(phoneNumber string, num uint16) (err error, details string) {
	if err := RedisCache.Get(phoneNumber).Err(); err != nil {
		RedisCache.Set(phoneNumber, num, 5*time.Minute)
		return nil, "上条信息已过期"
	} else {
		return errors.New("上条信息未过期"), "请待上条验证码过期后重试"
	}
}

func DeleteMsgCacheIfExist(phoneNumber string) {
	if err := RedisCache.Get(phoneNumber).Err(); err != nil {
		RedisCache.Del(phoneNumber)
	}
}

func ValidateVeriCode(phoneNumber string, veriCode uint16) (err error, details string) {
	if err = RedisCache.Get(phoneNumber).Err(); err != nil {
		return err, "验证码已过期"
	} else {
		snum := RedisCache.Get(phoneNumber).Val()
		unum, _ := strconv.ParseUint(snum, 10, 16)
		if unum == uint64(veriCode) {
			return nil, "验证码正确"
		} else {
			return errors.New("验证码错误"), "验证码错误"
		}
	}
}
