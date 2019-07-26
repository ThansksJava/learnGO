package main

import "fmt"

type Human struct {
	name string
	age  int
}

type fly interface {
	fly()
}

func (p *Human) fly() {
	fmt.Println(p.name + "可以飞")
}
