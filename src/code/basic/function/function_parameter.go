package function

import "fmt"

func ParamFunc(a int, b int) {
	fmt.Printf("a=%d,b=%d", a, b)
}

func CallParmaFunc(a int, b int, f func(int, int)) {
	f(a, b)
}
