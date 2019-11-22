package main

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"go.etcd.io/etcd/clientv3"
)

func TestWatch(t *testing.T) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect failed, err:", err)
		return
	}

	fmt.Println("connect succ")
	defer cli.Close()
	for i := 0; i < 10; i++ {
		cli.Put(context.Background(), "/logagent/conf/", strconv.Itoa(rand.Int()))
	}

}
