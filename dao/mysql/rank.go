package mysql

import (
	"log"
	"recruit/serializer"
)

func Rank() (rankSerialList []*serializer.RankSerializer, err error) {
	str := "select uid,uname,identification,succeed_times from user order by succeed_times desc limit 10"
	if err = db.Select(&rankSerialList, str); err != nil {
		log.Printf("mysql.Rank:%s\n", err.Error())
	}
	return
}
