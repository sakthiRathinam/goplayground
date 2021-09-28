package main

import "fmt"

func main() {
	exercise1()
	exercise2()
	exercise3()
}

func exercise1() {
	x := 10.10
	fmt.Printf("%p\n", &x)
	ptr := &x
	fmt.Printf("%v %T\n", ptr, ptr)
	fmt.Printf("%p %v\n", &ptr, *ptr)

}

func exercise2() {
	x, y := 10, 2
	ptrx, ptry := &x, &y
	z := *ptrx / *ptry
	fmt.Println(z)
}

func swap(x *float64, y *float64) {
	*x, *y = *y, *x
}

func exercise3() {
	x, y := 5.5, 8.8
	swap(&x, &y)
	fmt.Println(x, y)
}
