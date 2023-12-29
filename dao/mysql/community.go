package mysql

import (
	"log"
	"recruit/model"
	"recruit/serializer"
)

func PostCommodities(commodity *model.Commodity) error {
	str := "insert into commodity(price,cmdt_name,on_sale,owner_id,description) values(?,?,?,?)"
	if _, err := db.Exec(str, commodity.Price, commodity.CmdtName, commodity.OnSale, commodity.OwnerID, commodity.Description); err != nil {
		log.Printf("PostCommodities:%s", err)
		return err
	}
	return nil
}

func GetMyCommodities(uid string) (cmdtList []*model.Commodity, err error) {
	str := "select * from commodity where owner_id=?"
	if err = db.Select(&cmdtList, str, uid); err != nil {
		log.Printf("GetMyCommodities:%s\n", err.Error())
		return
	}
	return
}

func GetCommodity(cid int) (cmdt *model.Commodity, err error) {
	str := "select * from commodity where cmdt_id=?"
	if err = db.Get(&cmdt, str, cid); err != nil {
		log.Printf("GetCommodity:%s\n", err.Error())
		return
	}
	return
}

/*
cmdt_id初次生成后就不允许更改，
用户无法决定cmdt_id的一切
*/
func UpdateCommodity(commodity model.Commodity) (err error) {
	str := "update commodity set cmdt_name=?,price=?,quantity=?,on_sale=?,owner_id=?,description=? where cmdt_id=?"
	_, err = db.Exec(str, commodity.CmdtName, commodity.Price, commodity.Quantity, commodity.OnSale, commodity.OwnerID, commodity.Description, commodity.CmdtID)
	if err != nil {
		log.Printf("mysql.UpdateCommodity:%s\n", err.Error())
	}
	return
}

func GetCommodities() (cmdt []*serializer.MainPageSerial, err error) {
	str := "select ta.cmdt_name,ta.price,ta.quantity,ta.on_sale,tb.uid,tb.uname from (SELECT * FROM commodity WHERE cmdt_id >= ( SELECT floor( RAND() * ( SELECT MAX( cmdt_id ) FROM commodity ) ) ) ORDER BY cmdt_id LIMIT 0,8) ta inner join user tb on ta.owner_id=tb.uid"
	err = db.Select(&cmdt, str)
	if err != nil {
		log.Printf("mysql.GetCommodities:%s\n", err.Error())
	}
	return
}

func DeleteCommodity(cid int) (err error) {
	str := "delete from commodity where cmdt_id=?"
	_, err = db.Exec(str, cid)
	if err != nil {
		log.Printf("mysql.DeleteCommodity:%s\n", err.Error())
	}
	return
}

func GetOwnerID(cid int) (oid *int, err error) {
	str := "select owner_id from commodity where cmdt_id=?"
	err = db.Get(oid, str, cid)
	if err != nil {
		log.Printf("mysql.GetOwnerID:%s\n", err.Error())
	}
	return
}

func GetCmdtName(cid int) (cname *string, err error) {
	str := "select cmdt_name from commodity where cmdt_id=?"
	err = db.Get(cname, str, cid)
	if err != nil {
		log.Printf("mysql.GetCmdtName:%s\n", err.Error())
	}
	return
}
