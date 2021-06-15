package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		for i := 0; i < 1000; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	go func() {
		for {
			x, ok := <-ch1
			if !ok {
				break
			}
			ch2 <- (x * x)
		}
		close(ch2)
	}()
	for {
		x, ok := <-ch2
		if !ok {
			break
		}
		fmt.Println(x)
	}
}
