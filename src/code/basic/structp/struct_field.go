package structp

import "fmt"

type people struct {
	name string
	age  int
}

//StructPrint 结构体
func StructPrint() {
	var fengjie people
	// fengjie.age = 26
	// fengjie.name = "fengjie"
	fmt.Println("使用直接.赋值")
	fmt.Print(fengjie.name, fengjie.age)
	fmt.Println("使用&struct{1,2}赋值")
	someone := &people{"someone", 16}
	fmt.Println(someone)
}
