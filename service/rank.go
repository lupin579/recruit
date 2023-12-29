package service

import (
	"recruit/dao/mysql"
	"recruit/serializer"
)

type RankService struct {
}

func (rankService *RankService) Rank() (rankSerialList []*serializer.RankSerializer, err error) {
	rankSerialList, err = mysql.Rank()
	return
}
