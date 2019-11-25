package main

import (
	"fmt"
	"net"
)

var (
	localIPArray []string
)

//init 会在main函数之前执行
func init() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(fmt.Sprintf("get local ip faild,%v", err))
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				localIPArray = append(localIPArray, ipnet.IP.String())
			}
		}
	}
	fmt.Println(localIPArray)
}
