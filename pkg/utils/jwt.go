package utils

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var myKey = []byte("tokisaki")

/*
token签发
*/
func JWTIssue(userName string) (signedString string, err error) {
	ms := MyClaims{
		Username: userName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + 10000,
			Issuer:    "233",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &ms)
	signedString, err = token.SignedString(myKey)
	return
}

/*
token验证
*/
func JWTValidator(token string) (details string, err error) {
	claims, err := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err != nil {
		return "解析错误", err
	}
	if claims.Valid {
		return "token有效", nil
	}
	return "token已过期", errors.New("token is overdue(out of date)")
}
