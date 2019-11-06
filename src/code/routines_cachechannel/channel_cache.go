package main

import (
	"fmt"
)

//func main() {
//	ch := make(chan string, 3)
//	ch <- "fengjie"
//	fmt.Print("长度：", len(ch), " ")
//	fmt.Println("容量：", cap(ch))
//}

//func f1(in chan int) {
//	for{
//		fmt.Println(<-in)
//	}
//}
//func f2(out chan int){
//	for i:=1;i < 100;i++{
//		out <- i
//	}
//}
//func main() {
//
//	out := make(chan int)
//	go f2(out)
//	time.Sleep(10)
//	go f1(out)
//}
func f1(in chan int) {
	fmt.Println(<-in)
}

func main() {
	out := make(chan int)
	out <- 2
	go f1(out)
}