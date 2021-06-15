package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/astaxie/beego/logs"
)

var (
	client sarama.SyncProducer
)

//InitKafka 初始化kafka
func InitKafka(addr string) (err error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	client, err = sarama.NewSyncProducer([]string{addr}, config)
	if err != nil {
		logs.Error("init kafka producer failed,err", err)
		return
	}
	logs.Debug("init kafka success")
	return
}

//SendToKafka 往kafka发送消息
func SendToKafka(data, topic string) (err error) {
	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Value = sarama.StringEncoder(data)
	//pid, offset, errSend := client.SendMessage(msg)
	_, _, errSend := client.SendMessage(msg)
	if errSend != nil {
		fmt.Println("Send msg failed,", err)
		logs.Error("Send msg failed,err :%v data:%v topic:%v", err, data, topic)
		err = errSend
		return
	}
	//logs.Debug("send success pid:%v offset:%v\n", pid, offset)
	return
}
