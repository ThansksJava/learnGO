package method

import (
	"fmt"
	"math"
)

//Circle 圆形
type Circle struct {
	radius float32
}

//TypeAssert 类型断言
func TypeAssert() {
	var areaIntf Shaper
	sq1 := new(Square)
	sq1.side = 5

	areaIntf = sq1
	// // Is Square the type of areaIntf?
	// if t, ok := areaIntf.(*Square); ok {
	// 	fmt.Printf("The type of areaIntf is: %T\n", t)
	// }
	// if u, ok := areaIntf.(*Circle); ok {
	// 	fmt.Printf("The type of areaIntf is: %T\n", u)
	// } else {
	// 	fmt.Println("areaIntf does not contain a variable of type Circle")
	// }

	switch t := areaIntf.(type) {
	case *Square:
		fmt.Printf("Type Square %T with value %v\n", t, t)
	case *Circle:
		fmt.Printf("Type Circle %T with value %v\n", t, t)
	case nil:
		fmt.Printf("nil value: nothing to check?\n")
	default:
		fmt.Printf("Unexpected type %T\n", t)
	}


}

// Area 圆形的面积实现
func (ci *Circle) Area() float32 {
	return ci.radius * ci.radius * math.Pi
}
