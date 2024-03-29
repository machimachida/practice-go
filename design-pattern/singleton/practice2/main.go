package main

import (
	"fmt"
)

type TriplesElement struct {
	id int
}

var triple = []*TriplesElement{{id: 0}, {id: 1}, {id: 2}}

func GetInstance(id int) *TriplesElement {
	return triple[id]
}

func main() {
	obj1 := GetInstance(0)
	obj2 := GetInstance(1)
	obj3 := GetInstance(2)
	obj4 := GetInstance(0)

	fmt.Println(obj1 == obj2)
	fmt.Println(obj1 == obj3)
	fmt.Println(obj1 == obj4)
}
