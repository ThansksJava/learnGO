package main

import (
	"fmt"

	"github.com/astaxie/beego/config"
)

//Config 存储配置结构体
type Config struct {
	logLevel string
	logPath  string

	chanSize  int
	kafkaAddr string
	// collectConf []tailf.CollectConf

	etcdAddr string
	etcdKey  string
}

func loadConf(confType string, fileName string) (err error) {
	//调用beego读取配置API
	conf, err := config.NewConfig(confType, fileName)
	if err != nil {
		fmt.Println("new config faild,err:", err)
	}

	//init configuration struct
	appConfig := &Config{}

	//load log_level from configuration
	appConfig.logLevel = conf.String("logs::log_level")
	if len(appConfig.logLevel) == 0 {
		appConfig.logLevel = "debug"
	}

	//load log_path from configuration
	appConfig.logPath = conf.String("logs::log_path")
	if len(appConfig.logPath) == 0 {
		appConfig.logPath = "./logs"
	}

	//load chan_size from configuration
	appConfig.chanSize, err = conf.Int("collect::chan_size")
	if err != nil {
		appConfig.chanSize = 100
	}
	
	return nil
}
