package mysql

import (
	"log"
	"recruit/model"
)

func Me(id int) (user *model.User, err error) {
	str := "select * from user where uid=?"
	err = db.Get(user, str, id)
	if err != nil {
		log.Printf("mysql.Me:%s\n")
	}
	return
}
