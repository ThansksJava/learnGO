package main

import (
	"fmt"
	"gopkg.in/olivere/elastic.v2"
)

var (
	esClient *elastic.Client
	addr     string
)

func init() {
	addr = "http://129.223.237.83:9200"
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(addr))
	if err != nil {
		fmt.Println("connect es error", err)
		return
	}
	esClient = client
}
func main() {
	esClient.Delete().Index("nginx_log").Do()
}
