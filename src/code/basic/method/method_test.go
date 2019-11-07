package method

import (
	"fmt"
	"testing"
)

type father interface {
	PrintHello()
}

type son struct {
	name string
}

func (s son) PrintHello() {
	fmt.Print("hello " + s.name)
}
func TestInherited(t *testing.T) {
	var s son
	var f father
	f = s
	f.PrintHello()
}
