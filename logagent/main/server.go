package main

import (
	"fmt"
	"learnGO/logagent/kafka"
	"learnGO/logagent/tailf"
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
	fmt.Println(msg.Msg)
	err = kafka.SendToKafka(msg.Msg, msg.Topic)
	return
}
