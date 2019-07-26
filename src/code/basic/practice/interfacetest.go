package main

import "fmt"

//接口中的函数也必须全被实现
type query interface {
	queryById(name string) string
	//queryByName() string
}

//相当于一个实现类
type queryImpl struct {
}

func (queryImpl) queryById(name string) string {
	return name + "继承了接口query"
}

func testInteface() {
	//var qu  = queryImpl{}
	var qu query = queryImpl{}
	fmt.Println(qu.queryById("fengjie"))
}
