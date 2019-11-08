//测试一下值传递跟引用传递
//结果证明取决于函数头 func (p *People) 将会是引用 func (p People)将会是值
//调用时go会自动将调用者类型转换
//使用值接收者声明方法，调用时会使
//用这个值的一个副本来执行
package main

import (
	"fmt"
	"testing"
)

//People 定义一个人类型
type People struct {
	name    string            //姓名
	age     int8              //年龄
	status  bool              //状态 生 or 死
	feature map[string]string //其他的特征放到map中
}

//特指黑人
type BlackPeople People

//InitParam People的初始化方法，初始化基本属性
func (p People) InitParam(name string, age int8, status bool) {
	p.name = name
	p.age = age
	p.status = status

}

//UpdateStatus 更新其状态
func (p *People) UpdateStatus() {
	if p.status {
		p.status = false
	} else {
		p.status = true
	}
}

//UpdateStatus 更新其状态
func (p *BlackPeople) UpdateBlackStatus() {
	if p.status {
		p.status = false
	} else {
		p.status = true
	}
}

//使用值调用
func TestParam1(t *testing.T) {
	p1 := People{
		name:    "fengjie",
		age:     12,
		status:  true,
		feature: map[string]string{"heigit": "172", "hairColor": "black"},
	}
	p1.UpdateStatus()
	fmt.Println(p1)
	t.Log(p1)
	bp1 := BlackPeople{
		name:    "fengjie",
		age:     12,
		status:  true,
		feature: map[string]string{"heigit": "172", "hairColor": "black"},
	}

	bp1.UpdateBlackStatus()

	fmt.Println(bp1)

	t.Log(bp1)
}

//使用指针调用
//结果证明go会自动将指针转变为值类型
func TestParam2(t *testing.T) {
	p1 := People{
		name:    "fengjie",
		age:     12,
		status:  true,
		feature: map[string]string{"heigit": "172", "hairColor": "black"},
	}
	(&p1).UpdateStatus()
	fmt.Println(p1)
	t.Log(p1)
	bp1 := BlackPeople{
		name:    "fengjie",
		age:     12,
		status:  true,
		feature: map[string]string{"heigit": "172", "hairColor": "black"},
	}

	(&bp1).UpdateBlackStatus()

	fmt.Println(bp1)

	t.Log(bp1)
}

//使用指针调用
//结果证明go会自动将指针转变为值类型
func TestParam3(t *testing.T) {
	p1 := People{
		name:    "fengjie",
		age:     12,
		status:  true,
		feature: map[string]string{"heigit": "172", "hairColor": "black"},
	}
	p1.UpdateStatus()
	fmt.Println(p1)
	t.Log(p1)
	bp1 := BlackPeople{
		name:    "fengjie",
		age:     12,
		status:  true,
		feature: map[string]string{"heigit": "172", "hairColor": "black"},
	}

	bp1.UpdateBlackStatus()

	fmt.Println(bp1)

	t.Log(bp1)
}
func TestParam4(t *testing.T) {
	p := People{
		name:    "fengjie",
		age:     12,
		status:  true,
		feature: map[string]string{"heigit": "172", "hairColor": "black"},
	}
	callUpdate(p)
}
func callUpdate(p People)  {
	p.UpdateStatus()
	fmt.Println(p)
}
