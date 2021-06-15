package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))

}
func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	inputChan := make(chan int)
	for input.Text() != "" {
		inputChan <- 1
	}
	timeout := time.After(10 * time.Second)
	select {
	case <-timeout:
		fmt.Fprintln(c, "\t", strings.ToLower("你要是不说话我就先挂了啊"))
		c.Close()
		fmt.Print("超过十秒钟没有消息过来，断开连接")
	case <-inputChan:
		for input.Scan() {
			go echo(c, input.Text(), 1*time.Second)
		}
	}
	c.Close()
}
func main() {
	listener, err := net.Listen("tcp", "localhost:7777")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}
		go handleConn(conn)
	}

}
