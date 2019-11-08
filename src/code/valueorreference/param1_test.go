package main

import (
	"fmt"
	"testing"
)

//UpdateStatus 更新其状态
func UpdateStatus(p People) {
	if p.status {
		p.status = false
	} else {
		p.status = true
	}
}

//TestParam 使用值调用
func TestParam(t *testing.T) {
	p1 := People{
		name:    "fengjie",
		age:     12,
		status:  false,
		feature: map[string]string{"heigit": "172", "hairColor": "black"},
	}

	UpdateStatus(p1)

	fmt.Println(p1)
}
