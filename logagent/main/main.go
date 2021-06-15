package main

import (
	"fmt"
	"learnGO/logagent/kafka"

	"github.com/astaxie/beego/logs"

	"learnGO/logagent/tailf"
)

func main() {
	//1、加载配置文件
	filename := "D:/workspace/learnGO/src/logagent/commonconf/logagent.conf"
	err := loadConf("ini", filename)
	if err != nil {
		fmt.Printf("load conf failed,err:%v\n", err)
		panic("load conf failed")
	}
	//2、初始化日志组件
	err = initLogger()
	if err != nil {
		fmt.Printf("load logger failed,err:%v\n", err)
		panic("load logger failed")
	}
	logs.Debug("load conf succcess,config%v\n", appConfig)
	logs.Debug("initialize Logger success")

	collectConf, err := initEtcd(appConfig.etcdAddr, appConfig.etcdKey)
	if err != nil {
		logs.Error("init etcd failed, err:%v", err)
		return
	}
	logs.Debug("initialize etcd succ")
	//3、初始化tail组件
	err = tailf.InitTail(collectConf, appConfig.chanSize)
	if err != nil {
		logs.Error("init tail failed,err:%v\n", err)
		return
	}
	logs.Debug("initialize Tail success")
	err = kafka.InitKafka(appConfig.kafkaAddr)
	if err != nil {
		logs.Error("init kafka failed,err:%v\n", err)
		return
	}
	logs.Debug("initialize Kafka success")
	logs.Debug("initialize all success")
	err = serverRun()
	if err != nil {
		logs.Error("serverRun failed,err:%v", err)
		return
	}
	logs.Info("Program Exit!")
}
