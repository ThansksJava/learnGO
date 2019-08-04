//Package method 学习方法
// 类型和作用在它上面定义的方法    必须在同一个包里定义
// 间接的方式：可以先定义该类型（比如：int 或 float）的别名类型，
// 然后再为别名类型定义方法。或者像下面这样将它作为匿名类型嵌入在一个新的结构体中。当然方法只在这个别名类型上有效。
package method

import (
	"fmt"
)

// Calc 计算器
type calc struct {
	op string
	a  int
	b  int
}

// Do 计算器计算方法
// 性能的原因，`recv` 最常见的是一个指向 receiver_type 的指针
func (c *calc) Do() {
	switch c.op {
	case "+":
		fmt.Println(c.a + c.b)
	case "-":
		fmt.Println(c.a - c.b)
	case "*":
		fmt.Println(c.a * c.b)
	case "/":
		fmt.Println(c.a / c.b)
	default:
		fmt.Println("不支持的运算")
	}
}

  //PracMeth 方法测试
func PracMeth() {
	do := calc{"+", 1, 2}
	do.Do()
}
