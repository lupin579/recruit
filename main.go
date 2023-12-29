package main

import (
	"fmt"
	"recruit/dao/mysql"
	"recruit/routers"
	"recruit/settings"
)

func main() {
	if err := settings.Init(); err != nil {
		fmt.Printf("load config failed, err: %v\n", err)
		return
	}
	if err := mysql.Init(settings.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer mysql.Close() // 程序退出关闭数据库连接
	r := routers.SetupRouter(*settings.Conf)
	err := r.Run(fmt.Sprintf("0.0.0.0:%d", settings.Conf.Port))
	if err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		return
	}

}
