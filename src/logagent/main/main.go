package main

import (
	"fmt"

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

}
