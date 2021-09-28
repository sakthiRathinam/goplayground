package main

import "fmt"

func iota_test() {
	const (
		a = -(iota + 2)
		b
		c
	)
	fmt.Printf("iota value of 3 %d/n", c)

}
