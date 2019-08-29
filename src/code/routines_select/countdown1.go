package main

import (
	"fmt"
	"time"
)

func countdown1() {
	fmt.Println("Commencing countdown.")
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		fmt.Println(<-tick)
	}
	launch()
}
func launch() {
	fmt.Println("Lift off!")
}
