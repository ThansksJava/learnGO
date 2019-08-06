package method

import (
	"fmt"
	"sort"
)

//Sorter 排序接口 默认实现了sort包中的Interface接口，也就是说这个golang的接口也可以继承父接口
type Sorter interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

//IntArray int数组
type IntArray []int

//Len 实现计算元素个数
func (arr IntArray) Len() int {
	return len(arr)
}

//Less 实现比较大小
func (arr IntArray) Less(i, j int) bool {
	return arr[i] > arr[j]
}

//Swap 交换元素
func (arr IntArray) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

//SortTest 测试
func SortTest() {
	data := IntArray{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}

	sort.Sort(data)
	fmt.Println(data)

}
