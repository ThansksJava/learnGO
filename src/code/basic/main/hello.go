package main

import "fmt"

func main() {
	var idx int
	var val int
	for i := 1; i <= 10; i++ {
		idx, val = Fib(i)
		fmt.Printf("序号: %d,值: %d", idx, val)
	}
}
