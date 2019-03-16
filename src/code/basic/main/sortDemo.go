package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type humans struct {
	name string
	age  int
}

type black []humans

func (af black) Len() int {
	return len(af)
}
func (af black) Less(i, j int) bool {
	return af[i].age > af[j].age
}
func (af black) Swap(i, j int) {
	temp := af[i]
	af[i] = af[j]
	af[j] = temp
}
func testSort() {
	var persons black
	var i int
	for i = 0; i < 10; i++ {
		p := humans{
			name: fmt.Sprintf("黑人%d", i),
			age:  rand.Intn(100),
		}
		persons = append(persons, p)
	}
	fmt.Println("before sort")
	fmt.Println(persons)
	fmt.Println("after sort")
	sort.Sort(persons)
	fmt.Println(persons)

}
