package main

import (
	"fmt"
	"logagent/kafka"

	"github.com/astaxie/beego/logs"

	"logagent/tailf"
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
	logs.Debug("initialize success")
	logs.Debug("load conf succcess,config%v\n", appConfig)
	//3、初始化tail组件
	err = tailf.InitTail(appConfig.collectConf, appConfig.chanSize)
	if err != nil {
		logs.Error("init tail failed,err:%v\n", err)
		return
	}
	logs.Debug("initialize all success")
	err = kafka.InitKafka(appConfig.kafkaAddr)
	if err != nil {
		logs.Error("init kafka failed,err:%v\n", err)
		return
	}
	err = serverRun()
	if err != nil {
		logs.Error("serverRun failed,err:%v", err)
		return
	}
	logs.Info("Program Exit!")

}
