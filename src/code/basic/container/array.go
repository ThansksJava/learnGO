package container

import "fmt"

//ForArray 循环遍历数组
func ForArray() {
	var arr1 [10]int
	for i := 1; i <= 10; i++ {
		arr1[i-1] = i
	}
	for i := 1; i <= 10; i++ {
		fmt.Print(arr1[i-1])
		fmt.Print(" ")
	}
	for _, v := range arr1 {
		fmt.Print(v)
		fmt.Print(" ")
	}
}
