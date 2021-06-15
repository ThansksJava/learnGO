package container

import "fmt"

func stringOut() {
	var x string
	x = "1234567890"
	var b = make([]byte, 10, 11)
	b = append(b, []byte(x)[1])
	fmt.Print(b)

}
