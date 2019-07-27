package practice

func Fib(n int) (idx int, val int) {
	if n == 0 {
		idx = 0
		val = 0
		return
	} else if n == 1 || n == 2 {
		idx = n
		val = 1
		return
	} else {
		var _, tmval1 = Fib(n - 1)
		var _, tmval2 = Fib(n - 2)
		idx = n
		val = tmval1 + tmval2
		return
	}
}
