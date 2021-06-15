package function

import (
	"fmt"
	"runtime"
	"sync"
)

//BiBao 闭包
func BiBao() {
	f := func(i int) {
		fmt.Printf("the num is %d\n", i)
	}
	for i := 0; i < 4; i++ {
		f(i)
	}
}

//BiBao1 闭包
func BiBao1() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			fmt.Println(i)
			wg.Done()
		}()
	}
	wg.Wait()
}
