package function

//Fib 使用闭包实现斐波那契数列,网上答案
func Fib() func() int {
	var preVal1 = 0
	var preVal2 = 1
	return func() int {
		fibVal := preVal1
		preVal1, preVal2 = preVal2, (preVal1 + preVal2)
		return fibVal
	}
}
