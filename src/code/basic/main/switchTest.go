package main

func switchTest() {
	x := []int{1, 2, 3}
	//;不可省略
	switch i := x[2]; { // 带初始化语句
	case i > 0:
		println("a")
	case i < 0:
		println("b")
	default:
		println("c")
	}
}
