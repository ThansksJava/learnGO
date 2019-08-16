package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	MustCopy(os.Stdout, conn)
}

//MustCopy 从source流往destination流写入数据
func MustCopy(dst io.Writer, src net.Conn) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
