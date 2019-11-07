package main

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
func main() {
	// 声明第二个包含 3 个元素的指向字符串的指针数组
	// 使用字符串指针初始化这个数组

	var array2 [3]*string
	// 使用颜色为每个元素赋值
	s1 := "red"
	array2[0] = &s1
	fmt.Println(array2)
	var array3 [3]int

	fmt.Println(array3)

	array4 := [3]*string{new(string), new(string), new(string)}
	// 使用颜色为每个元素赋值
	*array4[0] = "Red"
	*array4[1] = "Blue"
	*array4[2] = "Green"
	// 将 array2 复制给 array1
	fmt.Println(array4)

	//切片
	fmt.Println("~~~~切片~~~~")
	slice1 := make([]int, 3, 10)
	fmt.Println("slice1:", len(slice1), cap(slice1))

	slice2 := []int{10, 20, 30, 40, 50}
	newSlice2 := slice2[1:]
	fmt.Println("before append slice2:", len(slice2), cap(slice2),slice2)

	fmt.Println("befor append newSlice2:", len(newSlice2), cap(newSlice2),newSlice2)

	newSlice2 = append(newSlice2, 999)

	fmt.Println("after append slece2:", len(slice2), cap(slice2),slice2)

	fmt.Println("after append newSlice2:", len(newSlice2), cap(newSlice2),newSlice2)
}
