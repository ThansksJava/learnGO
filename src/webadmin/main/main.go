package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"webadmin/model"
)

func initDb() (err error) {
	database, err := sqlx.Open("mysql", "root:root@tcp(127.0.0.1:3306)/logadmin")
	if err != nil {
		logs.Warn("open mysql failed,", err)
		return
	}

	model.InitDb(database)
	return
}
func main() {
	//初始化数据库
	err := initDb()
	if err != nil {
		logs.Warn("initDb failed, err:%v", err)
		return
	}
	////初始化ETCD
	//err = initEtcd()
	//if err != nil {
	//	logs.Warn("init etcd failed, err:%v", err)
	//	return
	//}
	beego.Run()
}
