package method

import "fmt"

//Shaper 接口名字
type Shaper interface {
	Area() float32
}

//Square 正方形结构体
type Square struct {
	side float32
}

//Area 计算面积的方法
func (s *Square) Area() float32 {
	return s.side * s.side
}

//AreaTest 测试这个接口
func AreaTest() {
	sq1 := new(Square)
	sq1.side = 5
	var areaIntf Shaper
	areaIntf = sq1

	fmt.Printf("面积是：%f", areaIntf.Area())
}
