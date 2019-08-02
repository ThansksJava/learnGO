package container

import "fmt"

//SlinceArray 切片赋值
func SlinceArray() {
	var arr1 [10]int
	var s1 []int
	for i := 1; i <= 10; i++ {
		arr1[i-1] = i
	}
	s1 = arr1[2:4]
	fmt.Printf("切片的长度是：%d,容量是：%d\n", len(s1), cap(s1))
	for _, v := range s1 {
		fmt.Print(v)
		fmt.Print(" ")
	}
}
