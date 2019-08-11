package iop

import "fmt"

var (
	firstName, lastName, s string
	i                      int
	f                      float32
	input                  = "24.5 25 abc"
	format                 = "%f %d %s"
)

func readInputTest() {
	fmt.Println("请输入：")
	fmt.Scanln(&firstName, &lastName)
	fmt.Printf("Hi %s,%s", lastName, firstName)
	fmt.Sscanf(input, format, &f, &i, &s)
	fmt.Println("使用Sscanf读取的数据", f, i, s)
}
