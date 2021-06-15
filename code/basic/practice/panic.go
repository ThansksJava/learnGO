package practice
import "fmt"
import "testing"
//PanicTest panic
func PanicTest() {
    f(3)
}
func f(x int) {
    fmt.Printf("f(%d)\n", x+0/x) // panics if x == 0
    
    defer fmt.Printf("defer %d\n", x)
    f(x - 1)
}

func TestF(t *testing.T){
    PanicTest()
}