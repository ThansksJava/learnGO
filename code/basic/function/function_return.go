package function

//Add 函数作为返回值
func Add() func(b int) int{
	return func(b int) int {
		return b+b
	}

}