//elastic search 初始化
package main

import (
	"fmt"

	elastic "gopkg.in/olivere/elastic.v2"
)

type LogMessage struct {
	App     string
	Topic   string
	Message string
}

var (
	esClient *elastic.Client
)

func initES(addr string) (err error) {

	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(addr))
	if err != nil {
		fmt.Println("connect es error", err)
		return
	}
	esClient = client
	return
}
func sendToES(topic string, data []byte) (err error) {

	msg := &LogMessage{}
	msg.Topic = topic
	msg.Message = string(data)
	esClient.Delete()
	_, err = esClient.Index().
		Index(topic).
		Type(topic).
		//Id(fmt.Sprintf("%d", i)).
		BodyJson(msg).
		Do()
	if err != nil {
		// Handle error
		panic(err)
		return
	}
	return
}
