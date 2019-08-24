package routines_channal

import "fmt"

func counter(out chan<- int) {
	for x := 0; x < 10; x++ {
		out <- x
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}

func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func nocacheChannel() {
	num := make(chan int)
	sqrtNum := make(chan int)
	go counter(num)
	go squarer(sqrtNum, num)
	printer(sqrtNum)
}
