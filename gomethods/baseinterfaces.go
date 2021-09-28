package main

import "fmt"

func interfaces() {
	exercise1()
	exercise2()
}

type vehicle interface {
	License() string
	Name() string
}

type car struct {
	licenceNo string
	brand     string
}

func (c car) License() string {
	return c.licenceNo
}
func (c car) Name() string {
	return c.brand
}

func exercise1() {
	var v vehicle = car{licenceNo: "awefweaf", brand: "awfawefffff"}
	fmt.Println(v.License())
	fmt.Println(v.Name())
}

func exercise2() {
	var empty interface{}
	fmt.Printf("%T\n", empty)
	empty = 10
	fmt.Printf("%T\n", empty)
	empty = 10.12
	fmt.Printf("%T\n", empty)
	empty = []int{}
	fmt.Printf("%T\n", empty)
	empty = append(empty.([]int), 20)
	fmt.Printf("%v\n", empty)
}
