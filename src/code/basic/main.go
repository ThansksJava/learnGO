package main

import fmc "code/basic/function"
// import pac "code/basic/practice"
import "fmt"
func main() {
	// fmc.CallParmaFunc(1, 2, fmc.ParamFunc)
	// pac.SayH ello()
	// fmc.BiBao()

	// retf := fmc.Add();
	// fmt.Printf("2b=%d",retf(2))

	fib := fmc.Fib()
	for i := 0;i < 10;i++{
		fmt.Printf("%d ",fib())
	}


}