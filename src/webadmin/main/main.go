package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.etcd.io/etcd/clientv3"
	"time"
	"webadmin/model"
	_ "webadmin/router"
)

func initDb() (err error) {
	database, err := sqlx.Open("mysql", "root:@tcp(127.0.0.1:3306)/logadmin")
	if err != nil {
		logs.Warn("open mysql failed,", err)
		return
	}

	model.InitDb(database)
	return
}
func initEtcd() (err error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect failed, err:", err)
		return
	}

	model.InitEtcd(cli)
	return
}
func main() {
	//初始化数据库
	err := initDb()
	if err != nil {
		logs.Warn("initDb failed, err:%v", err)
		return
	}
	//初始化ETCD
	err = initEtcd()
	if err != nil {
		logs.Warn("init etcd failed, err:%v", err)
		return
	}
	beego.Run()
}
