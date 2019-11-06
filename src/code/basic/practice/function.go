package practice

import "fmt"

type person struct {
	name string
	age  int
}

func nort(person1 person) {
	fmt.Println("这个函数没有返回值")
}
func rt(person1 person) (string, int) {
	fmt.Println("这个函数有返回值") 
	return "冯杰", 18
}

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}
func testDefer(x, y int) (z int) {
	defer func() {
		fmt.Println(z)
	}()
	z = x + y
	return z + 200
}
func testFunction() {
	var person person
	person.name = "fengjie1"
	person.age = 19
	nort(person)
	fmt.Println(rt(person))
	fmt.Println("输出斐波那契数列")
	for i := 0; i < 10; i++ {
		fmt.Printf("%d\t", fibonacci(i))
	}
}
func main_testDefer() {
	//会输出两次
	//203
	//203
	fmt.Println(testDefer(1, 2))
}
