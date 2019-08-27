package main

import (
	"fmt"
)

func main() {
	ch := make(chan string, 3)
	for {
		ch <- "fengjie"
		fmt.Print("长度：", len(ch), " ")
		fmt.Println("容量：", cap(ch))
	}
}
