package function

import "fmt"

//BiBao 闭包
func BiBao() {
	f := func(i int) {
		fmt.Printf("the num is %d\n", i)
	}
	for i := 0; i < 4; i++ {
		f(i)
	}
}
