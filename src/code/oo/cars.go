package main

import (
	"fmt"
)

//Any 定义空接口
type Any interface{}

//Car 定义车的结构体，相当于Java中的类了
type Car struct {
	//定义属性
	Model        string
	ManuFacturer string
	BuildYear    int
}

//Cars 定义Car的数组类型
type Cars []*Car

//Process 给Car数组类型定义一个方法
func (cs Cars) Process(f func(car *Car)) {
	for _, c := range cs {
		f(c)
	}
}

//FindAll t
func (cs Cars) FindAll(f func(car *Car) bool) Cars {
	//car的slince
	cars := make([]*Car, 0)

	cs.Process(func(c *Car) {
		if f(c) {
			cars = append(cars, c)
		}
	})

	return cars
}

//Map t
func (cs Cars) Map(f func(car *Car) Any) []Any {
	result := make([]Any, len(cs))
	ix := 0
	cs.Process(func(c *Car) {
		result[ix] = f(c)
		ix++
	})

	return result
}

//MakeSortedAppender tt
func MakeSortedAppender(manufacturers []string) (func(car *Car), map[string]Cars) {
	sortedCars := make(map[string]Cars)
	for _, m := range manufacturers {
		sortedCars[m] = make([]*Car, 0)
	}

	sortedCars["Default"] = make([]*Car, 0)

	appender := func(c *Car) {
		if _, ok := sortedCars[c.ManuFacturer]; ok {
			sortedCars[c.ManuFacturer] = append(sortedCars[c.ManuFacturer], c)
		} else {
			sortedCars["Default"] = append(sortedCars["Default"], c)
		}
	}
	return appender, sortedCars
}
func main() {
	// make some cars:
	ford := &Car{"Fiesta", "Ford", 2008}
	bmw := &Car{"XL 450", "BMW", 2011}
	merc := &Car{"D600", "Mercedes", 2009}
	bmw2 := &Car{"X 800", "BMW", 2008}
	// query:
	allCars := Cars([]*Car{ford, bmw, merc, bmw2})
	allNewBMWs := allCars.FindAll(func(car *Car) bool {
		return (car.ManuFacturer == "BMW") && (car.BuildYear > 2010)
	})
	fmt.Println("AllCars: ", allCars)
	fmt.Println("New BMWs: ", allNewBMWs)

	manufacturers := []string{"Ford", "Aston Martin", "Land Rover", "BMW", "Jaguar"}
	sortedAppender, sortedCars := MakeSortedAppender(manufacturers)
	allCars.Process(sortedAppender)
	fmt.Println("Map sortedCars: ", sortedCars)
	BMWCount := len(sortedCars["BMW"])
	fmt.Println("We have ", BMWCount, " BMWs")
}
