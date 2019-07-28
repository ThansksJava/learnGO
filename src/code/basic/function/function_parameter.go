package function

import "fmt"

//ParamFunc 定义一个函数
func ParamFunc(a int, b int) {
	fmt.Printf("a=%d,b=%d", a, b)
}

//CallParmaFunc 调用闭包函数
func CallParmaFunc(a int, b int, f func(int, int)) {
	f(a, b)
}
