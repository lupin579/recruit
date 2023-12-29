package serializer

import "recruit/model"

type CmdtSerial struct {
	ImagePath *string
	model.Commodity
}

type MainPageSerial struct {
	ImagePath *string `json:"imagePath"`
	CmdtName  string  `json:"cmdtName" db:"cmdt_name"`
	Uname     string  `json:"uname" db:"uname"`
	Uid       int     `json:"uid" db:"uid"`
	Price     int     `json:"price" db:"price"`
	Quantiry  int     `json:"quantiry" db:"quantity"`
	Onsale    int     `json:"onsale" db:"on_sale"`
}
