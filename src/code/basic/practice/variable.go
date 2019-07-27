package practice

import "fmt"

//只能用于定义全局变量
var (
	name1 string
	age1  int
)

func testVariable() {
	////此种方式只能用于函数内部
	//name := "fengjie"
	//age := 18
	//fmt.Print("我的名字是"+name)
	//fmt.Print(",今年已经")
	//fmt.Print(age)
	//fmt.Println("岁了!")
	////常量
	//const a = "常量"
	////var unused string = "未使用变量" 无法编译通过
	////数组
	////var array[5] int
	////array[1] = 0
	////
	////var array1 = [5] int {1,2,3,4,5}
	////fmt.Print(array1[1])
	////
	////var array2 = []int {1,2,3,4,5}
	////fmt.Println(array2)
	//
	//var b int = 20 /* 声明实际变量 */
	//var ip *int    /* 声明指针变量 */
	//
	//ip = &b  /* 指针变量的存储地址 */
	//
	//fmt.Printf("b 变量的地址是: %x\n", &b  )
	//
	///* 指针变量的存储地址 */
	//fmt.Printf("ip 变量储存的指针地址: %x\n", ip )
	//
	///* 使用指针访问值 */
	//fmt.Printf("*ip 变量的值: %d\n", *ip )
	//
	////结构体
	//type struct1 struct {
	//	pid int
	//	pname string
	//	page int
	//}
	//fmt.Println(struct1{18,"jack",10})
	//fmt.Println(struct1{pid:18,pname:"jack",page:10})
	//var someperson1 struct1
	//someperson1.pid = 2222
	//someperson1.pname = "tom"
	//someperson1.page = 18
	//fmt.Print("someperson1:")
	//fmt.Println(someperson1)

	//这是我们使用range去求一个slice的和。使用数组跟这个很类似
	//nums := []int{2, 3, 4}
	//sum := 0
	//for _, num := range nums {
	//	sum += num
	//}
	//fmt.Println("sum:", sum)
	////在数组上使用range将传入index和值两个变量。上面那个例子我们不需要使用该元素的序号，所以我们使用空白符"_"省略了。有时侯我们确实需要知道它的索引。
	//for i, num := range nums {
	//	if num == 3 {
	//		fmt.Println("index:", i)
	//	}
	//}
	////range也可以用在map的键值对上。
	//kvs := map[string]string{"a": "apple", "b": "banana"}
	//for k, v := range kvs {
	//	fmt.Printf("%s -> %s\n", k, v)
	//}
	////range也可以用来枚举Unicode字符串。第一个参数是字符的索引，第二个是字符（Unicode的值）本身。
	//for i, c := range "go" {
	//	fmt.Println(i, c)
	//}

	var mymap map[string]string = make(map[string]string)
	mymap["name"] = "fengjie"
	mymap["age"] = "18"
	for k, v := range mymap {
		fmt.Println("key:" + k)
		fmt.Println("value:" + v)
	}
	fmt.Println("删除name项")
	delete(mymap, "name")
	for k, v := range mymap {
		fmt.Println("key:" + k)
		fmt.Println("value:" + v)
	}
}
