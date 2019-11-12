package main

import (
	"errors"
	"fmt"

	"github.com/astaxie/beego/config"

	cm "logagent/commonconf"
)

var (
	appConfig *Config
)

//Config 存储配置结构体
type Config struct {
	logLevel    string
	logPath     string
	collectConf []cm.CollectConf
	chanSize    int
}

func loadCollectConf(conf config.Configer) (err error) {
	var cc cm.CollectConf
	cc.LogPath = conf.String("collect::log_path")
	if len(cc.LogPath) == 0 {
		err = errors.New("invalid collect::log_path")
		return
	}

	cc.Topic = conf.String("collect::topic")
	if len(cc.LogPath) == 0 {
		err = errors.New("invalid collect::topic")
		return
	}

	//追加配置
	appConfig.collectConf = append(appConfig.collectConf, cc)

	return
}
func loadConf(confType string, fileName string) (err error) {
	//调用beego读取配置API
	conf, err := config.NewConfig(confType, fileName)
	if err != nil {
		fmt.Println("new config faild,err:", err)
	}

	//init configuration struct
	appConfig = &Config{}

	//load log_level from configuration
	appConfig.logLevel = conf.String("logs::log_level")
	if len(appConfig.logLevel) == 0 {
		appConfig.logLevel = "debug"
	}
	err = loadCollectConf(conf)
	if err != nil {
		fmt.Printf("load collect conf failed, err:%v\n", err)
		return
	}
	//load log_path from configuration
	appConfig.logPath = conf.String("logs::log_path")
	if len(appConfig.logPath) == 0 {
		appConfig.logPath = "./logs"
	}
	appConfig.chanSize, err = conf.Int("main::chan_size")
	//if load chanSize failed
	if err != nil {
		appConfig.chanSize = 100
	}

	return
}
