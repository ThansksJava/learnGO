package main

import (
	"github.com/astaxie/beego/logs"
	"logagent/kafka"
	"logagent/tailf"
	"time"
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
	err = kafka.SendToKafka(msg.Msg, msg.Topic)
	return
}
