package structp

import (
	"fmt"
	"reflect"
)

type people struct { //Tag此处的警告是说成对出现比较好，但是还不是很明白怎么写 TODO
	name string "这是name字段"
	age  int    "这是age字段"
}

type apartment struct {
	name string
	num  string
	// 在一个结构体中对于  每一种数据类型  只能有一个匿名字段。
	string
	people
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

//StructTag 隐藏的tag标签
func StructTag() {
	p := people{"fengjie", 26}
	pType := reflect.TypeOf(p)
	pAgeField := pType.Field(1)
	fmt.Printf("%v\n", pAgeField.Tag)
}

//StructAnonymousFields 匿名字段
func StructAnonymousFields() {
	apart := new(apartment)
	apart.name = "云计算业务部"
	apart.num = "NO.1"
	apart.string = "瑞飞第一"
	// 重名不能这么写，内部的匿名字段的属性值不能赋值，fengjie会赋值给apart的name字段
	// 	当两个字段拥有相同的名字（可能是继承来的名字）时该怎么办呢？
	// 1. 外层名字会覆盖内层名字（但是两者的内存空间都保留），这提供了一种重载字段或方法的方式；
	// 2. 如果相同的名字在同一级别出现了两次，如果这个名字被程序使用了，将会引发一个错误（不使用没关系）。
	//    没有办法来解决这种问题引起的二义性，必须由程序员自己修正。
	// apart.name = "fengjie"
	// apart.age = 26
	apart.people = people{"某人", 100}
	fmt.Println(apart)

	fmt.Println(apart.name)
	fmt.Println(apart.people.name)
	// 外层结构体直接访问内部结构体的字段
	fmt.Println(apart.age)

}
