package main

import (
	"fmt"
	"logagent/tailf"
	"time"

	"github.com/astaxie/beego/logs"
)

func serverRun() (err error) {
	for {
		msg := tailf.GetOneOnLine()
		err = sendToKafka(msg)
		if err != nil {
			logs.Error("send to kafka failed,err:%v", err)
			time.Sleep(time.Second)
			continue
		}
	}
	return
}
func sendToKafka(msg *tailf.TextMsg) (err error) {
	fmt.Printf("read msg:%s,read topic:%s\n", msg.Msg, msg.Topic)
	return
}
