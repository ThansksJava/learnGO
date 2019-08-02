package container

import "fmt"

//Map 给map赋值
func Map() {
	fmt.Print("make初始化赋值即使只是大概知道容量，也最好先标明\n")
	// var Map1 = make(map[string]int,100)
	var Map1 = make(map[string]int)
	Map1["name"] = 2222
	Map1["age"] = 3333
	for k, v := range Map1 {
		fmt.Printf("key:%s,value:%d\n", k, v)
	}
	fmt.Print("用 {key1: val1, key2: val2} 的描述方法来初始化赋值\n")
	var map2 = map[string]string{"name": "fengjie", "age": "22"}
	for k, v := range map2 {
		fmt.Printf("key:%s,value:%s\n", k, v)
	}
	fmt.Print("val1, isPresent = map1[key1] 如果 key1 存在于 map1，val1 就是 key1 对应的 value 值，并且 isPresent为true\n")
	value, exsit := map2["name"]
	value1, exsit1 := map2["ssss"]

	fmt.Print(value, exsit, "\n")
	fmt.Print(value1, exsit1,"\n")

		// Version A:
		items := make([]map[int]int, 5)
		for i:= range items {
			items[i] = make(map[int]int, 1)
			items[i][1] = 2
		}
		fmt.Printf("Version A: Value of items: %v\n", items)
	
		// Version B: NOT GOOD!
		items2 := make([]map[int]int, 5)
		for _, item := range items2 {
			item = make(map[int]int, 1) // item is only a copy of the slice element.
			item[1] = 2 // This 'item' will be lost on the next iteration.
		}
		fmt.Printf("Version B: Value of items: %v\n", items2)
}
