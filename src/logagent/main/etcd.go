package main

import (
	"context"
	"encoding/json"
	"fmt"
	cm "logagent/commonconf"
	"logagent/tailf"
	"strings"
	"time"

	"github.com/astaxie/beego/logs"

	etcd_client "go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/mvcc/mvccpb"
)

//EtcdClient 结构体聚合Etcd配置
type EtcdClient struct {
	client *etcd_client.Client
	keys   []string
}

var (
	etcdClient *EtcdClient
)

func initEtcd(addr string, key string) (collectConf []cm.CollectConf, err error) {
	cli, err := etcd_client.New(etcd_client.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:22379"},
		DialTimeout: 5 * time.Second,
	})

	if err != nil {
		logs.Error("connect etcd faild,err:", err)
		return
	}
	etcdClient = &EtcdClient{
		client: cli,
	}
	if strings.HasSuffix(key, "/") == false {
		key = key + "/"
	}
	for _, ip := range localIPArray {
		etcdKey := fmt.Sprintf("%s%s", key, ip)
		etcdClient.keys = append(etcdClient.keys, etcdKey)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		resp, err := cli.Get(ctx, etcdKey)
		if err != nil {
			logs.Error("client get from etcd failed,err :%v", err)
			continue
		}
		cancel()

		logs.Debug("response from etcd: %v", resp.Kvs)
		for _, v := range resp.Kvs {
			if string(v.Key) == etcdKey {
				err = json.Unmarshal(v.Value, &collectConf)
				if err != nil {
					logs.Error("unmarshal failed,err:%v", err)
					continue
				}
				logs.Debug("log config is %v", collectConf)
			}
		}
	}
	initEtcdWatcher()
	return
}

func initEtcdWatcher() {
	for _, key := range etcdClient.keys {
		// 开启一个routine监听每一个key
		go watchKey(key)
	}
}

func watchKey(key string) {
	cli, err := etcd_client.New(etcd_client.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		logs.Error("connect etcd failed, err:", err)
		return
	}

	logs.Debug("begin watch key:%s", key)
	for {
		rch := cli.Watch(context.Background(), key)
		var collectConf []cm.CollectConf
		var getConfSucc = true
		for wresp := range rch {
			for _, ev := range wresp.Events {
				if ev.Type == mvccpb.DELETE {
					logs.Warn("key[%s] 's config deleted", key)
					continue
				}

				if ev.Type == mvccpb.PUT && string(ev.Kv.Key) == key {
					err = json.Unmarshal(ev.Kv.Value, &collectConf)
					if err != nil {
						logs.Error("key [%s], Unmarshal[%s], err:%v ", err)
						getConfSucc = false
						continue
					}
				}
				logs.Debug("get config from etcd, %s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
			}

			if getConfSucc {
				logs.Debug("get config from etcd succ, %v", collectConf)
				//传入新的全量配置，通过判断配置是否在已有配置中进行筛选更新
				tailf.UpdateConfig(collectConf)
			}
		}

	}
}
