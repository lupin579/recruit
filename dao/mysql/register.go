package mysql

import (
	"database/sql"
	"log"
)

func UserDataSave(id int, uname, identification, phoneNumber, QQNumber, wxNumber string) (err error, sqlResult sql.Result) {
	str := "insert into user values (?,?,?,?,?,?)"
	if sqlResult, err = db.Exec(str, id, uname, identification, phoneNumber, QQNumber, wxNumber); err != nil {
		return
	}
	return
}

func UserCount() (num *int, err error) {
	str := "select count(*) from user"
	if err = db.Get(num, str); err != nil {
		log.Printf("mysql.UserCount:%s\n", err)
	}
	return num, err
}
